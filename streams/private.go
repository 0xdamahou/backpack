package streams

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	DefaultWsURL = "wss://ws.backpack.exchange"
)

type BackpackWebsocketClient struct {
	ApiKey     string
	ApiSecret  string
	PrivateKey ed25519.PrivateKey
	wsURL      string
}

func NewBackpackWebsocketClient(wsURL string, apiKey string, apiSecret string) *BackpackWebsocketClient {
	client := &BackpackWebsocketClient{ApiKey: apiKey, ApiSecret: apiSecret, wsURL: wsURL}
	seed, err := base64.StdEncoding.DecodeString(apiSecret)
	if err != nil {
		panic(err)
		return nil
	}
	client.PrivateKey = ed25519.NewKeyFromSeed(seed)
	return client
}

// SubscriptionRequest 订阅请求结构体
type SubscriptionRequest struct {
	Method    string   `json:"method"`
	Params    []string `json:"params"`
	Signature []string `json:"signature,omitempty"` // 仅私有订阅时使用
}

// Message 接收消息结构体
type Message struct {
	Stream string      `json:"stream,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	ID     *int        `json:"id,omitempty"`
	Error  *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// ED25519
func (bpc *BackpackWebsocketClient) generateED25519Signature(instruction string) string {
	signature := ed25519.Sign(bpc.PrivateKey, []byte(instruction))
	return base64.StdEncoding.EncodeToString(signature)
}

func (bpc *BackpackWebsocketClient) OrderUpdate(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := "account.orderUpdate"
	if symbol != "" {
		stream = stream + "." + symbol
	}
	return bpc.Subscribe([]string{stream}, true, done, result)
}
func (bpc *BackpackWebsocketClient) PositionUpdate(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := "account.positionUpdate"
	if symbol != "" {
		stream = stream + "." + symbol
	}
	return bpc.Subscribe([]string{stream}, true, done, result)
}

func (bpc *BackpackWebsocketClient) RfqUpdate(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := "account.rfqUpdate"
	if symbol != "" {
		stream = stream + "." + symbol
	}
	return bpc.Subscribe([]string{stream}, true, done, result)
}

func (bpc *BackpackWebsocketClient) Unsubscribe(c *websocket.Conn, streams []string) {
	defer c.Close()
	unsub := SubscriptionRequest{
		Method: "UNSUBSCRIBE",
		Params: streams,
	}
	err := c.WriteJSON(unsub)
	if err != nil {
		log.Println("write unsubscription:", err)
	}
	log.Printf("Unsubscribe : %v", streams)
}

// 连接并订阅WebSocket
func (bpc *BackpackWebsocketClient) Subscribe(streams []string, isPrivate bool, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	connect := func() (*websocket.Conn, error) {
		conn, _, err := websocket.DefaultDialer.Dial(bpc.wsURL, nil)
		if err != nil {
			log.Printf("dial: %v", err)
			return nil, err
		}

		subReq := SubscriptionRequest{
			Method: "SUBSCRIBE",
			Params: streams,
		}

		if isPrivate {
			timestamp := time.Now().UnixMilli()
			window := int64(5000)
			instruction := fmt.Sprintf("instruction=subscribe&timestamp=%d&window=%d", timestamp, window)
			signature := bpc.generateED25519Signature(instruction)
			subReq.Signature = []string{
				bpc.ApiKey,
				signature,
				fmt.Sprintf("%d", timestamp),
				fmt.Sprintf("%d", window),
			}
		}

		if err := conn.WriteJSON(subReq); err != nil {
			log.Printf("write subscription: %v", err)
			conn.Close()
			return nil, err
		}
		log.Printf("Subscribed to: %v", streams)

		conn.SetPingHandler(func(appData string) error {
			return conn.WriteMessage(websocket.PongMessage, nil)
		})
		conn.SetCloseHandler(func(code int, text string) error {
			log.Printf("Received close frame: code=%d, text=%s", code, text)
			return nil
		})

		return conn, nil
	}

	conn, err := connect()
	if err != nil {
		return nil, err
	}

	const maxReconnectAttempts = 5
	const initialBackoff = 1 * time.Second

	go func() {
		defer close(result) // 确保退出时关闭 result 通道
		reconnectAttempts := 0
		backoff := initialBackoff

		readCh := make(chan []byte)
		readErrCh := make(chan error)

		readMessages := func(c *websocket.Conn) {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					readErrCh <- err
					return
				}
				readCh <- message
			}
		}

		go readMessages(conn)

		for {
			select {
			case <-done:
				log.Println("Received done signal, closing connection")
				conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				conn.Close()
				return

			case err := <-readErrCh:
				log.Printf("WebSocket read error: %v, attempting to reconnect...", err)
				conn.Close()

				if reconnectAttempts >= maxReconnectAttempts {
					log.Printf("Max reconnection attempts (%d) reached, giving up", maxReconnectAttempts)
					return
				}

				reconnectAttempts++
				log.Printf("Reconnection attempt %d/%d in %v", reconnectAttempts, maxReconnectAttempts, backoff)
				time.Sleep(backoff)
				backoff *= 2
				if backoff > 30*time.Second {
					backoff = 30 * time.Second
				}

				newConn, connErr := connect()
				if connErr != nil {
					log.Printf("Reconnection failed: %v", connErr)
					continue
				}
				conn = newConn
				reconnectAttempts = 0
				backoff = initialBackoff
				go readMessages(conn)

			case msg := <-readCh:
				result <- msg
			}
		}
	}()

	return conn, nil
}

func main() {
	// 示例：订阅公共streams
	//publicStreams := []string{
	//	"bookTicker.MEW_USDC",
	//}
	done := make(chan struct{})
	result := make(chan []byte, 20)
	apiKey := ""
	apiSecret := ""
	bpc := NewBackpackWebsocketClient(DefaultWsURL, apiKey, apiSecret)

	//bpc.Subscribe(publicStreams, false, done, result)
	//
	//// 示例：订阅私有streams
	////privateStreams := []string{
	////	"account.orderUpdate:SOL_USDC",
	////	"account.balanceUpdate",
	////}
	////go bpc.Subscribe(privateStreams, true)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	conn, err := bpc.OrderUpdate("", done, result)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		select {
		case <-interrupt:
			log.Println("Interrupt received, closing connection...")
			done <- struct{}{}
			return
		default:
			bs := <-result
			log.Println(string(bs))
		}

	}
}

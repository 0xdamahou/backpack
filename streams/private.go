package streams

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	DefaultWsURL = "wss://ws.backpack.exchange"
	apiKey       = "" // 替换为你的 API Key
	apiSecret    = "" // 替换为你的 Private Key

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
	c, _, err := websocket.DefaultDialer.Dial(bpc.wsURL, nil)
	if err != nil {
		log.Println("dial:", err)
		return nil, err
	}
	// 建立WebSocket连接
	subReq := SubscriptionRequest{
		Method: "SUBSCRIBE",
		Params: streams,
	}

	if isPrivate {
		// 私有订阅需要签名
		timestamp := time.Now().UnixMilli()
		window := int64(5000)
		instruction := fmt.Sprintf("instruction=subscribe&timestamp=%d&window=%d", timestamp, window)
		signature := bpc.generateED25519Signature(instruction)

		subReq.Signature = []string{
			apiKey,
			signature,
			fmt.Sprintf("%d", timestamp),
			fmt.Sprintf("%d", window),
		}
	}
	marshal, err := json.Marshal(subReq)

	log.Printf("%+v %+v\n", string(marshal), err)
	// 发送订阅请求
	err = c.WriteJSON(subReq)
	if err != nil {
		log.Println("write subscription:", err)
		c.Close()
		time.Sleep(5 * time.Second)
		return nil, err
	}
	log.Printf("Subscribed to: %v", streams)

	// 设置Ping处理器以响应服务器的Ping
	c.SetPingHandler(func(appData string) error {
		//log.Println("Received ping from server, sending pong")
		return c.WriteMessage(websocket.PongMessage, nil)
	})

	// 设置Close处理器
	c.SetCloseHandler(func(code int, text string) error {
		log.Printf("Received close frame: code=%d, text=%s, reconnecting in 5 seconds...", code, text)
		return nil
	})

	// 处理接收消息
	go func() {
		// 创建一个通道用于传递读取结果
		readCh := make(chan []byte)
		readErrCh := make(chan error)

		// 完全分离读取操作
		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					readErrCh <- err
					return
				}
				readCh <- message
			}
		}()

		// 主循环监听多个通道
		for {
			select {
			case <-done:
				// 收到关闭信号
				log.Println("Received done signal, closing connection")
				err := c.WriteMessage(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("Error sending close message:", err)
				}
				time.Sleep(time.Second)
				return

			case err := <-readErrCh:
				// 读取错误
				log.Println("WebSocket read error:", err)
				return

			case msg := <-readCh:
				// 成功读取消息
				result <- msg
			}
		}
	}()
	return c, nil

}

func main() {
	// 示例：订阅公共streams
	//publicStreams := []string{
	//	"bookTicker.MEW_USDC",
	//}
	done := make(chan struct{})
	result := make(chan []byte, 20)
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

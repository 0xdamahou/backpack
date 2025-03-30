package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

// EventHeader represents the common header fields for all event types
// It can be used to determine the specific event type before unmarshalling
// the complete message into the appropriate struct
type EventHeader struct {
	EventType string `json:"e"` // Event type
	EventTime int64  `json:"E"` // Event time in microseconds
}

// Constants for different event types
const (
	EventTypeOrderAccepted  = "orderAccepted"
	EventTypeOrderCancelled = "orderCancelled"
	EventTypeOrderExpired   = "orderExpired"
	EventTypeOrderFill      = "orderFill"
	EventTypeOrderModified  = "orderModified"
	EventTypeTriggerPlaced  = "triggerPlaced"
	EventTypeTriggerFailed  = "triggerFailed"

	// Position events
	EventTypePositionAdjusted = "positionAdjusted"
	EventTypePositionOpened   = "positionOpened"
	EventTypePositionClosed   = "positionClosed"

	// RFQ (Request for Quote) events
	EventTypeRFQActive      = "rfqActive"      // Indicates an RFQ is active and open for quotes
	EventTypeQuoteAccepted  = "quoteAccepted"  // Indicates a quote submitted by the maker has been accepted
	EventTypeQuoteCancelled = "quoteCancelled" // Indicates a quote has been cancelled (due to submission, RFQ being filled, refreshed, cancelled, or expired)
	EventTypeRFQFilled      = "rfqFilled"      // Indicates an RFQ has been fully filled with a quote from the maker

	// Market data events
	EventTypeBookTicker   = "bookTicker"
	EventTypeDepth        = "depth"
	EventTypeKline        = "kline"
	EventTypeLiquidation  = "liquidation"
	EventTypeMarkPrice    = "markPrice"
	EventTypeTicker       = "ticker"
	EventTypeOpenInterest = "openInterest"
	EventTypeTrade        = "trade"
)

func (bpc *BackpackWebsocketClient) BookTicker(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := fmt.Sprintf("bookTicker.%s", symbol)
	return bpc.Subscribe([]string{stream}, false, done, result)
}

func (bpc *BackpackWebsocketClient) Depth(symbol string, isRealtime bool, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	var stream string
	if isRealtime {
		stream = fmt.Sprintf("depth.%s", symbol)
	} else {
		stream = fmt.Sprintf("depth.200ms.%s", symbol)
	}
	return bpc.Subscribe([]string{stream}, false, done, result)
}

func (bpc *BackpackWebsocketClient) Kline(symbol string, interval string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := fmt.Sprintf("kline.%s.%s", interval, symbol)
	return bpc.Subscribe([]string{stream}, false, done, result)
}
func (bpc *BackpackWebsocketClient) Liquidation(done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := "liquidation"
	return bpc.Subscribe([]string{stream}, false, done, result)
}

func (bpc *BackpackWebsocketClient) MarkPrice(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := fmt.Sprintf("markPrice.%s", symbol)
	return bpc.Subscribe([]string{stream}, false, done, result)
}

func (bpc *BackpackWebsocketClient) Ticker(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := fmt.Sprintf("ticker.%s", symbol)
	return bpc.Subscribe([]string{stream}, false, done, result)
}

func (bpc *BackpackWebsocketClient) OpenInterest(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := fmt.Sprintf("openInterest.%s", symbol)
	return bpc.Subscribe([]string{stream}, false, done, result)
}

func (bpc *BackpackWebsocketClient) Trade(symbol string, done chan struct{}, result chan []byte) (*websocket.Conn, error) {
	stream := fmt.Sprintf("trade.%s", symbol)
	return bpc.Subscribe([]string{stream}, false, done, result)
}

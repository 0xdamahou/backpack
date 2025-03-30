package stream

// BookTickerEvent represents a book ticker update with best bid and ask prices
type BookTickerEvent struct {
	EventType       string `json:"e"` // Event type
	EventTime       int64  `json:"E"` // Event time in microseconds
	Symbol          string `json:"s"` // Symbol
	AskPrice        string `json:"a"` // Inside ask price
	AskQuantity     string `json:"A"` // Inside ask quantity
	BidPrice        string `json:"b"` // Inside bid price
	BidQuantity     string `json:"B"` // Inside bid quantity
	UpdateID        string `json:"u"` // Update ID of event
	EngineTimestamp int64  `json:"T"` // Engine timestamp in microseconds
}

// DepthEvent represents an order book depth update event
type DepthEvent struct {
	EventType       string     `json:"e"` // Event type
	EventTime       int64      `json:"E"` // Event time in microseconds
	Symbol          string     `json:"s"` // Symbol
	Asks            [][]string `json:"a"` // Asks - [price, quantity] pairs
	Bids            [][]string `json:"b"` // Bids - [price, quantity] pairs
	FirstUpdateID   int64      `json:"U"` // First update ID in event
	LastUpdateID    int64      `json:"u"` // Last update ID in event
	EngineTimestamp int64      `json:"T"` // Engine timestamp in microseconds
}

// KlineEvent represents a candlestick/kline update event
type KlineEvent struct {
	EventType       string `json:"e"` // Event type
	EventTime       int64  `json:"E"` // Event time in microseconds
	Symbol          string `json:"s"` // Symbol
	StartTime       int64  `json:"t"` // K-Line start time in seconds
	CloseTime       int64  `json:"T"` // K-Line close time in seconds
	OpenPrice       string `json:"o"` // Open price
	ClosePrice      string `json:"c"` // Close price
	HighPrice       string `json:"h"` // High price
	LowPrice        string `json:"l"` // Low price
	BaseAssetVolume string `json:"v"` // Base asset volume
	NumberOfTrades  int64  `json:"n"` // Number of trades
	IsClosed        bool   `json:"X"` // Is this k-line closed?
}

// LiquidationEvent represents a liquidation event
type LiquidationEvent struct {
	EventType       string `json:"e"` // Event type
	EventTime       int64  `json:"E"` // Event time in microseconds
	Quantity        string `json:"q"` // Quantity
	Price           string `json:"p"` // Price
	Side            string `json:"S"` // Side
	Symbol          string `json:"s"` // Symbol
	EngineTimestamp int64  `json:"T"` // Engine timestamp in microseconds
}

// MarkPriceEvent represents a mark price update event
type MarkPriceEvent struct {
	EventType       string `json:"e"` // Event type
	EventTime       int64  `json:"E"` // Event time in microseconds
	Symbol          string `json:"s"` // Symbol
	MarkPrice       string `json:"p"` // Mark price
	EstFundingRate  string `json:"f"` // Estimated funding rate
	IndexPrice      string `json:"i"` // Index price
	NextFundingTime int64  `json:"n"` // Next funding timestamp in microseconds
}

// TickerEvent represents a 24hr ticker price change statistics
type TickerEvent struct {
	EventType        string `json:"e"` // Event type
	EventTime        int64  `json:"E"` // Event time in microseconds
	Symbol           string `json:"s"` // Symbol
	FirstPrice       string `json:"o"` // First price
	LastPrice        string `json:"c"` // Last price
	HighPrice        string `json:"h"` // High price
	LowPrice         string `json:"l"` // Low price
	BaseAssetVolume  string `json:"v"` // Base asset volume
	QuoteAssetVolume string `json:"V"` // Quote asset volume
	NumberOfTrades   int64  `json:"n"` // Number of trades
}

// OpenInterestEvent represents an open interest update event
type OpenInterestEvent struct {
	EventType    string `json:"e"` // Event type
	EventTime    int64  `json:"E"` // Event time in microseconds
	Symbol       string `json:"s"` // Symbol
	OpenInterest string `json:"o"` // Open interest in contracts
}

// TradeEvent represents a trade execution event
type TradeEvent struct {
	EventType       string `json:"e"` // Event type
	EventTime       int64  `json:"E"` // Event time in microseconds
	Symbol          string `json:"s"` // Symbol
	Price           string `json:"p"` // Price
	Quantity        string `json:"q"` // Quantity
	BuyerOrderID    string `json:"b"` // Buyer order ID
	SellerOrderID   string `json:"a"` // Seller order ID
	TradeID         int64  `json:"t"` // Trade ID
	EngineTimestamp int64  `json:"T"` // Engine timestamp in microseconds
	IsBuyerMaker    bool   `json:"m"` // Is the buyer the maker?
}

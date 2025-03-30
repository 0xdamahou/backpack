package main

// OrderEvent represents a trading order event with all associated data
type OrderEvent struct {
	EventType        string `json:"e"` // Event type
	EventTime        int64  `json:"E"` // Event time in microseconds
	Symbol           string `json:"s"` // Symbol
	ClientOrderID    int64  `json:"c"` // Client order ID
	Side             string `json:"S"` // Side
	OrderType        string `json:"o"` // Order type
	TimeInForce      string `json:"f"` // Time in force
	Quantity         string `json:"q"` // Quantity
	QuantityInQuote  string `json:"Q"` // Quantity in quote
	Price            string `json:"p"` // Price
	TriggerPrice     string `json:"P"` // Trigger price
	TriggerBy        string `json:"B"` // Trigger by
	TakeProfitPrice  string `json:"a"` // Take profit trigger price
	StopLossPrice    string `json:"b"` // Stop loss trigger price
	TakeProfitBy     string `json:"d"` // Take profit trigger by
	StopLossBy       string `json:"g"` // Stop loss trigger by
	TriggerQuantity  string `json:"Y"` // Trigger quantity
	OrderState       string `json:"X"` // Order state
	ExpiryReason     string `json:"R"` // Order expiry reason
	OrderID          string `json:"i"` // Order ID
	TradeID          int64  `json:"t"` // Trade ID
	FillQuantity     string `json:"l"` // Fill quantity
	ExecutedQty      string `json:"z"` // Executed quantity
	ExecutedQtyQuote string `json:"Z"` // Executed quantity in quote
	FillPrice        string `json:"L"` // Fill price
	IsMaker          bool   `json:"m"` // Whether the order was maker
	Fee              string `json:"n"` // Fee
	FeeSymbol        string `json:"N"` // Fee symbol
	SelfTradePrev    string `json:"V"` // Self trade prevention
	EngineTimestamp  int64  `json:"T"` // Engine timestamp in microseconds
	Origin           string `json:"O"` // Origin of the update
	RelatedOrderID   string `json:"I"` // Related order ID
}

// PositionEvent represents a trading position event with all associated data
type PositionEvent struct {
	EventType           string  `json:"e"` // Event type
	EventTime           int64   `json:"E"` // Event time in microseconds
	Symbol              string  `json:"s"` // Symbol
	BreakEventPrice     float64 `json:"b"` // Break event price
	EntryPrice          float64 `json:"B"` // Entry price
	LiquidationPrice    float64 `json:"l"` // Estimated liquidation price
	InitialMarginFrac   float64 `json:"f"` // Initial margin fraction
	MarkPrice           float64 `json:"M"` // Mark price
	MaintenanceMargin   float64 `json:"m"` // Maintenance margin fraction
	NetQuantity         float64 `json:"q"` // Net quantity
	NetExposureQty      float64 `json:"Q"` // Net exposure quantity
	NetExposureNotional float64 `json:"n"` // Net exposure notional
	PositionID          string  `json:"i"` // Position ID
	PnLRealized         string  `json:"p"` // PnL realized
	PnLUnrealized       string  `json:"P"` // PnL unrealized
	EngineTimestamp     int64   `json:"T"` // Engine timestamp in microseconds
}

package trading

// RFQActiveEvent represents an active RFQ event
type RFQActiveEvent struct {
	EventType        string `json:"e"`  // Event type
	EventTime        int64  `json:"E"`  // Event time in microseconds
	RFQID            int64  `json:"R"`  // RFQ ID
	Symbol           string `json:"s"`  // Symbol
	Quantity         string `json:"q"`  // Quantity
	SubmissionTime   int64  `json:"w"`  // Submission time in milliseconds
	ExpiryTime       int64  `json:"W"`  // Expiry time in milliseconds
	RFQStatus        string `json:"X"`  // RFQ status
	EngineTimestamp  int64  `json:"T"`  // Engine timestamp in microseconds
}

// QuoteAcceptedEvent represents a quote accepted event
type QuoteAcceptedEvent struct {
	EventType        string `json:"e"`  // Event type
	EventTime        int64  `json:"E"`  // Event time in microseconds
	RFQID            int64  `json:"R"`  // RFQ ID
	QuoteID          int64  `json:"Q"`  // Quote ID
	ClientQuoteID    string `json:"C"`  // Client Quote ID (optional)
	QuoteStatus      string `json:"X"`  // Quote status
	EngineTimestamp  int64  `json:"T"`  // Engine timestamp in microseconds
}

// QuoteCancelledEvent represents a quote cancelled event
type QuoteCancelledEvent struct {
	EventType        string `json:"e"`  // Event type
	EventTime        int64  `json:"E"`  // Event time in microseconds
	RFQID            int64  `json:"R"`  // RFQ ID
	QuoteID          int64  `json:"Q"`  // Quote ID
	ClientQuoteID    string `json:"C"`  // Client Quote ID (optional)
	QuoteStatus      string `json:"X"`  // Quote status
	EngineTimestamp  int64  `json:"T"`  // Engine timestamp in microseconds
}

// RFQFilledEvent represents an RFQ filled event
type RFQFilledEvent struct {
	EventType        string `json:"e"`  // Event type
	EventTime        int64  `json:"E"`  // Event time in microseconds
	RFQID            int64  `json:"R"`  // RFQ ID
	QuoteID          int64  `json:"Q"`  // Quote ID
	ClientQuoteID    string `json:"C"`  // Client Quote ID (optional)
	Side             string `json:"S"`  // RFQ side (Bid or Ask)
	FillPrice        string `json:"p"`  // Fill price
	QuoteStatus      string `json:"X"`  // Quote status
	EngineTimestamp  int64  `json:"T"`  // Engine timestamp in microseconds
}
package authenticated

type FillHistoryResponse struct {
	ClientID        string `json:"clientId"`
	Fee             string `json:"fee"`
	FeeSymbol       string `json:"feeSymbol"`
	IsMaker         bool   `json:"isMaker"`
	OrderID         string `json:"orderId"`
	Price           string `json:"price"`
	Quantity        string `json:"quantity"`
	Side            string `json:"side"`
	Symbol          string `json:"symbol"`
	SystemOrderType string `json:"systemOrderType"`
	Timestamp       string `json:"timestamp"`
	TradeID         int    `json:"tradeId"`
}
type OrderHistoryResponse struct {
	ClientID        string `json:"clientId"`
	Fee             string `json:"fee"`
	FeeSymbol       string `json:"feeSymbol"`
	IsMaker         bool   `json:"isMaker"`
	OrderID         string `json:"orderId"`
	Price           string `json:"price"`
	Quantity        string `json:"quantity"`
	Side            string `json:"side"`
	Symbol          string `json:"symbol"`
	SystemOrderType string `json:"systemOrderType"`
	Timestamp       string `json:"timestamp"`
	TradeID         int    `json:"tradeId"`
}
type BorrowHistoryItem struct {
	EventType         string `json:"eventType"`
	PositionId        string `json:"positionId"`
	PositionQuantity  string `json:"positionQuantity"`
	Quantity          string `json:"quantity"`
	Source            string `json:"source"`
	Symbol            string `json:"symbol"`
	Timestamp         string `json:"timestamp"`
	SpotMarginOrderId string `json:"spotMarginOrderId"`
}

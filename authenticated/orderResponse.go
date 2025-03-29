package authenticated

type OrderResponse struct {
	OrderType              string `json:"orderType"`
	ID                     string `json:"id"`
	ClientID               int    `json:"clientId"`
	CreatedAt              int64  `json:"createdAt"`
	ExecutedQuantity       string `json:"executedQuantity"`
	ExecutedQuoteQuantity  string `json:"executedQuoteQuantity"`
	Quantity               string `json:"quantity"`
	QuoteQuantity          string `json:"quoteQuantity"`
	ReduceOnly             bool   `json:"reduceOnly"`
	TimeInForce            string `json:"timeInForce"`
	SelfTradePrevention    string `json:"selfTradePrevention"`
	Side                   string `json:"side"`
	Status                 string `json:"status"`
	StopLossTriggerPrice   string `json:"stopLossTriggerPrice"`
	StopLossLimitPrice     string `json:"stopLossLimitPrice"`
	StopLossTriggerBy      string `json:"stopLossTriggerBy"`
	Symbol                 string `json:"symbol"`
	TakeProfitTriggerPrice string `json:"takeProfitTriggerPrice"`
	TakeProfitLimitPrice   string `json:"takeProfitLimitPrice"`
	TakeProfitTriggerBy    string `json:"takeProfitTriggerBy"`
	TriggerBy              string `json:"triggerBy"`
	TriggerPrice           string `json:"triggerPrice"`
	TriggerQuantity        string `json:"triggerQuantity"`
	TriggeredAt            int64  `json:"triggeredAt"`
	RelatedOrderID         string `json:"relatedOrderId"`
}

type OpenOrdersResponse []OrderResponse

// OrderExecuteResponse 创建订单的响应结构体
type OrderExecuteResponse struct {
	OrderType              string `json:"orderType"`
	ID                     string `json:"id"`
	ClientID               int    `json:"clientId"`
	CreatedAt              int64  `json:"createdAt"`
	ExecutedQuantity       string `json:"executedQuantity"`
	ExecutedQuoteQuantity  string `json:"executedQuoteQuantity"`
	Quantity               string `json:"quantity"`
	QuoteQuantity          string `json:"quoteQuantity"`
	ReduceOnly             bool   `json:"reduceOnly"`
	TimeInForce            string `json:"timeInForce"`
	SelfTradePrevention    string `json:"selfTradePrevention"`
	Side                   string `json:"side"`
	Status                 string `json:"status"`
	StopLossTriggerPrice   string `json:"stopLossTriggerPrice"`
	StopLossLimitPrice     string `json:"stopLossLimitPrice"`
	StopLossTriggerBy      string `json:"stopLossTriggerBy"`
	Symbol                 string `json:"symbol"`
	TakeProfitTriggerPrice string `json:"takeProfitTriggerPrice"`
	TakeProfitLimitPrice   string `json:"takeProfitLimitPrice"`
	TakeProfitTriggerBy    string `json:"takeProfitTriggerBy"`
	TriggerBy              string `json:"triggerBy"`
	TriggerPrice           string `json:"triggerPrice"`
	TriggerQuantity        string `json:"triggerQuantity"`
	TriggeredAt            int64  `json:"triggeredAt"`
	RelatedOrderID         string `json:"relatedOrderId"`
}

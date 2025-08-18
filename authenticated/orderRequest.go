package authenticated

func NewExecuteOrderRequest(orderType, side, symbol string) *ExecuteOrderRequest {
	return &ExecuteOrderRequest{
		OrderType: orderType,
		Side:      side,
		Symbol:    symbol,
	}
}

func (r *ExecuteOrderRequest) WithMargin() {
	r.WithAutoLend(true)
	r.WithAutoBorrow(true)
	r.WithAutoLendRedeem(true)
	r.WithAutoBorrowRepay(true)
}
func (r *ExecuteOrderRequest) WithQuantity(quantity string) *ExecuteOrderRequest {
	r.Quantity = &quantity
	return r
}

func (r *ExecuteOrderRequest) WithPrice(price string) *ExecuteOrderRequest {
	r.Price = &price
	return r
}

func (r *ExecuteOrderRequest) WithTimeInForce(tif string) *ExecuteOrderRequest {
	r.TimeInForce = &tif
	return r
}

// ExecuteOrderRequest 表示订单请求
type ExecuteOrderRequest struct {
	// 必填字段
	OrderType string `json:"orderType"` // "Market" 或 "Limit"
	Side      string `json:"side"`      // "Bid" 或 "Ask"
	Symbol    string `json:"symbol"`    // 交易对

	// 可选字段
	AutoLend               *bool   `json:"autoLend,omitempty"`
	AutoLendRedeem         *bool   `json:"autoLendRedeem,omitempty"`
	AutoBorrow             *bool   `json:"autoBorrow,omitempty"`
	AutoBorrowRepay        *bool   `json:"autoBorrowRepay,omitempty"`
	ClientId               *uint32 `json:"clientId,omitempty"`
	PostOnly               *bool   `json:"postOnly,omitempty"`
	Price                  *string `json:"price,omitempty"`
	Quantity               *string `json:"quantity,omitempty"`
	QuoteQuantity          *string `json:"quoteQuantity,omitempty"`
	ReduceOnly             *bool   `json:"reduceOnly,omitempty"`
	SelfTradePrevention    *string `json:"selfTradePrevention,omitempty"`
	StopLossLimitPrice     *string `json:"stopLossLimitPrice,omitempty"`
	StopLossTriggerPrice   *string `json:"stopLossTriggerPrice,omitempty"`
	TakeProfitLimitPrice   *string `json:"takeProfitLimitPrice,omitempty"`
	TakeProfitTriggerPrice *string `json:"takeProfitTriggerPrice,omitempty"`
	TimeInForce            *string `json:"timeInForce,omitempty"`
	TriggerPrice           *string `json:"triggerPrice,omitempty"`
	TriggerQuantity        *string `json:"triggerQuantity,omitempty"`
	MarketType             *string `json:"marketType,omitempty"`
}

func (r *ExecuteOrderRequest) WithMarketType(marketType string) *ExecuteOrderRequest {
	r.MarketType = &marketType
	return r
}

func (r *ExecuteOrderRequest) WithBoolOption(value bool, field **bool) *ExecuteOrderRequest {
	*field = &value
	return r
}

func (r *ExecuteOrderRequest) WithAutoLend(value bool) *ExecuteOrderRequest {
	return r.WithBoolOption(value, &r.AutoLend)
}
func (r *ExecuteOrderRequest) WithPostOnly(value bool) *ExecuteOrderRequest {
	return r.WithBoolOption(value, &r.PostOnly)
}
func (r *ExecuteOrderRequest) WithAutoLendRedeem(value bool) *ExecuteOrderRequest {
	return r.WithBoolOption(value, &r.AutoLendRedeem)
}
func (r *ExecuteOrderRequest) WithAutoBorrow(value bool) *ExecuteOrderRequest {
	return r.WithBoolOption(value, &r.AutoBorrow)
}
func (r *ExecuteOrderRequest) WithAutoBorrowRepay(value bool) *ExecuteOrderRequest {
	return r.WithBoolOption(value, &r.AutoBorrowRepay)
}
func (r *ExecuteOrderRequest) WithReduceOnly(value bool) *ExecuteOrderRequest {
	return r.WithBoolOption(value, &r.ReduceOnly)
}

func (r *ExecuteOrderRequest) ToURLQueryString() string {
	p := NewParams()

	// Add fields in alphabetical order, ignoring nil pointers and empty strings
	if r.AutoBorrow != nil && *r.AutoBorrow {
		p.AddBoolean("autoBorrow", r.AutoBorrow)
	}

	if r.AutoBorrowRepay != nil && *r.AutoBorrowRepay {
		p.AddBoolean("autoBorrowRepay", r.AutoBorrowRepay)
	}

	if r.AutoLend != nil && *r.AutoLendRedeem {
		p.AddBoolean("autoLend", r.AutoLend)
	}

	if r.AutoLendRedeem != nil && *r.AutoLendRedeem {
		p.AddBoolean("autoLendRedeem", r.AutoLendRedeem)
	}

	if r.ClientId != nil {
		p.AddUint32("clientId", r.ClientId)
	}

	if r.MarketType != nil && *r.MarketType != "" {
		p.Add("marketType", r.MarketType)
	}

	// Required field, but check if not empty

	p.Add("orderType", &r.OrderType)

	if r.PostOnly != nil {
		p.AddBoolean("postOnly", r.PostOnly)
	}

	if r.Price != nil && *r.Price != "" {
		p.Add("price", r.Price)
	}

	if r.Quantity != nil && *r.Quantity != "" {
		p.Add("quantity", r.Quantity)
	}

	if r.QuoteQuantity != nil && *r.QuoteQuantity != "" {
		p.Add("quoteQuantity", r.QuoteQuantity)
	}

	if r.ReduceOnly != nil && *r.ReduceOnly {
		p.AddBoolean("reduceOnly", r.ReduceOnly)
	}

	if r.SelfTradePrevention != nil && *r.SelfTradePrevention != "" {
		p.Add("selfTradePrevention", r.SelfTradePrevention)
	}

	// Required field, but check if not empty

	p.Add("side", &r.Side)

	if r.StopLossLimitPrice != nil && *r.StopLossLimitPrice != "" {
		p.Add("stopLossLimitPrice", r.StopLossLimitPrice)
	}

	if r.StopLossTriggerPrice != nil && *r.StopLossTriggerPrice != "" {
		p.Add("stopLossTriggerPrice", r.StopLossTriggerPrice)
	}

	// Required field, but check if not empty

	p.Add("symbol", &r.Symbol)

	if r.TakeProfitLimitPrice != nil && *r.TakeProfitLimitPrice != "" {
		p.Add("takeProfitLimitPrice", r.TakeProfitLimitPrice)
	}

	if r.TakeProfitTriggerPrice != nil && *r.TakeProfitTriggerPrice != "" {
		p.Add("takeProfitTriggerPrice", r.TakeProfitTriggerPrice)
	}

	if r.TimeInForce != nil && *r.TimeInForce != "" {
		p.Add("timeInForce", r.TimeInForce)
	}

	if r.TriggerPrice != nil && *r.TriggerPrice != "" {
		p.Add("triggerPrice", r.TriggerPrice)
	}

	if r.TriggerQuantity != nil && *r.TriggerQuantity != "" {
		p.Add("triggerQuantity", r.TriggerQuantity)
	}

	return p.String()
}

type OpenOrderRequest struct {
	Symbol string `json:"symbol"` // 交易对

	ClientId *uint32 `json:"clientId,omitempty"`
	OrderId  *string `json:"orderId,omitempty"`
}

func NewOpenOrderRequest(symbol string) *OpenOrderRequest {
	return &OpenOrderRequest{Symbol: symbol}
}
func (r *OpenOrderRequest) WithClientId(clientId uint32) *OpenOrderRequest {
	r.ClientId = &clientId
	return r
}
func (r *OpenOrderRequest) WithOrderId(orderId string) *OpenOrderRequest {
	r.OrderId = &orderId
	return r
}

func (r *OpenOrderRequest) ToURLQueryString() string {
	p := NewParams()
	p.AddUint32("clientId", r.ClientId)
	p.Add("orderId", r.OrderId)
	p.Add("symbol", &r.Symbol)
	return p.String()

}

type CancelOrdersRequest struct {
	Symbol    string `json:"symbol"`    // 交易对
	OrderType string `json:"orderType"` // 交易对
}

func NewCancelOrdersRequest(symbol string, orderType CancelOrderType) *CancelOrdersRequest {
	return &CancelOrdersRequest{Symbol: symbol, OrderType: string(orderType)}
}

func (r *CancelOrdersRequest) ToURLQueryString() string {
	p := NewParams()
	p.Add("orderType", &r.OrderType)
	p.Add("symbol", &r.Symbol)
	return p.String()

}

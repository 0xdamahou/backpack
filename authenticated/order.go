package authenticated

import (
	//"encoding/json"
	json "github.com/bytedance/sonic"
)

// OrderSide 定义订单方向
type OrderSide string

const (
	OrderSideBuy  OrderSide = "Bid"
	OrderSideSell OrderSide = "Ask"
	POST          string    = "POST"
	GET           string    = "GET"
	PATCH         string    = "PATCH"
	DELETE        string    = "DELETE"
)

// OrderType 定义订单类型
type OrderType string

const (
	OrderTypeLimit    OrderType = "Limit"
	OrderTypeMarket   OrderType = "Market"
	OrderTypePostOnly OrderType = "PostOnly"
	OrderTypeTrigger  OrderType = "Trigger"
)

// TimeInForce 定义订单有效时间
type TimeInForce string

const (
	TimeInForceGTC TimeInForce = "GTC" // Good 'Til Canceled
	TimeInForceIOC TimeInForce = "IOC" // Immediate or Cancel
	TimeInForceFOK TimeInForce = "FOK" // Fill or Kill
)

// SelfTradePrevention 定义自成交预防策略
type SelfTradePrevention string

const (
	SelfTradePreventionRejectTaker SelfTradePrevention = "RejectTaker"
)

// CancelOrderType 定义取消订单的类型（用于 CancelOpenOrders）
type CancelOrderType string

const (
	CancelOrderTypeRestingLimit CancelOrderType = "RestingLimitOrder"
	CancelOrderTypeConditional  CancelOrderType = "ConditionalOrder"
)

// MarketType 定义市场类型（用于 GetOpenOrders）
func (c *BackpackClient) LimitOrder(symbol string, side OrderSide, quantity string, price string) (OrderExecuteResponse, error) {
	or := NewExecuteOrderRequest(string(OrderTypeLimit), string(side), symbol)
	or.WithPrice(price).WithQuantity(quantity)
	return c.ExecuteOrder(or)
}
func (c *BackpackClient) MarginLimitOrder(symbol string, side OrderSide, quantity string, price string) (OrderExecuteResponse, error) {
	oor := NewExecuteOrderRequest(string(OrderTypeLimit), string(side), symbol)
	oor.WithPrice(price)
	oor.WithQuantity(quantity)
	oor.WithAutoLend(true)
	oor.WithAutoBorrow(true)
	oor.WithAutoLendRedeem(true)
	oor.WithAutoBorrowRepay(true)
	return c.ExecuteOrder(oor)
}
func (c *BackpackClient) MarketOrder(symbol string, side OrderSide, quantity string) (OrderExecuteResponse, error) {
	or := NewExecuteOrderRequest(string(OrderTypeMarket), string(side), symbol)
	or.WithQuantity(quantity)
	return c.ExecuteOrder(or)
}

// ExecuteOrder 创建订单
func (c *BackpackClient) ExecuteOrder(order *ExecuteOrderRequest) (OrderExecuteResponse, error) {
	var resp OrderExecuteResponse
	instruction := "orderExecute"
	endPoint := "api/v1/order"
	jsonBody, err := json.Marshal(order)
	if err != nil {
		return resp, err
	}

	//log.Printf("Request: %s", string(jsonBody))
	err = c.DoPost(endPoint, instruction, jsonBody, &resp)
	return resp, err
}

// GetOpenOrder 查询单个未结订单 (/order/open)
func (c *BackpackClient) GetOpenOrder(oor *OpenOrderRequest) (OrderResponse, error) {
	var resp OrderResponse
	instruction := "orderQuery"
	endPoint := "api/v1/order"
	err := c.DoGet(endPoint, instruction, oor.ToURLQueryString(), &resp)
	return resp, err
}

// GetOpenOrders 查询所有未结订单 (/orders/open)
func (c *BackpackClient) GetOpenOrders(marketType MarketType, symbol string) (OpenOrdersResponse, error) {
	var resp OpenOrdersResponse
	instruction := "orderQueryAll"
	endPoint := "api/v1/orders"
	params := NewParams()
	mk := string(marketType)
	params.Add("marketType", &mk).Add("symbol", &symbol)
	err := c.DoGet(endPoint, instruction, params.String(), &resp)
	return resp, err
}

// CancelOpenOrder 取消单个未结订单 (/order/open/cancel)
func (c *BackpackClient) CancelOpenOrder(oor *OpenOrderRequest) (OrderResponse, error) {
	var resp OrderResponse
	instruction := "orderCancel"
	endPoint := "api/v1/order"
	body, err := json.Marshal(*oor)
	err = c.DoDelete(endPoint, instruction, body, &resp)
	return resp, err

}

// CancelOpenOrders 取消所有未结订单 (/orders/open/cancel)
func (c *BackpackClient) CancelOpenOrders(symbol string, orderType CancelOrderType) (OpenOrdersResponse, error) {
	var resp OpenOrdersResponse
	instruction := "orderCancelAll"
	endPoint := "api/v1/orders"
	params := map[string]interface{}{
		"symbol":    symbol,
		"orderType": string(orderType),
	}
	jsonBody, err := json.Marshal(params)
	if err != nil {
		return resp, err
	}
	err = c.DoDelete(endPoint, instruction, jsonBody, &resp)
	return resp, err
}

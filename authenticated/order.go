package authenticated

import (
	//"encoding/json"

	"bytes"

	"log"
	"strings"

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
	var buf bytes.Buffer
	err := json.ConfigDefault.NewEncoder(&buf).Encode(order)
	if err != nil {
		return resp, err
	}

	//log.Printf("Request: %s", string(jsonBody))
	//q := Body2Query(jsonBody)
	err = c.DoPost(endPoint, instruction, &buf, order.ToURLQueryString(), &resp)
	return resp, err
}

// ExecuteOrders 批量订单
func (c *BackpackClient) ExecuteOrders(orders []*ExecuteOrderRequest) ([]OrderExecuteResponse, error) {
	var resp []OrderExecuteResponse
	instruction := "orderExecute"
	endPoint := "api/v1/orders"
	//jsonBody, err := json.Marshal(orders)
	//if err != nil {
	//	return resp, err
	//}

	var queryBuilder strings.Builder
	queryBuilder.Grow(256 * len(orders))

	for i, order := range orders {
		if i > 0 {
			queryBuilder.WriteString("&instruction=orderExecute&")
		}
		queryBuilder.WriteString(order.ToURLQueryString())
	}
	q := queryBuilder.String()
	var buf bytes.Buffer
	err := json.ConfigDefault.NewEncoder(&buf).Encode(orders)
	if err != nil {
		log.Println(err)
		return resp, err
	}

	//log.Printf("[%s] %s\n", time.Now().Format("15:04:05.000"), q)
	err = c.DoPost(endPoint, instruction, &buf, q, &resp)
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
	var buf bytes.Buffer
	err := json.ConfigDefault.NewEncoder(&buf).Encode(*oor)
	err = c.DoDelete(endPoint, instruction, &buf, oor.ToURLQueryString(), &resp)
	return resp, err
}

// CancelOpenOrders 取消所有未结订单 (/orders/open/cancel)
func (c *BackpackClient) CancelOpenOrders(symbol string, orderType CancelOrderType) (OpenOrdersResponse, error) {
	var resp OpenOrdersResponse
	instruction := "orderCancelAll"
	endPoint := "api/v1/orders"
	params := NewCancelOrdersRequest(symbol, orderType)
	var buf bytes.Buffer
	err := json.ConfigDefault.NewEncoder(&buf).Encode(*params)
	if err != nil {
		return resp, err
	}
	err = c.DoDelete(endPoint, instruction, &buf, params.ToURLQueryString(), &resp)
	return resp, err
}

package authenticated

import (
	"log"
	"testing"
)

func TestBackpackClient_GetOpenOrders(t *testing.T) {
	openOrders, err := client.GetOpenOrders(MarketTypeSpot, "SOL_USDC")
	if err != nil {
		log.Println(err)
		return
	}
	for _, item := range openOrders {
		log.Printf("%+v\n", item)
	}
}

func TestBackpackClient_GetOpenOrder(t *testing.T) {
	oor := NewOpenOrderRequest("SOL_USDC")
	oor.WithOrderId("114217259366744064")
	order, err := client.GetOpenOrder(oor)

	if err != nil {
		return
	}
	log.Printf("%+v\n", order)
}

func TestBackpackClient_LimitOrder(t *testing.T) {
	order, err := client.LimitOrder("SOL_USDC_PERP", OrderSideSell, "3", "144.5")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", order)
}
func TestBackpackClient_ExecuteOrder(t *testing.T) {
	or := NewExecuteOrderRequest(string(OrderTypeLimit), string(OrderSideBuy), "SOL_USDC_PERP")
	or.WithPrice("143").WithQuantity("3").WithReduceOnly(true)

	order, err := client.ExecuteOrder(or)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n", order)
}

func TestBackpackClient_CancelOpenOrders(t *testing.T) {
	orders, err := client.CancelOpenOrders("SOL_USDC_PERP", CancelOrderTypeRestingLimit)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", orders)
}
func TestBackpackClient_CancelOpenOrder(t *testing.T) {
	symbol := "SOL_USDC_PERP"
	oor := NewOpenOrderRequest(symbol)
	oor.WithOrderId("114217259366744064")
	orders, err := client.CancelOpenOrder(oor)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", orders)
}

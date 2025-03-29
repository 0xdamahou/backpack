package main

import (
	"fmt"
	"github.com/0xdamahou/backpack/authenticated"
)

func main() {
	client, err := authenticated.NewBackpackClient("", "")
	if err != nil {
		return
	}
	response, err := client.LimitOrder("SOL_USDC", authenticated.OrderSideSell, "0.2", "150")
	if err != nil {
		fmt.Printf("Failed to limit order: %v\n", err)
		return
	}
	fmt.Printf("Limit Order Response: %+v\n", response)
	orders, err := client.GetOpenOrders(authenticated.MarketTypeSpot, "SOL_USDC")
	if err != nil {
		fmt.Printf("Failed to get open orders: %v\n", err)
		return
	}
	fmt.Printf("Orders: %+v\n", orders)
	oor := authenticated.NewOpenOrderRequest("SOL_USDC").WithOrderId("114200605223419906")

	oResp, err := client.GetOpenOrder(oor)
	if err != nil {
		fmt.Printf("Failed to get oor: %v\n", err)
	} else {
		fmt.Printf("GetOpenOrder: %+v\n", oResp)
	}

}

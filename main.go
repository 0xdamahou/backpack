package main

import (
	"fmt"
	"github.com/0xdamahou/backpack/authenticated"
	"log"
	"time"
)

const (
	apiKey     = "" // 替换为你的 API Key
	privateKey = "" // 替换为你的 Private Key
)

func TryOrderRelated(client *authenticated.BackpackClient) {
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
	oor := authenticated.NewOpenOrderRequest("SOL_USDC")
	oor.WithOrderId("114200605223419906")
	oResp, err := client.GetOpenOrder(oor)
	if err != nil {
		fmt.Printf("Failed to get oor: %v\n", err)
	} else {
		fmt.Printf("GetOpenOrder: %+v\n", oResp)
	}

	order, err := client.CancelOpenOrders("SOL_USDC", authenticated.CancelOrderTypeRestingLimit)
	if err != nil {
		fmt.Printf("Failed to cancel oor: %v\n", err)
		return
	}
	fmt.Printf("CancelOpenOrders: %+v\n", order)
}
func TryCapital(client *authenticated.BackpackClient) {
	//
	balances, err := client.GetBalances()
	if err != nil {
		fmt.Printf("Failed to get balances: %v\n", err)
		return
	}
	////114200605223419906
	fmt.Println("Balances:", balances)

	to := time.Now()
	from := to.AddDate(0, 0, -1)
	deposits, e := client.GetDeposit(from.UnixMilli(), to.UnixMilli(), 100, 0)
	if e != nil {
		fmt.Printf("Failed to get deposits: %v\n", e)
		//return
	}
	fmt.Printf("Deposits: %+v\n", deposits)

	address, e := client.GetDepositAddress("Solana")
	if e != nil {
		fmt.Printf("Failed to get deposit address: %v\n", e)

	}
	fmt.Printf("Deposit address: %s\n", address)
	withdrawals, err := client.GetWithdrawal(from.UnixMilli(), to.UnixMilli(), 100, 0)
	if err != nil {
		fmt.Printf("Failed to get withdrawals: %v\n", err)
		return
	}
	fmt.Printf("Withdrawals: %+v\n", withdrawals)

}

func TryAccount(client *authenticated.BackpackClient) {
	account, err := client.GetAccount()
	if err != nil {
		return
	}
	fmt.Printf("Account: %+v\n", account)
	quantity, err := client.GetMaxBorrowQuantity("BTC")
	if err != nil {
		return
	}
	fmt.Printf("Max Borrow Quantity: %+v\n", quantity)
	mor := authenticated.NewMaxOrderQuantityRequest("SOL_USDC", authenticated.OrderSideBuy)
	mor.WithPrice("150")
	maxOrderQuantity, err := client.GetMaxOrderQuantity(mor)
	if err != nil {
		log.Printf("Failed to get max order quantity: %v\n", err)
		return
	}
	fmt.Printf("Max Order Quantity: %+v\n", maxOrderQuantity)
	mwr := authenticated.NewMaxWithdrawalQuantityRequest("USDC")
	//mwr.WithAutoBorrow(true)
	maxWithdrawalQuantity, err := client.GetMaxWithdrawalQuantity(mwr)
	if err != nil {
		log.Printf("Failed to get max withdrawal quantity: %v\n", err)
		return
	}
	fmt.Printf("Max Withdrawal Quantity: %+v\n", maxWithdrawalQuantity)
}
func TryHistory(client *authenticated.BackpackClient) {
	//fhq := backpack.FillHistoryRequest{Symbol: "SOL_USDC"}
	symbol := "SOL_USDC"
	fhq := authenticated.FillHistoryRequest{Symbol: &symbol}
	fillHistory, err := client.GetFillHistory(&fhq)
	if err != nil {
		log.Printf("Failed to get fill history: %v\n", err)
		return
	}
	log.Printf("%+v\n", fillHistory)
	ohq := authenticated.OrderHistoryRequest{Symbol: &symbol}
	history, err := client.GetOrderHistory(&ohq)
	if err != nil {
		log.Printf("Failed to get order history: %v\n", err)
		return
	}
	fmt.Printf("%+v\n", history)
}

func main() {
	client, err := authenticated.NewBackpackClient(apiKey, privateKey)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}
	TryAccount(client)
	//TryOrderRelated(client)
	//history, err := client.GetOrderHistory(&backpack.OrderHistoryRequest{})
	//
	//if err != nil {
	//	fmt.Printf("Failed to get order history: %v\n", err)
	//	return
	//}
	//fmt.Printf("%+v\n", history)

}

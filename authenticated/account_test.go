package authenticated

import (
	"log"
	"testing"
)

const (
	apiKey     = "" // 替换为你的 API Key
	privateKey = "" // 替换为你的 Private Key
)

// The client is for test
var client, _ = NewBackpackClient(apiKey, privateKey)

func TestBackpackClient_GetAccount(t *testing.T) {
	account, err := client.GetAccount()
	if err != nil {
		log.Printf("Failed to get account: %v\n", err)
	}
	log.Printf("Account: %+v\n", account)
}

func TestBackpackClient_UpdateAccount(t *testing.T) {
	err := client.UpdateAccount(true, true, true, "10")
	if err != nil {
		log.Printf("Failed to update account: %v\n", err)
		return
	}

}

func TestBackpackClient_GetMaxBorrowQuantity(t *testing.T) {
	symbol := "SOL"
	quantity, err := client.GetMaxBorrowQuantity(symbol)
	if err != nil {
		log.Printf("Failed to get max borrow quantity: %v\n", err)
		return
	}
	log.Printf("Max borrow quantity: %+v\n", quantity)
}

func TestBackpackClient_GetMaxOrderQuantity(t *testing.T) {
	symbol := "SOL_USDC"
	moqr := NewMaxOrderQuantityRequest(symbol, OrderSideBuy)
	moqr.WithPrice("142")
	moqr.WithAutoBorrow(true)
	quantity, err := client.GetMaxOrderQuantity(moqr)
	if err != nil {
		log.Printf("Failed to get max order quantity: %v\n", err)
		return
	}
	log.Printf("Max order quantity: %+v\n", quantity)
}
func TestBackpackClient_GetMaxWithdrawalQuantity(t *testing.T) {
	symbol := "SOL"
	mwq := NewMaxWithdrawalQuantityRequest(symbol)
	mwq.WithAutoBorrow(true)
	quantity, err := client.GetMaxWithdrawalQuantity(mwq)
	if err != nil {
		log.Printf("Failed to get max withdrawal quantity: %v\n", err)
		return
	}
	log.Printf("Max withdrawal quantity: %+v\n", quantity)
}

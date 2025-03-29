package authenticated

import (
	"log"
	"testing"
	"time"
)

func TestBackpackClient_GetBalances(t *testing.T) {
	balances, err := client.GetBalances()
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Println(balances)
}

func TestBackpackClient_GetDeposit(t *testing.T) {
	to := time.Now().UnixMilli()
	from := time.Now().Add(-30 * 24 * time.Hour).UnixMilli()
	deposits, err := client.GetDeposit(from, to, 100, 0)
	if err != nil {
		return
	}
	for _, deposit := range deposits {
		log.Println(deposit)
	}
}

func TestBackpackClient_GetDepositAddress(t *testing.T) {
	address, err := client.GetDepositAddress("Solana")
	if err != nil {
		return
	}
	log.Println(address)
}

func TestBackpackClient_GetWithdrawals(t *testing.T) {
	to := time.Now().UnixMilli()
	from := time.Now().Add(-30 * 24 * time.Hour).UnixMilli()
	withdrawal, err := client.GetWithdrawal(from, to, 100, 0)
	if err != nil {
		return
	}
	log.Println(withdrawal)
}

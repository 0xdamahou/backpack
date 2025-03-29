package authenticated

import (
	"log"
	"testing"
	"time"
)

func TestBackpackClient_GetFillHistory(t *testing.T) {
	to := time.Now()
	from := to.AddDate(0, 0, -3)
	fromTime := from.UnixMilli()
	toTime := to.UnixMilli()
	fhq := &FillHistoryRequest{From: &fromTime, To: &toTime}
	history, err := client.GetFillHistory(fhq)
	if err != nil {
		log.Println(err)
		return
	}
	for _, historyItem := range history {
		log.Println(historyItem)
	}
	//log.Println(history)
}

func TestBackpackClient_GetOrderHistory(t *testing.T) {
	symbol := "SOL_USDC"
	query := &OrderHistoryRequest{Symbol: &symbol, MarketType: []string{string(MarketTypeSpot)}}
	history, err := client.GetOrderHistory(query)
	if err != nil {
		log.Println(err)
		return
	}
	for _, historyItem := range history {
		log.Printf("%+v", historyItem)
	}
}

func TestBackpackClient_GetBorrowHistory(t *testing.T) {
	//et := LendType
	usdc := "USDC"
	blhr := &BorrowLendHistoryRequest{}
	blhr.Symbol = &usdc

	history, err := client.GetBorrowHistory(blhr)
	if err != nil {
		log.Println(err)
		return
	}
	for _, item := range history {
		log.Printf("%+v\n", item)
	}
}

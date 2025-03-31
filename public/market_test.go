package public

import (
	"log"
	"strings"
	"testing"
	"time"
)

var bpbc = NewBackpackPublicClient()

func TestBackpackPublicClient_GetMarkets(t *testing.T) {
	markets, err := bpbc.GetMarkets()
	if err != nil {
		t.Fatal(err)
		return
	}
	//log.Println(len(markets))
	mks := make([]MarketSymbol, 0)
	for _, market := range markets {
		if strings.HasSuffix(market.Symbol, "PERP") {
			mks = append(mks, market)
		}
	}
	log.Println(len(mks))
	for _, mk := range mks {
		log.Printf("%+v", mk)
	}

}

func TestBackpackPublicClient_GetMarket(t *testing.T) {
	symbol := "SOL_USDC_PERP"
	market, err := bpbc.GetMarket(symbol)
	if err != nil {
		t.Fatal(err)
		return
	}
	if market.Symbol != symbol {
		t.Fail()
	}
	if market.MarketType != "PERP" {
		t.Fail()
	}
	log.Printf("%+v\n", market.Filters)
}

func TestBackpackPublicClient_GetTicker(t *testing.T) {
	symbol := "BTC_USDC_PERP"
	ticker, err := bpbc.GetTicker(symbol, "1w")
	if err != nil {
		log.Fatal(err)
		return
	}
	if ticker.Symbol != symbol {
		t.Fail()
	}
}

func TestBackpackPublicClient_GetTickers(t *testing.T) {
	tickers, err := bpbc.GetTickers("1d")
	if err != nil {
		t.Fatal(err)
		return
	}
	//if len(tickers) > 50 {
	//	return
	//
	//}
	log.Println(len(tickers))
	for _, ticker := range tickers {
		log.Printf("%+v\n", ticker)
	}
}

func TestBackpackPublicClient_GetDepth(t *testing.T) {
	symbol := "USDT_USDC"
	depth, err := bpbc.GetDepth(symbol)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Printf("%+v\n", depth)
	log.Printf("Asks: %+v Bids %+v\n", depth.Asks[0], depth.Bids[len(depth.Bids)-1])
}

func TestBackpackPublicClient_GetKline(t *testing.T) {
	symbol := "BTC_USDC"
	to := time.Now()
	from := to.AddDate(-1, 0, 0)
	kline, err := bpbc.GetKline(symbol, "1d", from.Unix(), to.Unix())
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Println(len(kline))

}

func TestBackpackPublicClient_GetAllMarketPrices(t *testing.T) {
	symbol := "BTC_USDC_PERP"
	prices, err := bpbc.GetAllMarkPrices(symbol)
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Printf("%+v\n", prices)
	prices, err = bpbc.GetAllMarkPrices("")
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Println(len(prices))
}

func TestBackpackPublicClient_GetOpenInterest(t *testing.T) {
	symbol := "BTC_USDC_PERP"
	interest, err := bpbc.GetOpenInterest(symbol)
	if err != nil {
		t.Fatal(err)
		return

	}
	log.Printf("%+v\n", interest)
	interest, err = bpbc.GetOpenInterest("")
	if err != nil {
		t.Fatal(err)
		return

	}
	log.Printf("%+v\n", len(interest))
}

func TestBackpackPublicClient_GetFundingIntervalRates(t *testing.T) {
	symbol := "BTC_USDC_PERP"
	rates, err := bpbc.GetFundingIntervalRates(symbol, 100, 0)
	if err != nil {
		t.Fatal(err)
		return
	}
	for _, rate := range rates {
		log.Printf("%+v\n", rate)
	}
	//log.Printf("%+v\n", rates)

}

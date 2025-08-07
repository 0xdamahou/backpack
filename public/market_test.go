package public

import (
	"log"
	"strconv"
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
	symbol := "FARTCOIN_USDC_PERP"
	to := time.Now()
	from := to.AddDate(-1, -2, -10)
	klines, err := bpbc.GetKline(symbol, "1d", from.Unix(), to.Unix())
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Println(len(klines))
	vol := 0.0
	num := 0.0
	vol2H := 0.0
	for i, kline := range klines {
		//log.Printf("%+v\n", kline)
		high, e := strconv.ParseFloat(kline.High, 64)
		if e != nil {
			continue
		}
		low, e := strconv.ParseFloat(kline.Low, 64)
		if e != nil {
			continue
		}
		vol += 2 * (high - low) / (high + low)
		num += 1
		if i > len(klines)-3 {
			c := 2 * (high - low) / (high + low)
			if c > vol2H {
				vol2H = c
			}
		}
	}
	vol240hAvg := vol / num
	Weight := 1.0
	switch {
	case vol2H > vol240hAvg*1.5:
		Weight = 2
	case vol2H < vol240hAvg*0.7:
		Weight = 0.5
	default:
		Weight = 1

	}
	log.Println(Weight, vol240hAvg, vol2H)

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

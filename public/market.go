package public

import (
	"net/url"
	"strconv"
)

func (bbc *BackpackPublicClient) GetMarkets() ([]MarketSymbol, error) {
	endpoint := "api/v1/markets"
	var markets []MarketSymbol
	err := bbc.DoGet(endpoint, "", &markets)
	return markets, err
}

func (bbc *BackpackPublicClient) GetMarket(symbol string) (MarketSymbol, error) {
	endpoint := "api/v1/market"
	var market MarketSymbol
	q := url.Values{}
	q.Add("symbol", symbol)
	err := bbc.DoGet(endpoint, q.Encode(), &market)
	return market, err
}

// GetTicker interval Enum: "1d" "1w"
func (bbc *BackpackPublicClient) GetTicker(symbol string, interval string) (Ticker, error) {
	endpoint := "api/v1/ticker"
	var ticker Ticker
	q := url.Values{}
	q.Add("symbol", symbol)
	q.Add("interval", interval)
	err := bbc.DoGet(endpoint, q.Encode(), &ticker)
	return ticker, err
}

func (bbc *BackpackPublicClient) GetTickers(interval string) ([]Ticker, error) {
	endpoint := "api/v1/tickers"
	var tickers []Ticker
	q := url.Values{}
	q.Add("interval", interval)
	err := bbc.DoGet(endpoint, q.Encode(), &tickers)
	return tickers, err
}

func (bbc *BackpackPublicClient) GetDepth(symbol string) (Depth, error) {
	endpoint := "api/v1/depth"
	var depth Depth
	q := url.Values{}
	q.Add("symbol", symbol)
	err := bbc.DoGet(endpoint, q.Encode(), &depth)
	return depth, err

}

// GetKline Enum: "1m" "3m" "5m" "15m" "30m" "1h" "2h" "4h" "6h" "8h" "12h" "1d" "3d" "1w" "1month"
func (bbc *BackpackPublicClient) GetKline(symbol string, interval string, startTime int64, endTime int64) ([]Kline, error) {
	endpoint := "api/v1/klines"
	var klines []Kline
	q := url.Values{}
	q.Add("symbol", symbol)
	q.Add("interval", interval)
	q.Add("startTime", strconv.FormatInt(startTime, 10))
	q.Add("endTime", strconv.FormatInt(endTime, 10))
	err := bbc.DoGet(endpoint, q.Encode(), &klines)
	return klines, err
}

func (bbc *BackpackPublicClient) GetAllMarkPrices(symbol string) ([]FundingInfo, error) {
	endpoint := "api/v1/markPrices"
	var query string
	if symbol != "" {
		q := url.Values{}
		q.Add("symbol", symbol)
		query = q.Encode()
	}
	var fundings []FundingInfo
	err := bbc.DoGet(endpoint, query, &fundings)
	return fundings, err
}

func (bbc *BackpackPublicClient) GetOpenInterest(symbol string) ([]OpenInterest, error) {
	endpoint := "api/v1/openInterest"
	var query string
	if symbol != "" {
		q := url.Values{}
		q.Add("symbol", symbol)
		query = q.Encode()
	}
	var interests []OpenInterest
	err := bbc.DoGet(endpoint, query, &interests)
	return interests, err
}

func (bbc *BackpackPublicClient) GetFundingIntervalRates(symbol string, limit uint64, offset uint64) ([]FundingRateHistory, error) {
	endpoint := "api/v1/fundingRates"
	var rates []FundingRateHistory
	q := url.Values{}
	q.Add("symbol", symbol)
	q.Add("limit", strconv.FormatUint(limit, 10))
	q.Add("offset", strconv.FormatUint(offset, 10))
	err := bbc.DoGet(endpoint, q.Encode(), &rates)
	return rates, err
}

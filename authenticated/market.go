package authenticated

import (
	"time"
)

// MarketType 定义市场类型
type MarketType string

const (
	MarketTypeSpot       MarketType = "SPOT"
	MarketTypePerp       MarketType = "PERP"
	MarketTypeIperp      MarketType = "IPERP"
	MarketTypeDated      MarketType = "DATED"
	MarketTypePrediction MarketType = "PREDICTION"
	MarketTypeRFQ        MarketType = "RFQ"
)

// MarketStatus 定义市场状态
type MarketStatus string

const (
	MarketStatusOnline  MarketStatus = "Online"
	MarketStatusOffline MarketStatus = "Offline"
)

// KlineInterval 定义 K 线时间间隔
type KlineInterval string

const (
	KlineInterval1m  KlineInterval = "1m"
	KlineInterval5m  KlineInterval = "5m"
	KlineInterval15m KlineInterval = "15m"
	KlineInterval1h  KlineInterval = "1h"
	KlineInterval4h  KlineInterval = "4h"
	KlineInterval1d  KlineInterval = "1d"
)

// MarketResponse 市场信息响应结构体（用于 GetMarkets 和 GetMarket）
type MarketResponse struct {
	Symbol                 string       `json:"symbol"`
	BaseAsset              string       `json:"baseAsset"`
	QuoteAsset             string       `json:"quoteAsset"`
	Type                   MarketType   `json:"type"`
	Status                 MarketStatus `json:"status"`
	MinOrderSize           string       `json:"minOrderSize"`
	MaxOrderSize           string       `json:"maxOrderSize"`
	TickSize               string       `json:"tickSize"`
	MinLeverage            int          `json:"minLeverage"`
	MaxLeverage            int          `json:"maxLeverage"`
	MarginRequirement      string       `json:"marginRequirement"`
	FundingRate            string       `json:"fundingRate,omitempty"`
	NextFundingTime        time.Time    `json:"nextFundingTime,omitempty"`
	ContractSize           string       `json:"contractSize,omitempty"`
	ExpirationTimestamp    time.Time    `json:"expirationTimestamp,omitempty"`
	PredictionMarketStatus string       `json:"predictionMarketStatus,omitempty"`
}

// TickerResponse 行情数据响应结构体（用于 GetTickers 和 GetTicker）
type TickerResponse struct {
	Symbol         string    `json:"symbol"`
	LastPrice      string    `json:"lastPrice"`
	BestBidPrice   string    `json:"bestBidPrice"`
	BestBidQty     string    `json:"bestBidQty"`
	BestAskPrice   string    `json:"bestAskPrice"`
	BestAskQty     string    `json:"bestAskQty"`
	Volume24h      string    `json:"volume24h"`
	QuoteVolume24h string    `json:"quoteVolume24h"`
	High24h        string    `json:"high24h"`
	Low24h         string    `json:"low24h"`
	Open24h        string    `json:"open24h"`
	Timestamp      time.Time `json:"timestamp"`
}

// KlineResponse K 线数据响应结构体（用于 GetKline）
type KlineResponse struct {
	OpenTime  time.Time `json:"openTime"`
	Open      string    `json:"open"`
	High      string    `json:"high"`
	Low       string    `json:"low"`
	Close     string    `json:"close"`
	Volume    string    `json:"volume"`
	CloseTime time.Time `json:"closeTime"`
}

// DepthResponse 深度数据响应结构体（用于 GetDepth）
type DepthResponse struct {
	Bids      [][]string `json:"bids"` // [price, quantity]
	Asks      [][]string `json:"asks"` // [price, quantity]
	Timestamp time.Time  `json:"timestamp"`
}

// TradesResponse 交易数据响应结构体（用于 GetTrades 和 GetHistoricalTrades）
type TradesResponse struct {
	TradeID      string    `json:"tradeId"`
	Price        string    `json:"price"`
	Quantity     string    `json:"quantity"`
	Timestamp    time.Time `json:"timestamp"`
	IsBuyerMaker bool      `json:"isBuyerMaker"`
}

// FundingRateResponse 资金费率响应结构体（用于 GetFundingIntervalRates）
type FundingRateResponse struct {
	Symbol      string    `json:"symbol"`
	FundingRate string    `json:"fundingRate"`
	Timestamp   time.Time `json:"timestamp"`
	FundingTime time.Time `json:"fundingTime"`
}

// OpenInterestResponse 开放兴趣响应结构体（用于 GetOpenInterest）
type OpenInterestResponse struct {
	Symbol       string    `json:"symbol"`
	OpenInterest string    `json:"openInterest"`
	Timestamp    time.Time `json:"timestamp"`
}

// MarkPriceResponse 标记价格响应结构体（用于 GetAllMarkPrices）
type MarkPriceResponse struct {
	Symbol    string    `json:"symbol"`
	MarkPrice string    `json:"markPrice"`
	Timestamp time.Time `json:"timestamp"`
}

// GetMarkets 获取所有市场信息 (/api/v1/markets)
func (c *BackpackClient) GetMarkets(marketType MarketType) ([]MarketResponse, error) {
	var resp []MarketResponse
	// 如果 marketType 是 ""，则返回所有类型市场
	return resp, nil
}

// GetMarket 获取特定市场信息 (/api/v1/markets/{symbol})
func (c *BackpackClient) GetMarket(symbol string) (MarketResponse, error) {
	var resp MarketResponse
	return resp, nil
}

// GetTickers 获取所有市场行情数据 (/api/v1/tickers)
func (c *BackpackClient) GetTickers(marketType MarketType) ([]TickerResponse, error) {
	var resp []TickerResponse
	// 如果 marketType 是 ""，则返回所有类型市场的行情
	return resp, nil
}

// GetTicker 获取特定市场行情数据 (/api/v1/tickers/{symbol})
func (c *BackpackClient) GetTicker(symbol string) (TickerResponse, error) {
	var resp TickerResponse
	return resp, nil
}

// GetKline 获取 K 线数据 (/api/v1/klines)
func (c *BackpackClient) GetKline(symbol string, interval KlineInterval, startTime, endTime int64, limit int) ([]KlineResponse, error) {
	var resp []KlineResponse
	// startTime, endTime 如果是 0，则视为未指定
	// limit 如果是 0，则使用默认值
	return resp, nil
}

// GetDepth 获取市场深度数据 (/api/v1/depth)
func (c *BackpackClient) GetDepth(symbol string, limit int) (DepthResponse, error) {
	var resp DepthResponse
	// limit 如果是 0，则使用默认值
	return resp, nil
}

// GetFundingIntervalRates 获取资金费率数据 (/api/v1/fundingRates)
func (c *BackpackClient) GetFundingIntervalRates(symbol string, startTime, endTime int64, limit int) ([]FundingRateResponse, error) {
	var resp []FundingRateResponse
	// startTime, endTime 如果是 0，则视为未指定
	// limit 如果是 0，则使用默认值
	return resp, nil
}

// GetOpenInterest 获取开放兴趣数据 (假设接口为 /api/v1/openInterest)
func (c *BackpackClient) GetOpenInterest(symbol string) (OpenInterestResponse, error) {
	var resp OpenInterestResponse
	return resp, nil
}

// GetAllMarkPrices 获取所有标记价格 (假设接口为 /api/v1/markPrices)
func (c *BackpackClient) GetAllMarkPrices(marketType MarketType) ([]MarkPriceResponse, error) {
	var resp []MarkPriceResponse
	// 如果 marketType 是 ""，则返回所有类型市场的标记价格
	return resp, nil
}

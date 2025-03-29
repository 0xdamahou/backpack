package public

// MarketSymbol 表示交易对的完整信息
type MarketSymbol struct {
	Symbol            string   `json:"symbol"`
	BaseSymbol        string   `json:"baseSymbol"`
	QuoteSymbol       string   `json:"quoteSymbol"`
	MarketType        string   `json:"marketType"`
	Filters           Filters  `json:"filters"`
	ImfFunction       Function `json:"imfFunction"`
	MmfFunction       Function `json:"mmfFunction"`
	FundingInterval   int      `json:"fundingInterval"`
	OpenInterestLimit string   `json:"openInterestLimit"`
	OrderBookState    string   `json:"orderBookState"`
	CreatedAt         string   `json:"createdAt"`
}

// Filters 表示交易对的过滤条件
type Filters struct {
	Price    PriceFilter    `json:"price"`
	Quantity QuantityFilter `json:"quantity"`
}

// PriceFilter 表示价格相关的过滤条件
type PriceFilter struct {
	MinPrice                    string        `json:"minPrice"`
	MaxPrice                    string        `json:"maxPrice"`
	TickSize                    string        `json:"tickSize"`
	MaxMultiplier               string        `json:"maxMultiplier"`
	MinMultiplier               string        `json:"minMultiplier"`
	MaxImpactMultiplier         string        `json:"maxImpactMultiplier"`
	MinImpactMultiplier         string        `json:"minImpactMultiplier"`
	MeanMarkPriceBand           MarkPriceBand `json:"meanMarkPriceBand"`
	MeanPremiumBand             PremiumBand   `json:"meanPremiumBand"`
	BorrowEntryFeeMaxMultiplier string        `json:"borrowEntryFeeMaxMultiplier"`
	BorrowEntryFeeMinMultiplier string        `json:"borrowEntryFeeMinMultiplier"`
}

// MarkPriceBand 表示标记价格范围
type MarkPriceBand struct {
	MaxMultiplier string `json:"maxMultiplier"`
	MinMultiplier string `json:"minMultiplier"`
}

// PremiumBand 表示溢价范围
type PremiumBand struct {
	TolerancePct string `json:"tolerancePct"`
}

// QuantityFilter 表示数量相关的过滤条件
type QuantityFilter struct {
	MinQuantity string `json:"minQuantity"`
	MaxQuantity string `json:"maxQuantity"`
	StepSize    string `json:"stepSize"`
}

// Function 表示 IMF 或 MMF 函数配置
type Function struct {
	Type   string `json:"type"`
	Base   string `json:"base"`
	Factor string `json:"factor"`
}

type Ticker struct {
	Symbol             string `json:"symbol"`
	FirstPrice         string `json:"firstPrice"`
	LastPrice          string `json:"lastPrice"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	High               string `json:"high"`
	Low                string `json:"low"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	Trades             string `json:"trades"`
}

type Depth struct {
	Asks         [][]string `json:"asks"`         // 卖单列表，每个元素是 [价格, 数量]
	Bids         [][]string `json:"bids"`         // 买单列表，每个元素是 [价格, 数量]
	LastUpdateID string     `json:"lastUpdateId"` // 最后更新 ID
	Timestamp    int64      `json:"timestamp"`    // 时间戳
}

type Kline struct {
	Start       string `json:"start"`       // 开始时间
	End         string `json:"end"`         // 结束时间
	Open        string `json:"open"`        // 开盘价
	High        string `json:"high"`        // 最高价
	Low         string `json:"low"`         // 最低价
	Close       string `json:"close"`       // 收盘价
	Volume      string `json:"volume"`      // 成交量
	QuoteVolume string `json:"quoteVolume"` // 报价成交量
	Trades      string `json:"trades"`      // 交易次数
}

// FundingInfo 表示资金费率相关数据
type FundingInfo struct {
	FundingRate          string `json:"fundingRate"`          // 资金费率
	IndexPrice           string `json:"indexPrice"`           // 指数价格
	MarkPrice            string `json:"markPrice"`            // 标记价格
	NextFundingTimestamp int64  `json:"nextFundingTimestamp"` // 下次资金费时间戳
	Symbol               string `json:"symbol"`               // 交易对符号
}
type OpenInterest struct {
	Symbol       string `json:"symbol"`       // 交易对符号
	OpenInterest string `json:"openInterest"` // 持仓量
	Timestamp    int64  `json:"timestamp"`    // 时间戳
}
type FundingRateHistory struct {
	Symbol               string `json:"symbol"`               // 交易对符号
	IntervalEndTimestamp string `json:"intervalEndTimestamp"` // 间隔结束时间戳
	FundingRate          string `json:"fundingRate"`          // 资金费率
}

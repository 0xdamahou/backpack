package authenticated

type BorrowLendEventType string

const (
	BorrowType      BorrowLendEventType = "Borrow"
	BorrowRepayType BorrowLendEventType = "BorrowRepay"
	LendType        BorrowLendEventType = "Lend"
	LendRedeemType  BorrowLendEventType = "LendRedeem"
)

type BorrowLendHistoryRequest struct {
	Type       *BorrowLendEventType `json:"type,omitempty"`
	Sources    *string              `json:"sources,omitempty"`
	PositionId *string              `json:"positionId,omitempty"`
	Symbol     *string              `json:"symbol,omitempty"`
	Limit      *uint64              `json:"limit,omitempty"`
	Offset     *uint64              `json:"offset,omitempty"`
}

func (blhr *BorrowLendHistoryRequest) ToURLQueryString() string {
	p := NewParams()

	p.AddUint64("limit", blhr.Limit)
	p.AddUint64("offset", blhr.Offset)
	p.Add("positionId", blhr.PositionId)
	p.Add("sources", blhr.Sources)
	p.Add("symbol", blhr.Symbol)
	if blhr.Type != nil {
		typeStr := string(*blhr.Type)
		p.Add("type", &typeStr)
	}

	return p.String()
}

type FillHistoryRequest struct {
	// OrderID is an optional filter for a specific order.
	OrderID *string `json:"orderId,omitempty"`

	// From is an optional filter for minimum time in milliseconds (<int64>).
	From *int64 `json:"from,omitempty"`

	// To is an optional filter for maximum time in milliseconds (<int64>).
	To *int64 `json:"to,omitempty"`

	// Symbol is an optional filter for a specific symbol.
	Symbol *string `json:"symbol,omitempty"`

	// Limit is an optional maximum number of records to return (<int64>).
	// Default: 100, Maximum: 1000.
	Limit *int64 `json:"limit,omitempty"`

	// Offset is an optional offset for pagination (<int64>).
	// Default: 0.
	Offset *int64 `json:"offset,omitempty"`

	// FillType is an optional filter for different fill types.
	// Enum: "User", "BookLiquidation", "Adl", "Backstop", "Liquidation",
	// "AllLiquidation", "CollateralConversion", "CollateralConversionAndSpotLiquidation"
	FillType *string `json:"fillType,omitempty"`

	// MarketType is an optional array of market types.
	// Items Enum: "SPOT", "PERP", "IPERP", "DATED", "PREDICTION", "RFQ"
	MarketType []string `json:"marketType,omitempty"`
}

func (fhq *FillHistoryRequest) ToURLQueryString() string {
	p := NewParams()
	p.Add("fillType", fhq.FillType)
	p.AddInt64("from", fhq.From)
	p.AddInt64("limit", fhq.Limit)
	for _, marketType := range fhq.MarketType {
		p.Add("marketType", &marketType)
	}
	p.AddInt64("offset", fhq.Offset)
	p.Add("orderId", fhq.OrderID)
	p.Add("symbol", fhq.Symbol)
	p.AddInt64("to", fhq.To)
	return p.String()

}

type OrderHistoryRequest struct {
	// OrderID is an optional filter for a specific order.
	OrderID *string `json:"orderId,omitempty"`

	// Symbol is an optional filter for a specific symbol.
	Symbol *string `json:"symbol,omitempty"`

	// Limit is an optional maximum number of records to return (<int64>).
	// Default: 100, Maximum: 1000.
	Limit *int64 `json:"limit,omitempty"`

	// Offset is an optional offset for pagination (<int64>).
	// Default: 0.
	Offset *int64 `json:"offset,omitempty"`

	// MarketType is an optional array of market types.
	// Items Enum: "SPOT", "PERP", "IPERP", "DATED", "PREDICTION", "RFQ"
	MarketType []string `json:"marketType,omitempty"`
}

func (ohq *OrderHistoryRequest) ToURLQueryString() string {
	params := NewParams()
	params.AddInt64("limit", ohq.Limit)
	for _, marketType := range ohq.MarketType {
		params.Add("marketType", &marketType)
	}
	params.AddInt64("offset", ohq.Offset)
	params.Add("orderId", ohq.OrderID)
	params.Add("symbol", ohq.Symbol)
	return params.String()
}

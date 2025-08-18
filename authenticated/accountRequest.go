package authenticated

type AccountConfig struct {
	AutoBorrowSettlements bool   `json:"autoBorrowSettlements"`
	AutoLend              bool   `json:"autoLend"`
	AutoRepayBorrows      bool   `json:"autoRepayBorrows"`
	LeverageLimit         string `json:"leverageLimit"`
}

func (mo *AccountConfig) ToURLQueryString() string {
	p := NewParams()
	p.AddBoolean("autoBorrowSettlements", &mo.AutoBorrowSettlements).AddBoolean("autoLend", &mo.AutoLend).AddBoolean("autoRepayBorrows", &mo.AutoRepayBorrows).Add("leverageLimit", &mo.LeverageLimit)
	return p.String()
}

type MaxOrderQuantityRequest struct {
	// Symbol is the required market symbol to trade.
	Symbol string `json:"symbol"`

	// Side is the required side of the order.
	// Enum: "Bid", "Ask"
	Side string `json:"side"`

	// Price is the optional limit price of the order (decimal as string).
	// Not included for market orders.
	Price *string `json:"price,omitempty"`

	// ReduceOnly is an optional flag indicating whether the order is reduce-only.
	ReduceOnly *bool `json:"reduceOnly,omitempty"`

	// AutoBorrow is an optional flag indicating whether the order uses auto-borrow.
	AutoBorrow *bool `json:"autoBorrow,omitempty"`

	// AutoBorrowRepay is an optional flag indicating whether the order uses auto-borrow repay.
	AutoBorrowRepay *bool `json:"autoBorrowRepay,omitempty"`

	// AutoLendRedeem is an optional flag indicating whether the order uses auto-lend redeem.
	AutoLendRedeem *bool `json:"autoLendRedeem,omitempty"`
}

func NewMaxOrderQuantityRequest(symbol string, side OrderSide) *MaxOrderQuantityRequest {
	return &MaxOrderQuantityRequest{
		Symbol: symbol,
		Side:   string(side),
	}
}

func (mo *MaxOrderQuantityRequest) WithPrice(price string) *MaxOrderQuantityRequest {
	mo.Price = &price
	return mo
}
func (mo *MaxOrderQuantityRequest) WithReduceOnly(reduceOnly bool) *MaxOrderQuantityRequest {
	mo.ReduceOnly = &reduceOnly
	return mo
}
func (mo *MaxOrderQuantityRequest) WithAutoBorrow(autoBorrow bool) *MaxOrderQuantityRequest {
	mo.AutoBorrow = &autoBorrow
	return mo
}
func (mo *MaxOrderQuantityRequest) WithAutoBorrowRepay(autoBorrowRepay bool) *MaxOrderQuantityRequest {
	mo.AutoBorrowRepay = &autoBorrowRepay
	return mo
}

func (mo *MaxOrderQuantityRequest) WithAutoLendRedeem(autoLendRedeem bool) *MaxOrderQuantityRequest {
	mo.AutoLendRedeem = &autoLendRedeem
	return mo
}
func (mo *MaxOrderQuantityRequest) ToURLQueryString() string {
	p := NewParams()
	p.AddBoolean("autoBorrow", mo.AutoBorrow).AddBoolean("autoBorrowRepay", mo.AutoBorrowRepay).AddBoolean("autoLendRedeem", mo.AutoLendRedeem).Add("price", mo.Price)
	p.AddBoolean("reduceOnly", mo.ReduceOnly).Add("side", &mo.Side).Add("symbol", &mo.Symbol)
	return p.String()
}

type MaxWithdrawalQuantityRequest struct {
	// Symbol is the required asset to withdraw.
	Symbol string `json:"symbol"`

	// AutoBorrow is an optional flag indicating whether the withdrawal uses auto-borrow.
	AutoBorrow *bool `json:"autoBorrow,omitempty"`

	// AutoLendRedeem is an optional flag indicating whether the withdrawal uses auto-lend redeem.
	AutoLendRedeem *bool `json:"autoLendRedeem,omitempty"`
}

func NewMaxWithdrawalQuantityRequest(symbol string) *MaxWithdrawalQuantityRequest {
	return &MaxWithdrawalQuantityRequest{Symbol: symbol}
}
func (mo *MaxWithdrawalQuantityRequest) WithAutoBorrow(autoBorrow bool) *MaxWithdrawalQuantityRequest {
	mo.AutoBorrow = &autoBorrow
	return mo
}
func (mo *MaxWithdrawalQuantityRequest) WithAutoLendRedeem(autoLendRedeem bool) *MaxWithdrawalQuantityRequest {
	mo.AutoLendRedeem = &autoLendRedeem
	return mo
}
func (mo *MaxWithdrawalQuantityRequest) ToURLQueryString() string {
	p := NewParams()
	p.AddBoolean("autoBorrow", mo.AutoBorrow).AddBoolean("autoLendRedeem", mo.AutoLendRedeem).Add("symbol", &mo.Symbol)
	return p.String()
}

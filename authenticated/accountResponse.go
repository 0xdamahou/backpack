package authenticated

type AccountSettings struct {
	AutoBorrowSettlements bool   `json:"autoBorrowSettlements"`
	AutoLend              bool   `json:"autoLend"`
	AutoRealizePnl        bool   `json:"autoRealizePnl"`
	AutoRepayBorrows      bool   `json:"autoRepayBorrows"`
	BorrowLimit           string `json:"borrowLimit"`
	FuturesMakerFee       string `json:"futuresMakerFee"`
	FuturesTakerFee       string `json:"futuresTakerFee"`
	LeverageLimit         string `json:"leverageLimit"`
	LimitOrders           int    `json:"limitOrders"`
	Liquidating           bool   `json:"liquidating"`
	PositionLimit         string `json:"positionLimit"`
	SpotMakerFee          string `json:"spotMakerFee"`
	SpotTakerFee          string `json:"spotTakerFee"`
	TriggerOrders         int    `json:"triggerOrders"`
}
type BorrowLimit struct {
	MaxBorrowQuantity string `json:"maxBorrowQuantity"`
	Symbol            string `json:"symbol"`
}

type MaxOrderResponse struct {
	AutoBorrow       bool   `json:"autoBorrow"`
	AutoBorrowRepay  bool   `json:"autoBorrowRepay"`
	AutoLendRedeem   bool   `json:"autoLendRedeem"`
	MaxOrderQuantity string `json:"maxOrderQuantity"`
	Price            string `json:"price"`
	Side             string `json:"side"`
	Symbol           string `json:"symbol"`
	ReduceOnly       bool   `json:"reduceOnly"`
}
type MaxWithdrawalResponse struct {
	AutoBorrow            bool   `json:"autoBorrow"`
	AutoLendRedeem        bool   `json:"autoLendRedeem"`
	MaxWithdrawalQuantity string `json:"maxWithdrawalQuantity"`
	Symbol                string `json:"symbol"`
}

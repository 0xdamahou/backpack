package authenticated

type Collateral struct {
	Symbol            string `json:"symbol"`
	AssetMarkPrice    string `json:"assetMarkPrice"`
	TotalQuantity     string `json:"totalQuantity"`
	BalanceNotional   string `json:"balanceNotional"`
	CollateralWeight  string `json:"collateralWeight"`
	CollateralValue   string `json:"collateralValue"`
	OpenOrderQuantity string `json:"openOrderQuantity"`
	LendQuantity      string `json:"lendQuantity"`
	AvailableQuantity string `json:"availableQuantity"`
}

type Account struct {
	AssetsValue        string       `json:"assetsValue"`
	BorrowLiability    string       `json:"borrowLiability"`
	Collateral         []Collateral `json:"collateral"`
	Imf                string       `json:"imf"`
	UnsettledEquity    string       `json:"unsettledEquity"`
	LiabilitiesValue   string       `json:"liabilitiesValue"`
	MarginFraction     string       `json:"marginFraction"`
	Mmf                string       `json:"mmf"`
	NetEquity          string       `json:"netEquity"`
	NetEquityAvailable string       `json:"netEquityAvailable"`
	NetEquityLocked    string       `json:"netEquityLocked"`
	NetExposureFutures string       `json:"netExposureFutures"`
	PnlUnrealized      string       `json:"pnlUnrealized"`
}
type Deposit struct {
	ID                      int    `json:"id"`
	ToAddress               string `json:"toAddress"`
	FromAddress             string `json:"fromAddress"`
	ConfirmationBlockNumber int    `json:"confirmationBlockNumber"`
	Source                  string `json:"source"`
	Status                  string `json:"status"`
	TransactionHash         string `json:"transactionHash"`
	Symbol                  string `json:"symbol"`
	Quantity                string `json:"quantity"`
	CreatedAt               string `json:"createdAt"`
}

type Withdrawal struct {
	ID              int    `json:"id"`
	Blockchain      string `json:"blockchain"`
	ClientID        string `json:"clientId"`
	Identifier      string `json:"identifier"`
	Quantity        string `json:"quantity"`
	Fee             string `json:"fee"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	SubaccountID    int    `json:"subaccountId"`
	ToAddress       string `json:"toAddress"`
	TransactionHash string `json:"transactionHash"`
	CreatedAt       string `json:"createdAt"`
	IsInternal      bool   `json:"isInternal"`
}

type WithdrawalResult struct {
	ID              int    `json:"id"`
	Blockchain      string `json:"blockchain"`
	ClientID        string `json:"clientId"`
	Identifier      string `json:"identifier"`
	Quantity        string `json:"quantity"`
	Fee             string `json:"fee"`
	Symbol          string `json:"symbol"`
	Status          string `json:"status"`
	SubaccountID    int    `json:"subaccountId"`
	ToAddress       string `json:"toAddress"`
	TransactionHash string `json:"transactionHash"`
	CreatedAt       string `json:"createdAt"`
	IsInternal      bool   `json:"isInternal"`
}

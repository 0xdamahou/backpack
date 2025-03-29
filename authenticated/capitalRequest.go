package authenticated

type WithdrawRequest struct {
	// Address is the required address to withdraw to.
	Address string `json:"address"`

	// Blockchain is the required blockchain to withdraw on.
	// Enum: "Arbitrum", "Base", "Berachain", "Bitcoin", "BitcoinCash", "Bsc", "Cardano",
	// "Dogecoin", "EqualsMoney", "Ethereum", "Hyperliquid", "Litecoin", "Polygon",
	// "Sui", "Solana", "Story", "XRP"
	Blockchain string `json:"blockchain"`

	// ClientID is an optional custom client ID.
	ClientID string `json:"clientId,omitempty"`

	// Quantity is the required quantity to withdraw (decimal as string).
	Quantity string `json:"quantity"`

	// Symbol is the required symbol of the asset to withdraw.
	// Enum: "BTC", "ETH", "SOL", "USDC", "USDT", "PYTH", "JTO", "BONK", "HNT", "MOBILE",
	// "WIF", "JUP", "RENDER", "WEN", "W", "TNSR", "PRCL", "SHARK", "KMNO", "MEW", "BOME",
	// "RAY", "HONEY", "SHFL", "BODEN", "IO", "DRIFT", "PEPE", "SHIB", "LINK", "UNI",
	// "ONDO", "FTM", "MATIC", "STRK", "BLUR", "WLD", "GALA", "NYAN", "HLG", "MON",
	// "ZKJ", "MANEKI", "HABIBI", "UNA", "ZRO", "ZEX", "AAVE", "LDO", "MOTHER", "CLOUD",
	// "MAX", "POL", "TRUMPWIN", "HARRISWIN", "MOODENG", "DBR", "GOAT", "ACT", "DOGE",
	// "BCH", "LTC", "APE", "ENA", "ME", "EIGEN", "CHILLGUY", "PENGU", "EUR", "SONIC",
	// "J", "TRUMP", "MELANIA", "ANIME", "XRP", "SUI", "VINE", "ADA", "MOVE", "BERA",
	// "IP", "HYPE", "BNB", "KAITO", "PEPE1000", "BONK1000", "SHIB1000", "AVAX", "S",
	// "POINTS", "ROAM", "AI16Z"
	Symbol string `json:"symbol"`

	// TwoFactorToken is an optional two-factor authentication token.
	TwoFactorToken string `json:"twoFactorToken,omitempty"`

	// AutoBorrow is an optional flag to auto-borrow if required.
	AutoBorrow bool `json:"autoBorrow,omitempty"`

	// AutoLendRedeem is an optional flag to auto-redeem a lend if required.
	AutoLendRedeem bool `json:"autoLendRedeem,omitempty"`
}

func NewWithdrawRequest(address string, blockchain string, symbol string, quantity string) *WithdrawRequest {
	return &WithdrawRequest{
		Address:    address,
		Blockchain: blockchain,
		Symbol:     symbol,
		Quantity:   quantity,
	}
}

func (req *WithdrawRequest) WithClientId(clientId string) *WithdrawRequest {
	req.ClientID = clientId
	return req
}
func (req *WithdrawRequest) WithAutoBorrow(autoBorrow bool) *WithdrawRequest {
	req.AutoBorrow = autoBorrow
	return req
}

func (req *WithdrawRequest) WithAutoLendRedeem(autoLendRedeem bool) *WithdrawRequest {
	req.AutoLendRedeem = autoLendRedeem
	return req
}
func (req *WithdrawRequest) WithTwoFactorToken(twoFactorToken string) *WithdrawRequest {
	req.TwoFactorToken = twoFactorToken
	return req
}

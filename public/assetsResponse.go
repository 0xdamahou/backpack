package public

// SymbolInfo 表示交易对及其代币信息
type AssetInfo struct {
	Symbol string      `json:"symbol"` // 交易对符号，例如 "BTC"
	Tokens []TokenInfo `json:"tokens"` // 代币信息列表
}

// TokenInfo 表示单个代币的详细信息
type TokenInfo struct {
	Blockchain        string `json:"blockchain"`        // 区块链名称，例如 "Arbitrum"
	ContractAddress   string `json:"contractAddress"`   // 合约地址
	DepositEnabled    bool   `json:"depositEnabled"`    // 是否启用存款
	MinimumDeposit    string `json:"minimumDeposit"`    // 最小存款金额
	WithdrawEnabled   bool   `json:"withdrawEnabled"`   // 是否启用提款
	MinimumWithdrawal string `json:"minimumWithdrawal"` // 最小提款金额
	MaximumWithdrawal string `json:"maximumWithdrawal"` // 最大提款金额
	WithdrawalFee     string `json:"withdrawalFee"`     // 提款费用
}

// CollateralParameters 表示交易对及其函数配置
type CollateralParameters struct {
	Symbol          string             `json:"symbol"`          // 交易对符号
	ImfFunction     CollateralFunction `json:"imfFunction"`     // IMF 函数配置
	MmfFunction     CollateralFunction `json:"mmfFunction"`     // MMF 函数配置
	HaircutFunction HaircutFunction    `json:"haircutFunction"` // Haircut 函数配置
}

// Function 表示 IMF 或 MMF 函数配置
type CollateralFunction struct {
	Type   string `json:"type"`   // 函数类型，例如 "sqrt"
	Base   string `json:"base"`   // 基础值
	Factor string `json:"factor"` // 因子
}

// HaircutFunction 表示 Haircut 函数配置
type HaircutFunction struct {
	Weight string      `json:"weight"` // 权重
	Kind   HaircutKind `json:"kind"`   // Haircut 类型配置
}

// HaircutKind 表示 Haircut 函数的具体类型
type HaircutKind struct {
	Type string `json:"type"` // 类型，例如 "identity"
}

package authenticated

import (
	"encoding/json"
	"fmt"
)

// GetBalances  return map[SOL:map[available:0.800001 locked:0 staked:0]]
func (c *BackpackClient) GetBalances() (map[string]map[string]string, error) {
	instruction := "balanceQuery"
	endpoint := "api/v1/capital"
	var result map[string]map[string]string
	err := c.DoGet(endpoint, instruction, "", &result)
	return result, err
}

func (c *BackpackClient) GetCollateral(subAccountId uint16) (Account, error) {
	instruction := "collateralQuery"
	endpoint := "api/v1/capital/collateral"
	params := fmt.Sprintf("subaccountId=%d", subAccountId)
	var account Account
	err := c.DoGet(endpoint, instruction, params, &account)
	return account, err
}

func (c *BackpackClient) GetDeposit(from int64, to int64, limit int64, offset int64) ([]Deposit, error) {
	instruction := "depositQueryAll"
	endpoint := "wapi/v1/capital/deposits"
	if limit > 1000 || limit < 0 {
		limit = 100
	}
	if offset < 0 || offset > limit {
		offset = 0
	}

	params := fmt.Sprintf("from=%d&limit=%d&offset=%d&to=%d", from, limit, offset, to)
	var deposit []Deposit
	err := c.DoGet(endpoint, instruction, params, &deposit)
	return deposit, err
}

/*
Enum: "Arbitrum" "Base" "Berachain" "Bitcoin" "BitcoinCash" "Bsc" "Cardano" "Dogecoin" "EqualsMoney" "Ethereum" "Hyperliquid" "Litecoin" "Polygon" "Sui" "Solana" "Story" "XRP"
Blockchain symbol to get a deposit address for.
*/
func (c *BackpackClient) GetDepositAddress(blockChain string) (string, error) {
	instruction := "depositAddressQuery"
	endpoint := "wapi/v1/capital/deposit/address"

	params := fmt.Sprintf("blockchain=%s", blockChain)
	var address struct {
		Address string `json:"address"`
	}
	err := c.DoGet(endpoint, instruction, params, &address)
	return address.Address, err
}
func (c *BackpackClient) GetWithdrawal(from int64, to int64, limit int64, offset int64) ([]Withdrawal, error) {
	instruction := "withdrawalQueryAll"
	endpoint := "wapi/v1/capital/withdrawals"
	if limit > 1000 || limit < 0 {
		limit = 100
	}
	if offset < 0 || offset > limit {
		offset = 0
	}
	params := fmt.Sprintf("from=%d&limit=%d&offset=%d&to=%d", from, limit, offset, to)
	var withdrawals []Withdrawal
	err := c.DoGet(endpoint, instruction, params, &withdrawals)
	return withdrawals, err
}

func (c *BackpackClient) RequestWithdrawal(dr *WithdrawRequest) (WithdrawalResult, error) {
	instruction := "withdraw"
	endpoint := "wapi/v1/capital/withdrawals"
	var result WithdrawalResult
	jsonBody, err := json.Marshal(dr)
	if err != nil {
		return result, err
	}
	err = c.DoPost(endpoint, instruction, jsonBody, &result)
	return result, err
}

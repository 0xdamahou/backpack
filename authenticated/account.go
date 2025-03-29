package authenticated

import (
	"encoding/json"
	"fmt"
	"log"
)

func (c *BackpackClient) GetAccount() (AccountSettings, error) {
	instruction := "accountQuery"
	endpoint := "api/v1/account"
	var result AccountSettings
	err := c.DoGet(endpoint, instruction, "", &result)
	return result, err
}

func (c *BackpackClient) UpdateAccount(autoBorrowSettlements bool, autoLend bool, autoRepayBorrows bool, leverageLimit string) error {
	instruction := "accountUpdate"
	endpoint := "api/v1/account"
	ac := AccountConfig{
		AutoBorrowSettlements: autoBorrowSettlements,
		AutoLend:              autoLend,
		AutoRepayBorrows:      autoRepayBorrows,
		LeverageLimit:         leverageLimit,
	}
	bs, err := json.Marshal(ac)
	if err != nil {
		log.Println(err)
		return err
	}
	err = c.DoPatch(endpoint, instruction, bs, nil)
	return err
}

func (c *BackpackClient) GetMaxBorrowQuantity(symbol string) (BorrowLimit, error) {
	instruction := "maxBorrowQuantity"
	endpoint := "api/v1/account/limits/borrow"
	params := fmt.Sprintf("symbol=%s", symbol)
	var result BorrowLimit
	err := c.DoGet(endpoint, instruction, params, &result)
	return result, err
}

func (c *BackpackClient) GetMaxOrderQuantity(mr *MaxOrderQuantityRequest) (MaxOrderResponse, error) {
	instruction := "maxOrderQuantity"
	endpoint := "api/v1/account/limits/order"
	var result MaxOrderResponse
	q := mr.ToURLQueryString()
	//log.Println(q)
	err := c.DoGet(endpoint, instruction, q, &result)
	return result, err
}

func (c *BackpackClient) GetMaxWithdrawalQuantity(mr *MaxWithdrawalQuantityRequest) (MaxWithdrawalResponse, error) {
	instruction := "maxWithdrawalQuantity"
	endpoint := "api/v1/account/limits/withdrawal"
	var result MaxWithdrawalResponse
	err := c.DoGet(endpoint, instruction, mr.ToURLQueryString(), &result)
	return result, err
}

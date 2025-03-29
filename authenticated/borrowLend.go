package authenticated

import (
	"encoding/json"
)

func (c *BackpackClient) GetBorrowLendPositions() (Positions, error) {
	endpoint := "api/v1/borrowLend/positions"
	instruction := "borrowLendPositionQuery"
	var positions Positions
	err := c.DoGet(endpoint, instruction, "", &positions)
	return positions, err
}

func (c *BackpackClient) ExecuteBorrowLend(symbol string, side string, quantity string) error {
	endpoint := "api/v1/borrowLend"
	instruction := "borrowLendExecute"
	type Request struct {
		Quantity string `json:"quantity"`
		Side     string `json:"side"`
		Symbol   string `json:"symbol"`
	}
	r := Request{Quantity: quantity, Side: side, Symbol: symbol}
	bs, e := json.Marshal(r)
	if e != nil {
		return e
	}
	return c.DoPost(endpoint, instruction, bs, nil)
}

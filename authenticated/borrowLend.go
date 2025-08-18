package authenticated

import (
	//"encoding/json"

	"bytes"

	json "github.com/bytedance/sonic"
)

func (c *BackpackClient) GetBorrowLendPositions() (Positions, error) {
	endpoint := "api/v1/borrowLend/positions"
	instruction := "borrowLendPositionQuery"
	var positions Positions
	err := c.DoGet(endpoint, instruction, "", &positions)
	return positions, err
}

type BorrowLendRequest struct {
	Quantity string `json:"quantity"`
	Side     string `json:"side"`
	Symbol   string `json:"symbol"`
}

func (r *BorrowLendRequest) ToURLQueryString() string {
	p := NewParams()
	p.Add("quantity", &r.Quantity)
	p.Add("side", &r.Side)
	p.Add("symbol", &r.Symbol)
	return p.String()

}
func (c *BackpackClient) ExecuteBorrowLend(symbol string, side string, quantity string) error {
	endpoint := "api/v1/borrowLend"
	instruction := "borrowLendExecute"
	r := BorrowLendRequest{Quantity: quantity, Side: side, Symbol: symbol}
	var buf bytes.Buffer
	err := json.ConfigDefault.NewEncoder(&buf).Encode(r)
	if err != nil {
		return err
	}
	return c.DoPost(endpoint, instruction, &buf, r.ToURLQueryString(), nil)
}

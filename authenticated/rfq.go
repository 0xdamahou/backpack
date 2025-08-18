package authenticated

import (
	"bytes"

	json "github.com/bytedance/sonic"
)

func (c *BackpackClient) SubmitQuote(rfq *RfqRequest) (RFQQuoteResponse, error) {
	endPoint := "api/v1/rfq/quote"
	instruction := "quoteSubmit"
	var resp RFQQuoteResponse
	var buf bytes.Buffer
	err := json.ConfigDefault.NewEncoder(&buf).Encode(*rfq)
	if err != nil {
		return resp, err
	}
	err = c.DoPost(endPoint, instruction, &buf, rfq.ToURLQueryString(), &resp)
	return resp, err
}

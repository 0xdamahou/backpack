package authenticated

import json "github.com/bytedance/sonic"

func (c *BackpackClient) SubmitQuote(rfq *RfqRequest) (RFQQuoteResponse, error) {
	endPoint := "api/v1/rfq/quote"
	instruction := "quoteSubmit"
	var resp RFQQuoteResponse
	bs, err := json.Marshal(*rfq)
	if err != nil {
		return resp, err
	}
	err = c.DoPost(endPoint, instruction, bs, &resp)
	return resp, err
}

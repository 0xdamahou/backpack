package authenticated

type RFQQuoteResponse struct {
	RfqId    string `json:"rfqId"`
	QuoteId  string `json:"quoteId"`
	ClientId int    `json:"clientId"`
	Price    string `json:"price"`
	Status   string `json:"status"`
}

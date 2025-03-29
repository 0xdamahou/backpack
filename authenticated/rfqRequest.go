package authenticated

type RfqRequest struct {
	AskPrice string  `json:"askPrice"`
	BidPrice string  `json:"bidPrice"`
	ClientId *uint32 `json:"clientId,omitempty"`
	RfqId    string  `json:"rfqId"`
}

func NewRfqRequest(rfqId string, bidPrice string, askPrice string) *RfqRequest {
	return &RfqRequest{RfqId: rfqId, BidPrice: bidPrice, AskPrice: askPrice}
}

func (r *RfqRequest) WithClientId(clientId uint32) {
	r.ClientId = &clientId
}

func (r *RfqRequest) ToURLQueryString() string {
	p := Params{}
	p.Add("askPrice", &r.AskPrice)
	p.Add("bidPrice", &r.BidPrice)
	p.AddUint32("clientId", r.ClientId)
	p.Add("rfqId", &r.RfqId)
	return p.String()
}

type UrlBuilding interface {
	ToURLQueryString() string
}

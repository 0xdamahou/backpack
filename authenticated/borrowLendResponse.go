package authenticated

type Position struct {
	CumulativeInterest  string       `json:"cumulativeInterest"`
	ID                  string       `json:"id"`
	IMF                 string       `json:"imf"`
	IMFFunction         RiskFunction `json:"imfFunction"`
	NetQuantity         string       `json:"netQuantity"`
	MarkPrice           string       `json:"markPrice"`
	MMF                 string       `json:"mmf"`
	MMFFunction         RiskFunction `json:"mmfFunction"`
	NetExposureQuantity string       `json:"netExposureQuantity"`
	NetExposureNotional string       `json:"netExposureNotional"`
	Symbol              string       `json:"symbol"`
}

type RiskFunction struct {
	Type   string `json:"type"`
	Base   string `json:"base"`
	Factor string `json:"factor"`
}

type Positions []Position

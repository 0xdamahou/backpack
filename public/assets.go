package public

func (bbc *BackpackPublicClient) GetAssets() ([]AssetInfo, error) {
	endpoint := "api/v1/assets"
	var rates []AssetInfo

	err := bbc.DoGet(endpoint, "", &rates)
	return rates, err
}

// Get collateral
func (bbc *BackpackPublicClient) GetCollateral() ([]CollateralParameters, error) {
	endpoint := "api/v1/collateral"
	var rates []CollateralParameters
	err := bbc.DoGet(endpoint, "", &rates)
	return rates, err
}

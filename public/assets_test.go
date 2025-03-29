package public

import (
	"log"
	"testing"
)

func TestBackpackPublicClient_GetAssets(t *testing.T) {
	assets, err := bpbc.GetAssets()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, asset := range assets {
		log.Printf("%+v\n", asset)
	}
}
func TestBackpackPublicClient_GetCollateral(t *testing.T) {
	collateral, err := bpbc.GetCollateral()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, collateral := range collateral {
		log.Printf("%+v\n", collateral)
	}
}

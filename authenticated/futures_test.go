package authenticated

import (
	"log"
	"testing"
)

func TestBackpackClient_GetOpenPositions(t *testing.T) {
	positions, err := client.GetOpenPositions()
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, position := range positions {
		log.Printf("%+v", position)
	}
}

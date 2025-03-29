package authenticated

import (
	"log"
	"testing"
)

func TestBackpackClient_GetBorrowLendPositions(t *testing.T) {
	positions, err := client.GetBorrowLendPositions()
	if err != nil {
		t.Fatal(err)
		return
	}
	log.Println(positions)
}

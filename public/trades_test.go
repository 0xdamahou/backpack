package public

import (
	"log"
	"testing"
)

func TestBackpackPublicClient_GetRecentTrades(t *testing.T) {
	recentTrades, err := bpbc.GetRecentTrades("SOL_USDC", 100)
	if err != nil {
		return
	}
	log.Println(recentTrades)
}
func TestBackpackPublicClient_GetHistoryTrades(t *testing.T) {
	historyTrades, err := bpbc.GetHistoryTrades("SOL_USDC", 100, 100)
	if err != nil {
		return
	}
	log.Println(historyTrades)
}

package public

import "fmt"

func (bbc *BackpackPublicClient) GetRecentTrades(symbol string, limit uint16) ([]RecentTrade, error) {
	endpoint := "api/v1/trades"
	query := fmt.Sprintf("limit=%d&symbol=%s", limit, symbol)
	var result []RecentTrade
	err := bbc.DoGet(endpoint, query, &result)
	return result, err
}
func (bbc *BackpackPublicClient) GetHistoryTrades(symbol string, limit uint64, offset uint64) ([]RecentTrade, error) {
	endpoint := "api/v1/trades/history"
	query := fmt.Sprintf("limit=%d&symbol=%s", limit, symbol)
	if offset > 0 {
		query = fmt.Sprintf("%s&offset=%d", query, offset)
	}
	var result []RecentTrade
	err := bbc.DoGet(endpoint, query, &result)
	return result, err
}

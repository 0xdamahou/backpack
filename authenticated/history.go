package authenticated

func (c *BackpackClient) GetBorrowHistory(blhr *BorrowLendHistoryRequest) ([]BorrowHistoryItem, error) {
	instruction := "borrowHistoryQueryAll"
	endpoint := "wapi/v1/history/borrowLend"
	var result []BorrowHistoryItem
	err := c.DoGet(endpoint, instruction, blhr.ToURLQueryString(), &result)
	return result, err
}

func (c *BackpackClient) GetFillHistory(fh *FillHistoryRequest) ([]FillHistoryResponse, error) {
	instruction := "fillHistoryQueryAll"
	endpoint := "wapi/v1/history/fills"
	var result []FillHistoryResponse
	err := c.DoGet(endpoint, instruction, fh.ToURLQueryString(), &result)
	return result, err
}
func (c *BackpackClient) GetOrderHistory(fh *OrderHistoryRequest) ([]OrderHistoryResponse, error) {
	instruction := "orderHistoryQueryAll"
	endpoint := "wapi/v1/history/orders"
	var result []OrderHistoryResponse
	err := c.DoGet(endpoint, instruction, fh.ToURLQueryString(), &result)
	return result, err
}

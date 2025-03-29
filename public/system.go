package public

func (bbc *BackpackPublicClient) Status() (StatusResponse, error) {
	endpoint := "api/v1/status"
	var response StatusResponse
	err := bbc.DoGet(endpoint, "", &response)
	return response, err
}

func (bbc *BackpackPublicClient) Ping() (bool, error) {
	endpoint := "api/v1/ping"
	result, err := bbc.Get(endpoint, "")
	return result == "pong", err
}

func (bbc *BackpackPublicClient) GetSystemTime() (string, error) {
	endpoint := "api/v1/time"
	return bbc.Get(endpoint, "")

}

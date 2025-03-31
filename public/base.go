package public

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/0xdamahou/backpack/authenticated"
	"io"
	"log"
	"net/http"
)

const (
	baseUrl = "https://api.backpack.exchange/"
)

type BackpackPublicClient struct {
	BaseUrl string
	Client  *http.Client
}

func NewBackpackPublicClient() *BackpackPublicClient {
	return &BackpackPublicClient{BaseUrl: baseUrl, Client: &http.Client{}}
}
func (bbc *BackpackPublicClient) DoGet(endpoint string, query string, result interface{}) error {
	return bbc.DoJson(authenticated.GET, endpoint, nil, query, result)
}

func (bbc *BackpackPublicClient) DoJson(method, endpoint string, reqBody []byte, query string, result interface{}) error {
	url := bbc.BaseUrl + endpoint
	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
		return err
	}

	if len(query) > 0 {
		req.URL.RawQuery = query
	}
	//log.Printf("Request URL: %s\n", req.URL)

	req.Header.Set("Content-Type", "application/json")

	resp, err := bbc.Client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bs, _ := io.ReadAll(resp.Body)
		log.Println(string(bs))
		return fmt.Errorf("request failed with status: %s\n", resp.Status)

	}
	//return json.Unmarshal(bs, result)
	return json.NewDecoder(resp.Body).Decode(result)

}
func (bbc *BackpackPublicClient) Get(endpoint string, query string) (string, error) {
	url := bbc.BaseUrl + endpoint
	req, err := http.NewRequest(authenticated.GET, url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if len(query) > 0 {
		req.URL.RawQuery = query
	}
	//log.Printf("Request URL: %s\n", req.URL)

	req.Header.Set("Content-Type", "application/json")

	resp, err := bbc.Client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	bs, err := io.ReadAll(resp.Body)

	return string(bs), err
}

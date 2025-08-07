package authenticated

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"net/http"
)

type BackpackClient struct {
	APIKey     string
	PrivateKey ed25519.PrivateKey
	Client     ClientProvider
}
type ClientProvider interface {
	GetClient() *http.Client
}

type DefaultClientProvider struct {
}

func (c *DefaultClientProvider) GetClient() *http.Client {
	return http.DefaultClient
}

func NewBackpackClient(apiKey, privateKey string) (*BackpackClient, error) {
	dcp := &DefaultClientProvider{}
	return NewBackpackClientEx(apiKey, privateKey, dcp)
}

func NewBackpackClientEx(apiKey, privateKey string, client ClientProvider) (*BackpackClient, error) {
	seed, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, fmt.Errorf("decode private key failed: %v", err)
	}
	if len(seed) != 32 {
		return nil, fmt.Errorf("invalid private key length: got %d, expected 32 bytes for seed", len(seed))
	}
	privKey := ed25519.NewKeyFromSeed(seed)

	return &BackpackClient{
		APIKey:     apiKey,
		PrivateKey: privKey,
		Client:     client,
	}, nil
}

package gokirimwa

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient HTTPClient
	apiKey     string
	baseURl    string
}

func NewKirimWA(apiKey string) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 30 * time.Second},
		apiKey:     apiKey,
		baseURl:    "https://api.kirimwa.id/v1",
	}
}

func NewKirimWAMockWithClient(apiKey string, httpClient HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
		apiKey:     apiKey,
		baseURl:    "https://api.kirimwa.id/v1",
	}
}

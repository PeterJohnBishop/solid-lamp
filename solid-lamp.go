package solidlamp

import "net/http"

type ClickUpClient struct {
	APIKey     string
	HTTPClient *http.Client
}

func NewClickUpClient(apiKey string) *ClickUpClient {
	return &ClickUpClient{
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

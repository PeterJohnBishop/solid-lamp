package solidlamp

import (
	"log"
	"net/http"
)

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

func Test() {
	client := NewClickUpClient("your_api_key")
	user, err := client.GetAuthorizedUser()
	if err != nil {
		log.Println("Error getting Authorized user!")
	}
	println("Authorized user:", user.Username)
}

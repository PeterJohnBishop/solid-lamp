# solid-lamp 

A basic 1:1 replication of the ClickUp V2 API accessible through Client. 

Go gettable at github.com/PeterJohnBishop/solid-lamp (to be updated).
- go get -u github.com/PeterJohnBishop/solid-lamp to update package

# example

func NewClickUpClient(apiKey string) *ClickUpClient {
	return &ClickUpClient{
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

client := NewClickUpClient("your_api_key")
user , err := client.GetAuthorizedUser()
	if err != nil {
		log.Println("Error getting Authorized user!")
	}
	println("Authorized user:", user.Username)



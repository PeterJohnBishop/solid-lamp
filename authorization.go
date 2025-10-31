package solidlamp

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (c *ClickUpClient) GetAccessToken(reqBody GetAccessTokenRequest) (*AccessTokenResponse, error) {
	baseURL := "https://api.clickup.com/api/v2/oauth/token"

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", baseURL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var token AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (c *ClickUpClient) GetAuthorizedUser() (*AuthorizedUserResponse, error) {
	baseURL := "https://api.clickup.com/api/v2/user"

	req, err := http.NewRequest("POST", baseURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var user AuthorizedUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

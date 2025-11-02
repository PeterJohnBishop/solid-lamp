package solidlamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/PeterJohnBishop/solid-lamp/cu"
)

func (c *ClickUpClient) GetAccessToken(reqBody cu.GetAccessTokenOptions) (*cu.AccessTokenResponse, error) {
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

	var token cu.AccessTokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (c *ClickUpClient) GetAuthorizedUser() (*cu.AuthorizedUser, error) {
	baseURL := "https://api.clickup.com/api/v2/user"

	req, err := http.NewRequest("GET", baseURL, nil)
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
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var authorizedUserResponse cu.AuthorizedUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&authorizedUserResponse); err != nil {
		return nil, err
	}

	return &authorizedUserResponse.User, nil
}

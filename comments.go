package solidlamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PeterJohnBishop/solid-lamp/cu"
)

func (c *ClickUpClient) GetTaskComments(taskId string, params cu.GetTaskCommentsQueryParams) (*[]cu.Comment, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/comment", url.PathEscape(taskId))

	q := url.Values{}

	if params.CustomTaskIds != nil {
		q.Set("custom_task_ids", fmt.Sprintf("%t", *params.CustomTaskIds))
	}
	if params.TeamId != nil {
		q.Set("team_id", fmt.Sprintf("%d", *params.TeamId))
	}
	if params.Start != nil {
		q.Set("start", fmt.Sprintf("%d", *params.Start))
	}
	if params.StartId != nil {
		q.Set("start_id", *params.StartId)
	}

	reqURL := baseURL
	if len(q) > 0 {
		reqURL = fmt.Sprintf("%s?%s", baseURL, q.Encode())
	}

	req, err := http.NewRequest("GET", reqURL, nil)
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

	var comments []cu.Comment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}

	return &comments, nil
}

func (c *ClickUpClient) CreateTaskComment(taskId string, params cu.CreateTaskCommentQueryParams, reqBody cu.CreateTaskCommentOptions) (*cu.CreateTaskCommentResponse, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/comment", url.PathEscape(taskId))

	q := url.Values{}

	if params.CustomTaskIds != nil {
		q.Set("custom_task_ids", fmt.Sprintf("%t", *params.CustomTaskIds))
	}
	if params.TeamId != nil {
		q.Set("team_id", fmt.Sprintf("%d", *params.TeamId))
	}

	reqURL := baseURL
	if len(q) > 0 {
		reqURL = fmt.Sprintf("%s?%s", baseURL, q.Encode())
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bodyBytes)

	req, err := http.NewRequest("POST", reqURL, body)
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

	var createCommentResp cu.CreateTaskCommentResponse
	if err := json.NewDecoder(resp.Body).Decode(&createCommentResp); err != nil {
		return nil, err
	}

	return &createCommentResp, nil
}

func (c *ClickUpClient) GetChatViewComments(viewId string, params cu.GetChatViewCommentsQueryParams) (*[]cu.ChatViewComment, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/view/%s/comment", url.PathEscape(viewId))

	q := url.Values{}

	if params.Start != nil {
		q.Set("start", fmt.Sprintf("%d", *params.Start))
	}
	if params.StartId != nil {
		q.Set("start_id", *params.StartId)
	}

	reqURL := baseURL
	if len(q) > 0 {
		reqURL = fmt.Sprintf("%s?%s", baseURL, q.Encode())
	}

	req, err := http.NewRequest("GET", reqURL, nil)
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

	var comments []cu.ChatViewComment
	if err := json.NewDecoder(resp.Body).Decode(&comments); err != nil {
		return nil, err
	}

	return &comments, nil
}

func (c *ClickUpClient) CreateChatViewComment(viewId string, reqBody cu.CreateChatViewCommentOptions) (*cu.CreateTaskCommentResponse, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/view/%s/comment", url.PathEscape(viewId))

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

	var createCommentResp cu.CreateTaskCommentResponse
	if err := json.NewDecoder(resp.Body).Decode(&createCommentResp); err != nil {
		return nil, err
	}

	return &createCommentResp, nil
}

func (c *ClickUpClient) GetListComments(listId string, params cu.GetListCommentsQueryParams) (*[]cu.ListComment, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/list/%s/comment", url.PathEscape(listId))

	q := url.Values{}

	if params.Start != nil {
		q.Set("start", fmt.Sprintf("%d", *params.Start))
	}
	if params.StartId != nil {
		q.Set("start_id", *params.StartId)
	}

	reqURL := baseURL
	if len(q) > 0 {
		reqURL = fmt.Sprintf("%s?%s", baseURL, q.Encode())
	}

	req, err := http.NewRequest("GET", reqURL, nil)
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

	var listComments []cu.ListComment
	if err := json.NewDecoder(resp.Body).Decode(&listComments); err != nil {
		return nil, err
	}

	return &listComments, nil
}

package solidlamp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *ClickUpClient) GetTasks(listID string, params GetTaskQueryParams) (*Tasks, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/list/%s/task", url.PathEscape(listID))

	q := url.Values{}

	if params.Archived != nil {
		q.Set("archived", fmt.Sprintf("%t", *params.Archived))
	}
	if params.IncludeMarkdownDescription != nil {
		q.Set("include_markdown_description", fmt.Sprintf("%t", *params.IncludeMarkdownDescription))
	}
	if params.Page != nil {
		q.Set("page", fmt.Sprintf("%d", *params.Page))
	}
	if params.OrderBy != nil {
		q.Set("order_by", *params.OrderBy)
	}
	if params.Reverse != nil {
		q.Set("reverse", fmt.Sprintf("%t", *params.Reverse))
	}
	if params.Subtasks != nil {
		q.Set("subtasks", fmt.Sprintf("%t", *params.Subtasks))
	}
	if params.IncludeClosed != nil {
		q.Set("include_closed", fmt.Sprintf("%t", *params.IncludeClosed))
	}
	if params.IncludeTiml != nil {
		q.Set("include_timl", fmt.Sprintf("%t", *params.IncludeTiml))
	}

	for _, status := range params.Statuses {
		q.Add("statuses[]", status)
	}
	for _, assignee := range params.Assignees {
		q.Add("assignees[]", assignee)
	}
	for _, watcher := range params.Watchers {
		q.Add("watchers[]", watcher)
	}
	for _, tag := range params.Tags {
		q.Add("tags[]", tag)
	}
	for _, customField := range params.CustomField {
		q.Add("custom_field[]", customField)
	}
	for _, item := range params.CustomItems {
		q.Add("custom_items[]", fmt.Sprintf("%d", item))
	}

	if params.DueDateGt != nil {
		q.Set("due_date_gt", fmt.Sprintf("%d", *params.DueDateGt))
	}
	if params.DueDateLt != nil {
		q.Set("due_date_lt", fmt.Sprintf("%d", *params.DueDateLt))
	}
	if params.DateCreatedGt != nil {
		q.Set("date_created_gt", fmt.Sprintf("%d", *params.DateCreatedGt))
	}
	if params.DateCreatedLt != nil {
		q.Set("date_created_lt", fmt.Sprintf("%d", *params.DateCreatedLt))
	}
	if params.DateUpdatedGt != nil {
		q.Set("date_updated_gt", fmt.Sprintf("%d", *params.DateUpdatedGt))
	}
	if params.DateUpdatedLt != nil {
		q.Set("date_updated_lt", fmt.Sprintf("%d", *params.DateUpdatedLt))
	}
	if params.DateDoneGt != nil {
		q.Set("date_done_gt", fmt.Sprintf("%d", *params.DateDoneGt))
	}
	if params.DateDoneLt != nil {
		q.Set("date_done_lt", fmt.Sprintf("%d", *params.DateDoneLt))
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
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return &tasks, nil
}

package solidlamp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PeterJohnBishop/solid-lamp/cu"
)

func (c *ClickUpClient) CreateTask(listId string, reqBody cu.CreateTaskOptions) (*cu.Task, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/list/%s/task", url.PathEscape(listId))

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

	var task cu.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *ClickUpClient) GetTasks(listID string, params cu.GetTasksQueryParams) (*cu.Tasks, error) {
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
	if params.CustomFields != nil {
		customFieldsBytes, err := json.Marshal(params.CustomFields)
		if err != nil {
			return nil, err
		}
		q.Set("custom_fields", string(customFieldsBytes))
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
		return nil, err
	}

	var tasks cu.Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (c *ClickUpClient) GetFilteredTeamTasks(teamId string, params cu.GetFilteredTeamTasksQueryParams) (*cu.Tasks, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/team/%s/task", url.PathEscape(teamId))

	q := url.Values{}

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
	if params.CustomFields != nil {
		customFieldsBytes, err := json.Marshal(params.CustomFields)
		if err != nil {
			return nil, err
		}
		q.Set("custom_fields", string(customFieldsBytes))
	}

	for _, space := range params.SpaceIds {
		q.Add("space_ids[]", space)
	}
	for _, project := range params.ProjectIds {
		q.Add("project_ids[]", project)
	}
	for _, list := range params.ListIds {
		q.Add("list_ids[]", list)
	}
	for _, status := range params.Statuses {
		q.Add("statuses[]", status)
	}
	for _, assignee := range params.Assignees {
		q.Add("assignees[]", assignee)
	}
	for _, tag := range params.Tags {
		q.Add("tags[]", tag)
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

	var tasks cu.Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (c *ClickUpClient) GetTask(taskId string, params cu.GetTaskQueryParams) (*cu.Task, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s", url.PathEscape(taskId))

	q := url.Values{}

	if params.CustomTaskIds != nil {
		q.Set("custom_task_ids", fmt.Sprintf("%t", *params.CustomTaskIds))
	}
	if params.TeamId != nil {
		q.Set("team_id", fmt.Sprintf("%d", *params.TeamId))
	}
	if params.IncludeMarkdownDescription != nil {
		q.Set("include_markdown_description", fmt.Sprintf("%t", *params.IncludeMarkdownDescription))
	}
	if params.CustomFields != nil {
		customFieldsBytes, err := json.Marshal(params.CustomFields)
		if err != nil {
			return nil, err
		}
		q.Set("custom_fields", string(customFieldsBytes))
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

	var task cu.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil

}

func (c *ClickUpClient) UpdateTask(taskId string, params cu.UpdateTaskQueryParams) (*cu.Task, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s", url.PathEscape(taskId))

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

	req, err := http.NewRequest("PUT", reqURL, nil)
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

	var task cu.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *ClickUpClient) DeleteTask(taskId string, params cu.DeleteTaskQueryParams) error {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s", url.PathEscape(taskId))

	q := url.Values{}

	if params.CustomTaskIds != nil {
		q.Set("custom_task_ids", fmt.Sprintf("%t", *params.CustomTaskIds))
	}
	if params.TeamId != nil {
		q.Set("team_id", fmt.Sprintf("%d", *params.TeamId))
	}

	req, err := http.NewRequest("DELETE", baseURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	return nil
}

func (c *ClickUpClient) MergeTasks(taskId string, reqBody cu.MergeTasksOptions) (*cu.Task, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/merge", url.PathEscape(taskId))

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

	var task cu.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *ClickUpClient) GetTaskTimeInStatus(taskId string, params cu.GetTaskTimeInStatusQueryParams) (*cu.TaskTimeInStatusResponse, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/task/%s/time_in_status", url.PathEscape(taskId))

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

	var taskTimeInStatus cu.TaskTimeInStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&taskTimeInStatus); err != nil {
		return nil, err
	}

	return &taskTimeInStatus, nil
}

func (c *ClickUpClient) GetBulkTaskTimeInStatus(taskId string, params cu.GetBulkTaskTimeInStatusQueryParams) (*cu.BulkStatusResponse, error) {
	baseURL := "https://api.clickup.com/api/v2/task/bulk_time_in_status/task_ids"

	q := url.Values{}

	if params.CustomTaskIds != nil {
		q.Set("custom_task_ids", fmt.Sprintf("%t", *params.CustomTaskIds))
	}
	if params.TeamId != nil {
		q.Set("team_id", fmt.Sprintf("%d", *params.TeamId))
	}
	for _, taskId := range params.TaskIds {
		q.Add("task_ids", taskId)
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

	var bulkTaskTimeInStatus cu.BulkStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&bulkTaskTimeInStatus); err != nil {
		return nil, err
	}

	return &bulkTaskTimeInStatus, nil
}

func (c *ClickUpClient) CreateTaskFromTemplate(listId, templateId string, reqBody cu.CreateTaskFromTemplateOptions) (*cu.Task, error) {
	baseURL := fmt.Sprintf("https://api.clickup.com/api/v2/list/%s/taskTemplate/%s", url.PathEscape(listId), url.PathEscape(templateId))

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

	var task cu.Task
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		return nil, err
	}

	return &task, nil
}

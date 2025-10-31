package cu

type CreateTaskOptions struct {
	Name                      string                 `json:"name"` // required
	Description               string                 `json:"description,omitempty"`
	Assignees                 []int                  `json:"assignees,omitempty"`
	Archived                  bool                   `json:"archived,omitempty"`
	GroupAssignees            bool                   `json:"group_assignees,omitempty"`
	Tags                      []string               `json:"tags,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	Priority                  int                    `json:"priority,omitempty"`
	DueDate                   int64                  `json:"due_date,omitempty"`
	DueDateTime               bool                   `json:"due_date_time,omitempty"`
	StartDate                 int64                  `json:"start_date,omitempty"`
	StartDateTime             bool                   `json:"start_date_time,omitempty"`
	Points                    int                    `json:"points,omitempty"`
	NotifyAll                 bool                   `json:"notify_all,omitempty"`
	Parent                    string                 `json:"parent,omitempty"`
	MarkdownContent           string                 `json:"markdown_content,omitempty"`
	LinksTo                   string                 `json:"links_to,omitempty"`
	CheckRequiredCustomFields bool                   `json:"check_required_custom_fields,omitempty"`
	CustomFields              map[string]interface{} `json:"custom_fields,omitempty"`
	CustomItemId              int                    `json:"custom_item_id,omitempty"`
}

type MergeTasksOptions struct {
	SourceTaskIds []string `json:"source_task_ids"` // required
}

type CreateTaskFromTemplateOptions struct {
	Name string `json:"name"` // required
}

type GetAccessTokenOptions struct {
	ClientId     string `json:"client_id"`     // required
	ClientSecret string `json:"client_secret"` // required
	Code         string `json:"code"`          // required
}

type CreateTaskCommentOptions struct {
	CommentText   string `json:"comment_text"` // required
	Assignee      int    `json:"assignee,omitempty"`
	GroupAssignee int    `json:"group_assignee,omitempty"`
	NotifyAll     bool   `json:"notify_all,omitempty"`
}

type CreateChatViewCommentOptions struct {
	CommentText string `json:"comment_text"` // required
	NotifyAll   bool   `json:"notify_all,omitempty"`
}

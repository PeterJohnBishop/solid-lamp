package cu

type GetTaskQueryParams struct {
	CustomTaskIds              *bool               `json:"custom_task_ids,omitempty"`
	TeamId                     *int                `json:"team_id,omitempty"`
	IncludeMarkdownDescription *bool               `json:"include_markdown_description,omitempty"`
	IncludeSubtasks            *bool               `json:"include_subtasks,omitempty"`
	CustomFields               []map[string]string `json:"custom_fields,omitempty"`
}

func SetGetTaskQueryParams(customTaskIds, includeMarkdown, includeSubtasks bool,
	teamId int,
	customFields ...[]map[string]string,
) *GetTaskQueryParams {
	var fields []map[string]string
	if len(customFields) > 0 {
		fields = customFields[0]
	}

	return &GetTaskQueryParams{
		CustomTaskIds:              &customTaskIds,
		TeamId:                     &teamId,
		IncludeMarkdownDescription: &includeMarkdown,
		IncludeSubtasks:            &includeSubtasks,
		CustomFields:               fields,
	}
}

type GetTasksQueryParams struct {
	Archived                   *bool               `json:"archived,omitempty"`
	IncludeMarkdownDescription *bool               `json:"include_markdown_description,omitempty"`
	Page                       *int                `json:"page,omitempty"`
	OrderBy                    *string             `json:"order_by,omitempty"`
	Reverse                    *bool               `json:"reverse,omitempty"`
	Subtasks                   *bool               `json:"subtasks,omitempty"`
	Statuses                   []string            `json:"statuses,omitempty"`
	IncludeClosed              *bool               `json:"include_closed,omitempty"`
	IncludeTiml                *bool               `json:"include_timl,omitempty"`
	Assignees                  []string            `json:"assignees,omitempty"`
	Watchers                   []string            `json:"watchers,omitempty"`
	Tags                       []string            `json:"tags,omitempty"`
	DueDateGt                  *int64              `json:"due_date_gt,omitempty"`
	DueDateLt                  *int64              `json:"due_date_lt,omitempty"`
	DateCreatedGt              *int64              `json:"date_created_gt,omitempty"`
	DateCreatedLt              *int64              `json:"date_created_lt,omitempty"`
	DateUpdatedGt              *int64              `json:"date_updated_gt,omitempty"`
	DateUpdatedLt              *int64              `json:"date_updated_lt,omitempty"`
	DateDoneGt                 *int64              `json:"date_done_gt,omitempty"`
	DateDoneLt                 *int64              `json:"date_done_lt,omitempty"`
	CustomFields               []map[string]string `json:"custom_fields,omitempty"`
	CustomField                []string            `json:"custom_field,omitempty"`
	CustomItems                []int               `json:"custom_items,omitempty"`
}

type GetFilteredTeamTasksQueryParams struct {
	IncludeMarkdownDescription *bool               `json:"include_markdown_description,omitempty"`
	Page                       *int                `json:"page,omitempty"`
	OrderBy                    *string             `json:"order_by,omitempty"`
	Reverse                    *bool               `json:"reverse,omitempty"`
	Subtasks                   *bool               `json:"subtasks,omitempty"`
	SpaceIds                   []string            `json:"space_ids,omitempty"`
	ProjectIds                 []string            `json:"project_ids,omitempty"`
	ListIds                    []string            `json:"list_ids,omitempty"`
	Statuses                   []string            `json:"statuses,omitempty"`
	IncludeClosed              *bool               `json:"include_closed,omitempty"`
	Assignees                  []string            `json:"assignees,omitempty"`
	Tags                       []string            `json:"tags,omitempty"`
	DueDateGt                  *int64              `json:"due_date_gt,omitempty"`
	DueDateLt                  *int64              `json:"due_date_lt,omitempty"`
	DateCreatedGt              *int64              `json:"date_created_gt,omitempty"`
	DateCreatedLt              *int64              `json:"date_created_lt,omitempty"`
	DateUpdatedGt              *int64              `json:"date_updated_gt,omitempty"`
	DateUpdatedLt              *int64              `json:"date_updated_lt,omitempty"`
	DateDoneGt                 *int64              `json:"date_done_gt,omitempty"`
	DateDoneLt                 *int64              `json:"date_done_lt,omitempty"`
	CustomFields               []map[string]string `json:"custom_fields,omitempty"`
	Parent                     string              `json:"parent,omitempty"`
	CustomItems                []int               `json:"custom_items,omitempty"`
}

type UpdateTaskQueryParams struct {
	CustomTaskIds *bool `json:"custom_task_ids,omitempty"`
	TeamId        *int  `json:"team_id,omitempty"`
}

type DeleteTaskQueryParams struct {
	CustomTaskIds *bool `json:"custom_task_ids,omitempty"`
	TeamId        *int  `json:"team_id,omitempty"`
}

type GetTaskTimeInStatusQueryParams struct {
	CustomTaskIds *bool `json:"custom_task_ids,omitempty"`
	TeamId        *int  `json:"team_id,omitempty"`
}

type GetBulkTaskTimeInStatusQueryParams struct {
	CustomTaskIds *bool    `json:"custom_task_ids,omitempty"`
	TeamId        *int     `json:"team_id,omitempty"`
	TaskIds       []string `json:"task_ids,omitempty"`
}

type GetTaskCommentsQueryParams struct {
	CustomTaskIds *bool   `json:"custom_task_ids,omitempty"`
	TeamId        *int    `json:"team_id,omitempty"`
	Start         *int    `json:"start,omitempty"`
	StartId       *string `json:"start_id,omitempty"`
}

type CreateTaskCommentQueryParams struct {
	CustomTaskIds *bool `json:"custom_task_ids,omitempty"`
	TeamId        *int  `json:"team_id,omitempty"`
}

type GetChatViewCommentsQueryParams struct {
	Start   *int64  `json:"start,omitempty"`
	StartId *string `json:"start_id,omitempty"`
}

type GetListCommentsQueryParams struct {
	Start   *int64  `json:"start,omitempty"`
	StartId *string `json:"start_id,omitempty"`
}

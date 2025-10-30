package solidlamp

type List struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Access bool   `json:"access"`
}

type Watcher struct {
	Email          string `json:"email"`
	ProfilePicture any    `json:"profilePicture"`
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Color          any    `json:"color"`
	Initials       string `json:"initials"`
}

type Project struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Access bool   `json:"access"`
}

type Folder struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Hidden bool   `json:"hidden"`
	Access bool   `json:"access"`
}

type Space struct {
	Id string `json:"id"`
}

type CustomField struct {
	Name           string     `json:"name"`
	HideFromGuests bool       `json:"hide_from_guests"`
	ValueRichtext  any        `json:"value_richtext"`
	Id             string     `json:"id"`
	Type           string     `json:"type"`
	TypeConfig     TypeConfig `json:"type_config"`
	DateCreated    string     `json:"date_created"`
	Value          string     `json:"value"`
	Required       bool       `json:"required"`
}

type Status struct {
	Orderindex int    `json:"orderindex"`
	Type       string `json:"type"`
	Id         string `json:"id"`
	Status     string `json:"status"`
	Color      string `json:"color"`
}

type Creator struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	Color          any    `json:"color"`
	Email          string `json:"email"`
	ProfilePicture any    `json:"profilePicture"`
}

type Sharing struct {
	Public               bool     `json:"public"`
	PublicShareExpiresOn any      `json:"public_share_expires_on"`
	PublicFields         []string `json:"public_fields"`
	Token                any      `json:"token"`
	SeoOptimized         bool     `json:"seo_optimized"`
}

type Task struct {
	Id              string        `json:"id"`
	DateCreated     string        `json:"date_created"`
	Checklists      []any         `json:"checklists"`
	DueDate         string        `json:"due_date"`
	TimeSpent       int           `json:"time_spent"`
	List            List          `json:"list"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	TopLevelParent  any           `json:"top_level_parent"`
	Url             string        `json:"url"`
	PermissionLevel string        `json:"permission_level"`
	CustomItemId    int           `json:"custom_item_id"`
	DateDone        any           `json:"date_done"`
	GroupAssignees  []any         `json:"group_assignees"`
	Watchers        []Watcher     `json:"watchers"`
	Parent          any           `json:"parent"`
	Project         Project       `json:"project"`
	TextContent     string        `json:"text_content"`
	Assignees       []any         `json:"assignees"`
	TimeEstimate    any           `json:"time_estimate"`
	Folder          Folder        `json:"folder"`
	CustomId        any           `json:"custom_id"`
	Archived        bool          `json:"archived"`
	Priority        any           `json:"priority"`
	Points          any           `json:"points"`
	TeamId          string        `json:"team_id"`
	Space           Space         `json:"space"`
	DateClosed      any           `json:"date_closed"`
	Tags            []any         `json:"tags"`
	StartDate       any           `json:"start_date"`
	CustomFields    []CustomField `json:"custom_fields"`
	Dependencies    []any         `json:"dependencies"`
	LinkedTasks     []any         `json:"linked_tasks"`
	Locations       []any         `json:"locations"`
	Status          Status        `json:"status"`
	Orderindex      string        `json:"orderindex"`
	DateUpdated     string        `json:"date_updated"`
	Creator         Creator       `json:"creator"`
	Sharing         Sharing       `json:"sharing"`
	Attachments     []any         `json:"attachments"`
}

type Tasks struct {
	Tasks    []Task `json:"tasks"`
	LastPage bool   `json:"last_page"`
}

type Option struct {
	ID         string `json:"id"`
	Name       string `json:"name,omitempty"`
	Label      string `json:"label,omitempty"`
	Color      string `json:"color"`
	OrderIndex int    `json:"orderindex"`
}

type ButtonIcon struct {
	Value  string `json:"value"`
	Source string `json:"source"`
}

type ButtonAutomation struct {
	ID string `json:"id"`
}

type TypeConfig struct {
	// Sorting & options
	Sorting *string  `json:"sorting,omitempty"`
	Options []Option `json:"options,omitempty"`

	// Currency / precision
	Precision    *int    `json:"precision,omitempty"`
	CurrencyType *string `json:"currency_type,omitempty"`

	// Fields
	Fields []string `json:"fields,omitempty"`

	// User inclusion flags
	IncludeGroups      *bool `json:"include_groups,omitempty"`
	IncludeGuests      *bool `json:"include_guests,omitempty"`
	IncludeTeamMembers *bool `json:"include_team_members,omitempty"`
	SingleUser         *bool `json:"single_user,omitempty"`

	// Count & code_point
	Count     *int    `json:"count,omitempty"`
	CodePoint *string `json:"code_point,omitempty"`

	// Range
	Start *int `json:"start,omitempty"`
	End   *int `json:"end,omitempty"`

	// Formula / calculation
	Simple           *bool    `json:"simple,omitempty"`
	Formula          *string  `json:"formula,omitempty"`
	Version          *string  `json:"version,omitempty"`
	ResetAt          *int64   `json:"reset_at,omitempty"`
	IsDynamic        *bool    `json:"is_dynamic,omitempty"`
	ReturnTypes      []string `json:"return_types,omitempty"`
	CalculationState *string  `json:"calculation_state,omitempty"`

	// Button config
	ButtonIcon       *ButtonIcon       `json:"button_icon,omitempty"`
	ButtonText       *string           `json:"button_text,omitempty"`
	ButtonColor      *string           `json:"button_color,omitempty"`
	ButtonAutomation *ButtonAutomation `json:"button_automation,omitempty"`
}

type GetTaskQueryParams struct {
	Archived                   *bool               `json:"archived,omitempty"`
	IncludeMarkdownDescription *bool               `json:"include_markdown_description,omitempty"`
	Page                       *int                `json:"page,omitempty"`
	OrderBy                    *string             `json:"order_by,omitempty"`
	Reverse                    *bool               `json:"reverse,omitempty"`
	Subtasks                   *bool               `json:"subtasks,omitempty"`
	Statuses                   []string            `json:"statuses,omitempty"` //?statuses[]=to%20do&statuses[]=in%20progress
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

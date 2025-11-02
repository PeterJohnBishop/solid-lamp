package cu

type List struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Access *bool   `json:"access,omitempty"`
}

type Watcher struct {
	Email          *string `json:"email,omitempty"`
	ProfilePicture any     `json:"profilePicture,omitempty"`
	Id             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Color          any     `json:"color,omitempty"`
	Initials       *string `json:"initials,omitempty"`
}

type Project struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Hidden *bool   `json:"hidden,omitempty"`
	Access *bool   `json:"access,omitempty"`
}

type Folder struct {
	Id     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Hidden *bool   `json:"hidden,omitempty"`
	Access *bool   `json:"access,omitempty"`
}

type Space struct {
	Id *string `json:"id,omitempty"`
}

type CustomField struct {
	Name           *string     `json:"name,omitempty"`
	HideFromGuests *bool       `json:"hide_from_guests,omitempty"`
	ValueRichtext  any         `json:"value_richtext,omitempty"`
	Id             *string     `json:"id,omitempty"`
	Type           *string     `json:"type,omitempty"`
	TypeConfig     *TypeConfig `json:"type_config,omitempty"`
	DateCreated    *string     `json:"date_created,omitempty"`
	Value          *string     `json:"value,omitempty"`
	Required       *bool       `json:"required,omitempty"`
}

type Status struct {
	Orderindex *int    `json:"orderindex,omitempty"`
	Type       *string `json:"type,omitempty"`
	Id         *string `json:"id,omitempty"`
	Status     *string `json:"status,omitempty"`
	Color      *string `json:"color,omitempty"`
}

type Creator struct {
	Id             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Color          any     `json:"color,omitempty"`
	Email          *string `json:"email,omitempty"`
	ProfilePicture any     `json:"profilePicture,omitempty"`
}

type Sharing struct {
	Public               *bool    `json:"public,omitempty"`
	PublicShareExpiresOn any      `json:"public_share_expires_on,omitempty"`
	PublicFields         []string `json:"public_fields,omitempty"`
	Token                any      `json:"token,omitempty"`
	SeoOptimized         *bool    `json:"seo_optimized,omitempty"`
}

type Task struct {
	Id              *string        `json:"id,omitempty"`
	DateCreated     *string        `json:"date_created,omitempty"`
	Checklists      []any          `json:"checklists,omitempty"`
	DueDate         *string        `json:"due_date,omitempty"`
	TimeSpent       *int           `json:"time_spent,omitempty"`
	List            *List          `json:"list,omitempty"`
	Name            *string        `json:"name,omitempty"`
	Description     *string        `json:"description,omitempty"`
	TopLevelParent  any            `json:"top_level_parent,omitempty"`
	Url             *string        `json:"url,omitempty"`
	PermissionLevel *string        `json:"permission_level,omitempty"`
	CustomItemId    *int           `json:"custom_item_id,omitempty"`
	DateDone        any            `json:"date_done,omitempty"`
	GroupAssignees  []any          `json:"group_assignees,omitempty"`
	Watchers        []*Watcher     `json:"watchers,omitempty"`
	Parent          any            `json:"parent,omitempty"`
	Project         *Project       `json:"project,omitempty"`
	TextContent     *string        `json:"text_content,omitempty"`
	Assignees       []any          `json:"assignees,omitempty"`
	TimeEstimate    any            `json:"time_estimate,omitempty"`
	Folder          *Folder        `json:"folder,omitempty"`
	CustomId        any            `json:"custom_id,omitempty"`
	Archived        *bool          `json:"archived,omitempty"`
	Priority        any            `json:"priority,omitempty"`
	Points          any            `json:"points,omitempty"`
	TeamId          *string        `json:"team_id,omitempty"`
	Space           *Space         `json:"space,omitempty"`
	DateClosed      any            `json:"date_closed,omitempty"`
	Tags            []any          `json:"tags,omitempty"`
	StartDate       any            `json:"start_date,omitempty"`
	CustomFields    []*CustomField `json:"custom_fields,omitempty"`
	Dependencies    []any          `json:"dependencies,omitempty"`
	LinkedTasks     []any          `json:"linked_tasks,omitempty"`
	Locations       []any          `json:"locations,omitempty"`
	Status          *Status        `json:"status,omitempty"`
	Orderindex      *string        `json:"orderindex,omitempty"`
	DateUpdated     *string        `json:"date_updated,omitempty"`
	Creator         *Creator       `json:"creator,omitempty"`
	Sharing         *Sharing       `json:"sharing,omitempty"`
	Attachments     []any          `json:"attachments,omitempty"`
}

type Tasks struct {
	Tasks    []*Task `json:"tasks,omitempty"`
	LastPage *bool   `json:"last_page,omitempty"`
}

type Option struct {
	ID         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	Label      *string `json:"label,omitempty"`
	Color      *string `json:"color,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
}

type ButtonIcon struct {
	Value  *string `json:"value,omitempty"`
	Source *string `json:"source,omitempty"`
}

type ButtonAutomation struct {
	ID *string `json:"id,omitempty"`
}

type TypeConfig struct {
	Sorting            *string           `json:"sorting,omitempty"`
	Options            []*Option         `json:"options,omitempty"`
	Precision          *int              `json:"precision,omitempty"`
	CurrencyType       *string           `json:"currency_type,omitempty"`
	Fields             []string          `json:"fields,omitempty"`
	IncludeGroups      *bool             `json:"include_groups,omitempty"`
	IncludeGuests      *bool             `json:"include_guests,omitempty"`
	IncludeTeamMembers *bool             `json:"include_team_members,omitempty"`
	SingleUser         *bool             `json:"single_user,omitempty"`
	Count              *int              `json:"count,omitempty"`
	CodePoint          *string           `json:"code_point,omitempty"`
	Start              *int              `json:"start,omitempty"`
	End                *int              `json:"end,omitempty"`
	Simple             *bool             `json:"simple,omitempty"`
	Formula            *string           `json:"formula,omitempty"`
	Version            *string           `json:"version,omitempty"`
	ResetAt            *int64            `json:"reset_at,omitempty"`
	IsDynamic          *bool             `json:"is_dynamic,omitempty"`
	ReturnTypes        []string          `json:"return_types,omitempty"`
	CalculationState   *string           `json:"calculation_state,omitempty"`
	ButtonIcon         *ButtonIcon       `json:"button_icon,omitempty"`
	ButtonText         *string           `json:"button_text,omitempty"`
	ButtonColor        *string           `json:"button_color,omitempty"`
	ButtonAutomation   *ButtonAutomation `json:"button_automation,omitempty"`
}

type TotalTime struct {
	ByMinute *int    `json:"by_minute,omitempty"`
	Since    *string `json:"since,omitempty"`
}

type CurrentStatus struct {
	TotalTime *TotalTime `json:"total_time,omitempty"`
	Status    *string    `json:"status,omitempty"`
	Color     *string    `json:"color,omitempty"`
}

type StatusHistory struct {
	Status     *string    `json:"status,omitempty"`
	Color      *string    `json:"color,omitempty"`
	Type       *string    `json:"type,omitempty"`
	TotalTime  *TotalTime `json:"total_time,omitempty"`
	Orderindex *int       `json:"orderindex,omitempty"`
}

type BulkTaskTimeInStatus struct {
	CurrentStatus *CurrentStatus   `json:"current_status,omitempty"`
	StatusHistory []*StatusHistory `json:"status_history,omitempty"`
}

type Comment struct {
	Reactions   []any       `json:"reactions,omitempty"`
	Date        *string     `json:"date,omitempty"`
	Comment     []*Comment  `json:"comment,omitempty"`
	Assignee    *Assignee   `json:"assignee,omitempty"`
	AssignedBy  *AssignedBy `json:"assigned_by,omitempty"`
	Resolved    *bool       `json:"resolved,omitempty"`
	ReplyCount  *string     `json:"reply_count,omitempty"`
	Id          *string     `json:"id,omitempty"`
	CommentText *string     `json:"comment_text,omitempty"`
	User        *User       `json:"user,omitempty"`
}

type Assignee struct {
	Id             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Initials       *string `json:"initials,omitempty"`
	Email          *string `json:"email,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
}

type AssignedBy struct {
	ProfilePicture *string `json:"profilePicture,omitempty"`
	Id             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Initials       *string `json:"initials,omitempty"`
	Email          *string `json:"email,omitempty"`
	Color          *string `json:"color,omitempty"`
}

type User struct {
	Username       *string `json:"username,omitempty"`
	Initials       *string `json:"initials,omitempty"`
	Email          *string `json:"email,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
	Id             *int    `json:"id,omitempty"`
}

type ChatViewComment struct {
	Comment     []*ChatComment `json:"comment,omitempty"`
	CommentText *string        `json:"comment_text,omitempty"`
	Resolved    *bool          `json:"resolved,omitempty"`
	Date        *string        `json:"date,omitempty"`
	ReplyCount  *string        `json:"reply_count,omitempty"`
	Id          *string        `json:"id,omitempty"`
	User        *User          `json:"user,omitempty"`
	Assignee    any            `json:"assignee,omitempty"`
	AssignedBy  any            `json:"assigned_by,omitempty"`
	Reactions   []any          `json:"reactions,omitempty"`
}

type ChatComment struct {
	Text *string `json:"text,omitempty"`
}

type ListComment struct {
	Id          *string        `json:"id,omitempty"`
	CommentText *string        `json:"comment_text,omitempty"`
	Assignee    *Assignee      `json:"assignee,omitempty"`
	ReplyCount  *string        `json:"reply_count,omitempty"`
	Comment     []*ChatComment `json:"comment,omitempty"`
	User        *User          `json:"user,omitempty"`
	Resolved    *bool          `json:"resolved,omitempty"`
	AssignedBy  *AssignedBy    `json:"assigned_by,omitempty"`
	Reactions   []any          `json:"reactions,omitempty"`
	Date        *string        `json:"date,omitempty"`
}

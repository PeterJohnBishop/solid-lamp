package solidlamp

type TaskTimeInStatusResponse struct {
	CurrentStatus CurrentStatus   `json:"current_status"`
	StatusHistory []StatusHistory `json:"status_history"`
}

type BulkStatusResponse map[string]TaskTimeInStatusResponse

type AuthorizedUser struct {
	Initials          string `json:"initials"`
	WeekStartDay      int    `json:"week_start_day"`
	GlobalFontSupport bool   `json:"global_font_support"`
	Timezone          string `json:"timezone"`
	Email             string `json:"email"`
	Color             string `json:"color"`
	ProfilePicture    string `json:"profilePicture"`
	Id                int    `json:"id"`
	Username          string `json:"username"`
}

type AuthorizedUserResponse map[string]AuthorizedUser

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

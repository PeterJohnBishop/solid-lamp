package solidlamp

type TaskTimeInStatusResponse struct {
	CurrentStatus CurrentStatus   `json:"current_status"`
	StatusHistory []StatusHistory `json:"status_history"`
}

type BulkStatusResponse map[string]TaskTimeInStatusResponse

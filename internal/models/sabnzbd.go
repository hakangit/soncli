package models

// SabnzbdResponse represents the base response from Sabnzbd API
type SabnzbdResponse struct {
	Status bool   `json:"status"`
	Error  string `json:"error,omitempty"`
}

// QueueResponse represents the queue response
type QueueResponse struct {
	Queue Queue `json:"queue"`
}

// Queue represents the download queue
type Queue struct {
	Version    string      `json:"version"`
	Paused     bool        `json:"paused"`
	PauseInt   string      `json:"pause_int"`
	SpeedLimit string      `json:"speedlimit"`
	Speed      string      `json:"speed"`
	Size       string      `json:"size"`
	SizeLeft   string      `json:"sizeleft"`
	TimeLeft   string      `json:"timeleft"`
	ETA        string      `json:"eta"`
	Status     string      `json:"status"`
	Slots      []QueueSlot `json:"slots"`
}

// QueueSlot represents a slot in the queue
type QueueSlot struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Category   string `json:"cat"`
	Size       string `json:"size"`
	SizeLeft   string `json:"sizeleft"`
	TimeLeft   string `json:"timeleft"`
	ETA        string `json:"eta"`
	Status     string `json:"status"`
	Index      int    `json:"index"`
	Percentage string `json:"percentage"`
	Missing    int    `json:"missing"`
}

// HistoryResponse represents the history response
type HistoryResponse struct {
	History History `json:"history"`
}

// History represents the download history
type History struct {
	Version string        `json:"version"`
	Paused  bool          `json:"paused"`
	Slots   []HistorySlot `json:"slots"`
}

// HistorySlot represents a slot in the history
type HistorySlot struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"cat"`
	Size        string `json:"size"`
	Status      string `json:"status"`
	ActionLine  string `json:"action_line"`
	ShowDetails bool   `json:"show_details"`
	ScriptLog   string `json:"script_log"`
	Meta        string `json:"meta"`
	Completed   int64  `json:"completed"`
	Downloaded  int64  `json:"downloaded"`
	FailMessage string `json:"fail_message"`
	URL         string `json:"url"`
	Bytes       int64  `json:"bytes"`
	Credentials string `json:"credentials"`
	PP          string `json:"pp"`
	Script      string `json:"script"`
	NzbName     string `json:"nzb_name"`
	Path        string `json:"path"`
	Storage     string `json:"storage"`
}

// AddResponse represents the response from adding an NZB
type AddResponse struct {
	SabnzbdResponse
	NZOIDS []string `json:"nzo_ids"`
}

// CategoriesResponse represents the categories response
type CategoriesResponse struct {
	Categories []string `json:"categories"`
}

// VersionResponse represents the version response
type VersionResponse struct {
	SabnzbdResponse
	Version string `json:"version"`
}

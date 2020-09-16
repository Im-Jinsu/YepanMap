package cerror

// Error Messages
const (
	ErrAuthFailed   = "Authentication Failed"
	ErrInvalidParam = "Invalid Parameters"
	ErrInvalidID    = "Invalid ID"
	ErrNotFound     = "Not Found"
	ErrDBFailed     = "DB maintenance in progress we apologize for the inconvenience"
)

// Message is api error message format
type Message struct {
	Msg string `json:"msg" xml:"msg"`
}

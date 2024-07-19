package contract

type Response struct {
	RequestID string `json:"request_id,omitempty"`
	Message   string `json:"message,omitempty"`
}

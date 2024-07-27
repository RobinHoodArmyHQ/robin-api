package models

// Status denotes the response status
type Status struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// StatusSuccess is a helper function to create a successful Status object
func StatusSuccess() *Status {
	return &Status{
		Success: true,
	}
}

// StatusFailed is a helper function to create a unsuccessful Status object
func StatusFailed(message string) *Status {
	return &Status{
		Success: false,
		Message: message,
	}
}

// StatusSomethingWentWrong is a helper function to create a generic failed Status object
func StatusSomethingWentWrong() *Status {
	return StatusFailed("something went wrong")
}

// StatusTimeout is a helper function to create a timed out Status object
func StatusTimedOut() *Status {
	return StatusFailed("timed out")
}

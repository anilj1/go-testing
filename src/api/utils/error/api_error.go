package error

type ApiError struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
	// Cause   []string      `json: "cause"`
}

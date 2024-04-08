package abstractions

import "fmt"

// ApiError is the parent type for errors thrown by the client when receiving failed responses to its requests
type ApiError struct {
	Message string
}

func (e *ApiError) Error() string {
	if len(e.Message) > 0 {
		return fmt.Sprint(e.Message)
	} else {
		return "error status code received from the API"
	}
}

// NewApiError creates a new ApiError instance
func NewApiError() *ApiError {
	return &ApiError{}
}

package shared

type APIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func SuccessResponse(message string, data any) APIResponse {
	return APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err any) APIResponse {
	return APIResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	}
}

package utils

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func ErrorResponse(error string, message string) Response {
	return Response{
		Success: false,
		Error:   error,
		Message: message,
	}
}

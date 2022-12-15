package Model

type ErrorResponse struct {
	Message interface{} `json:"message"`
	Error   interface{} `json:"error"`
}

type DefaultResponse struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(message interface{}, error interface{}) *ErrorResponse {
	return &ErrorResponse{Message: message, Error: error}
}

func NewDefaultResponse(message string, data interface{}) *DefaultResponse {
	return &DefaultResponse{Message: message, Data: data}
}

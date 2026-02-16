package utils

type SuccessResponse struct {
    Success bool        `json:"success"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
    Success bool              `json:"success"`
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}
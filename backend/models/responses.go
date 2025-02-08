package models

// Response represents a generic response
type Response struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Error message"`
}

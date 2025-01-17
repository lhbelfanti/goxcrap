package response

// DTO represents a standard HTTP response for a RESTful API
type DTO struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

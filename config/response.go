package config

import "github.com/gin-gonic/gin"

// Response estructura para respuestas JSON estándar
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse estructura para respuestas de error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// SendSuccess envía una respuesta exitosa
func SendSuccess(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SendError envía una respuesta de error
func SendError(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, ErrorResponse{
		Success: false,
		Error:   errorMessage,
	})
}

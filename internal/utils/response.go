package utils

import "github.com/gofiber/fiber/v2"

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SendSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SendError(c *fiber.Ctx, statusCode int, message, errorDetail string) error {
	return c.Status(statusCode).JSON(APIResponse{
		Status:  "error",
		Message: message,
		Error:   errorDetail,
	})
}

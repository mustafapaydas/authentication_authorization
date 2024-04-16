package utils

import (
	"github.com/gofiber/fiber/v2"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorHandler(c *fiber.Ctx, err error) error {
	response := CustomError{
		Code:    fiber.StatusInternalServerError,
		Message: err.Error(),
	}

	switch err.(type) {
	case *UniqueException:
		response.Code = fiber.StatusConflict
	case *NotNullException:
		response.Code = fiber.StatusNotAcceptable
	default:
		response.Message = err.Error()
	}

	return c.Status(response.Code).JSON(response)
}

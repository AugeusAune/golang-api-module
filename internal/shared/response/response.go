package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *fiber.Ctx, data interface{}, message string) error {
	return c.JSON(&Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Created(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusCreated).JSON(&Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(&Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

func BadRequest(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusBadRequest, message)
}

func Unauthorized(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusUnauthorized, message)
}

func NotFound(c *fiber.Ctx, message string) error {
	return Error(c, fiber.StatusNotFound, message)
}

func ErrorValidation(c *fiber.Ctx, err error) error {
	errs := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		var message string

		switch e.Tag() {
		case "required":
			message = e.Field() + " is required"
		case "email":
			message = e.Field() + " must be a valid email"
		case "min":
			message = e.Field() + " must be at least " + e.Param() + " characters"
		case "max":
			message = e.Field() + " max value is " + e.Param() + "characters"
		default:
			message = e.Field() + " is invalid"
		}

		errs[e.Field()] = message
	}

	return c.Status(fiber.StatusBadRequest).JSON(&Response{
		Success: false,
		Message: "Validation error",
		Data:    errs,
	})
}

func ErrorBodyParser(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).JSON(&Response{
		Success: false,
		Message: err.Error(),
	})
}

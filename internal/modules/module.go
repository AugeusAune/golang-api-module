package modules

import (
	"golang-api-module/internal/modules/vatrate"
	"golang-api-module/internal/shared/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InitModule(app *fiber.App, log *logrus.Logger) {
	router := app.Group("/api")

	validator := validator.New()

	registerVatRateModule(router, validator, log)

	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "Route not found")
	})
}

func registerVatRateModule(router fiber.Router, validator *validator.Validate, log *logrus.Logger) {
	handler := vatrate.NewHandler(validator, log)
	module := vatrate.NewModule(handler)

	module.RegisterRoutes(router.Group("/vat-rate"))
}

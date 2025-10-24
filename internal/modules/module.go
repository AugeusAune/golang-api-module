package modules

import (
	"context"
	"golang-api-module/internal/modules/vatrate"
	"golang-api-module/internal/queue"
	"golang-api-module/internal/shared/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func InitModule(ctx context.Context, app *fiber.App, queueClient *queue.Client, log *logrus.Logger) {
	router := app.Group("/api")

	validate := validator.New()

	registerVatRateModule(ctx, router, queueClient, validate, log)

	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "Route not found")
	})
}

func registerVatRateModule(ctx context.Context, router fiber.Router, queueClient *queue.Client, validate *validator.Validate, log *logrus.Logger) {
	service := vatrate.NewService(ctx, queueClient, log)
	handler := vatrate.NewHandler(service, validate, log)
	module := vatrate.NewModule(handler)

	module.RegisterRoutes(router.Group("/vat-rate"))
}

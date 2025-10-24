package modules

import (
	"context"
	"golang-api-module/internal/modules/vatrate"
	"golang-api-module/internal/queue"
	"golang-api-module/internal/shared/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func InitModule(ctx context.Context, app *fiber.App, db *gorm.DB, queueClient *queue.Client, log *logrus.Logger) {
	router := app.Group("/api")

	validate := validator.New()

	registerVatRateModule(ctx, router, db, queueClient, validate, log)

	app.Use(func(c *fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "Route not found")
	})
}

func registerVatRateModule(ctx context.Context, router fiber.Router, db *gorm.DB, queueClient *queue.Client, validate *validator.Validate, log *logrus.Logger) {
	service := vatrate.NewService(ctx, db, queueClient, log)
	handler := vatrate.NewHandler(service, validate, log)
	module := vatrate.NewModule(handler)

	module.RegisterRoutes(router.Group("/vat-rate"))
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-api-module/config"
	"golang-api-module/internal/modules"
	internalLog "golang-api-module/internal/shared/logger"
	"golang-api-module/internal/shared/response"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log := internalLog.NewLogger()

	if err := godotenv.Load(); err != nil {
		log.Warn("no .env found")
	}

	config := config.Load()

	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: customErrorHandler,
		Immutable:    true,
	})

	app.Use(requestid.New())
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path}\n",
		Output: io.MultiWriter(os.Stdout),
	}))

	modules.InitModule(app, log)

	go func() {
		port := fmt.Sprintf(":%s", config.Port)
		if err := app.Listen(port); err != nil {
			log.Fatalf("Failed Start app at port %s", port)
		}
	}()

	<-sigChan

	log.Info("shutting down...")

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Errorf("Error shutting down server: %v", err)
	}

	cancel()
}

func customErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(&response.Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

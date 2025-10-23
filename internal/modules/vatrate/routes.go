package vatrate

import (
	"github.com/gofiber/fiber/v2"
)

type Module struct {
	handler *Handler
}

func NewModule(handler *Handler) *Module {
	return &Module{
		handler: handler,
	}
}

func (m *Module) RegisterRoutes(router fiber.Router) {
	router.Get("/", m.handler.GetAll)
	router.Post("/", m.handler.Create)
}

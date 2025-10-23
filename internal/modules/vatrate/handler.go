package vatrate

import (
	"golang-api-module/internal/shared/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	validator *validator.Validate
	log       *logrus.Logger
}

func NewHandler(validator *validator.Validate, log *logrus.Logger) *Handler {
	return &Handler{
		validator: validator,
		log:       log,
	}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	return response.Success(c, nil, "routing from vat rate")
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req CreateVatRateRequest

	if err := c.BodyParser(&req); err != nil {
		h.log.Warn(err.Error())
		return response.ErrorBodyParser(c, err)
	}

	if err := h.validator.Struct(req); err != nil {
		h.log.Warn(err.Error())
		return response.ErrorValidation(c, err)
	}

	return response.Created(c, req, "VAT Rate berhasil dibuat")
}

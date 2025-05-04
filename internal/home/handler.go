package home

import (
	"bytes"
	"text/template"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

// /api/
// /api/error

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
	}
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	//"{{.Count}} - число пользователей"
	tmpl, err := template.New("test").Parse("{{.Count}} - число пользователей")
	data := struct{ Count int }{Count: 1}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Temlate error")
	}
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Temlate compile error")
	}
	return c.Send(tpl.Bytes())
}

func (h *HomeHandler) error(c *fiber.Ctx) error {

	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")

	return c.SendString("Error")
}

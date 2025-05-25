package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/tadapter"
	"go-fiber/views"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	//"github.com/rs/zerolog/log"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
}

type User struct {
	Id   int
	Name string
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *vacancy.VacancyRepository) {
	h := &HomeHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	//api := h.router.Group("/api")
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	vacancies, err := h.repository.GetAll()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(500)
	}
	component := views.Main(vacancies)
	//component.Render(context.Background(), os.Stdout)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {

	h.customLogger.Info().
		Bool("isAdmin", true).
		Str("email", "a@a.ru").
		Int("id", 10).
		Msg("Инфо")

	return c.SendString("Error")
}

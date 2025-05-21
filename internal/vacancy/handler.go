package vacancy

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *VacancRepository
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger, repository *VacancRepository) {
	h := &VacancyHandler{
		router:       router,
		customLogger: customLogger,
		repository:   repository,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacansy)
}

func (h *VacancyHandler) createVacansy(c *fiber.Ctx) error {
	form := VacancyCreateForm{
		Email:    c.FormValue("email"),
		Location: c.FormValue("location"),
		Type:     c.FormValue("type"),
		Company:  c.FormValue("company"),
		Role:     c.FormValue("role"),
		Salary:   c.FormValue("salary"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{Name: "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный"},
		&validators.StringIsPresent{Name: "Location",
			Field:   form.Location,
			Message: "Расположение не задано"},
		&validators.StringIsPresent{Name: "Type",
			Field:   form.Type,
			Message: "Сфера компаний не задана"},
		&validators.StringIsPresent{Name: "Company",
			Field:   form.Company,
			Message: "Название компании не задано"},
		&validators.StringIsPresent{Name: "Role",
			Field:   form.Role,
			Message: "Должность не задана"},
		&validators.StringIsPresent{Name: "Salary",
			Field:   form.Salary,
			Message: "Зароботная плата не задана"},
	)
	//time.Sleep(time.Second * 1)
	var component templ.Component
	if len(errors.Errors) > 0 {
		component = components.Notification(validator.FormatErrors(errors), components.NotificationFail)
		return tadapter.Render(c, component)
	}
	component = components.Notification("Вакансия успешно создана", components.NotificationSuccess)
	return tadapter.Render(c, component)
}

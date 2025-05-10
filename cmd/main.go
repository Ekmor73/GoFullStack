package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/pkg/logger"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()
	customLogger := logger.NewLogger(logConfig)
	// engine := html.New("./html", ".html")
	// engine.AddFuncMap(map[string]interface{}{
	// 	"ToUpper": func(c string) string {
	// 		return strings.ToUpper(c)
	// 	},
	// })

	app := fiber.New()

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: customLogger,
	}))
	app.Use(recover.New())
	app.Static("/public", "./public")

	home.NewHandler(app, customLogger)

	// Запускаем сервер на порту 3000
	app.Listen(":3000")
}

package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	// Создаем новый экземпляр Fiber
	app := fiber.New()
	app.Use(recover.New())

	home.NewHandler(app)

	// Запускаем сервер на порту 3000
	app.Listen(":3000")
}

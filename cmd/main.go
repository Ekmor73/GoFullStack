package main

import (
	"go-fiber/config"
	"go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)
	// Создаем новый экземпляр Fiber
	app := fiber.New()

	//Определяем обработчик для корневого маршрута
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Привет, Fiber!")
	// })
	home.NewHandler(app)

	// Запускаем сервер на порту 3000
	app.Listen(":3000")
}

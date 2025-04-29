package main

import (
	"go-fiber/internal/home"

	"github.com/gofiber/fiber/v2"
)

func main() {
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

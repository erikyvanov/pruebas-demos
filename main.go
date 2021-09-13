package main

import (
	"github.com/erikyvanov/pruebas-demos/routes"
	"github.com/erikyvanov/pruebas-demos/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	socketService := services.GetEventSocketService()
	go socketService.Run()

	app := fiber.New()
	routes.Router(app)

	app.Listen(":8080")
}

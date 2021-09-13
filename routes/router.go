package routes

import (
	"github.com/erikyvanov/pruebas-demos/handlers"
	"github.com/erikyvanov/pruebas-demos/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Router(app *fiber.App) {
	app.Static("/", "./public")

	app.Use("/ws", middlewares.IsWebSocketUpgrade)
	app.Use("/ws", middlewares.WebsocketCache())
	app.Get("/ws", websocket.New(handlers.WebsocketEventHandler))

	app.Get("/get_cache", handlers.GetCache)
}

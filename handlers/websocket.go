package handlers

import (
	"github.com/erikyvanov/pruebas-demos/models"
	"github.com/erikyvanov/pruebas-demos/services"
	"github.com/gofiber/websocket/v2"
)

func WebsocketEventHandler(c *websocket.Conn) {
	socketService := services.GetEventSocketService()

	var (
		err error
	)

	client := models.Client{
		Conn: c,
		ID:   c.Query("id"),
	}

	c.SetPingHandler(func(appData string) error {
		err := c.WriteMessage(websocket.PongMessage, nil)
		if err != nil {
			socketService.DeleteClient <- client.ID
			return err
		}
		return nil
	})

	socketService.NewClient <- client

	for {
		var msg models.Event

		if err = c.ReadJSON(&msg); err != nil {
			socketService.DeleteClient <- client.ID
			break
		}

		socketService.NewEvent <- msg

	}

}

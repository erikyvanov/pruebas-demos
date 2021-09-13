package models

import "github.com/gofiber/websocket/v2"

type Client struct {
	Conn *websocket.Conn
	ID   string
}

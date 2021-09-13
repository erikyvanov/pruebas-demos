package services

import (
	"sync"

	"github.com/erikyvanov/pruebas-demos/models"
	"github.com/gofiber/websocket/v2"
)

var onceSocketService sync.Once
var eventSocketService *EventSocketService

type EventSocketService struct {
	clients      map[string]*websocket.Conn
	NewClient    chan models.Client
	DeleteClient chan string
	NewEvent     chan models.Event
}

func (s *EventSocketService) Run() {
	cacheService = GetCacheService()

	for {
		select {
		case client := <-s.NewClient:
			s.clients[client.ID] = client.Conn
		case token := <-s.DeleteClient:
			delete(s.clients, token)
		case event := <-s.NewEvent:
			cacheService.EventToCache(event)
			s.notifyEvent()
		}
	}
}

func (s EventSocketService) notifyEvent() {
	for token, conn := range s.clients {
		err := conn.WriteMessage(websocket.BinaryMessage, []byte{})
		if err != nil {
			conn.Close()
			s.DeleteClient <- token
		}
	}
}

func GetEventSocketService() *EventSocketService {
	onceSocketService.Do(func() {
		eventSocketService = &EventSocketService{
			clients:      make(map[string]*websocket.Conn),
			NewClient:    make(chan models.Client),
			DeleteClient: make(chan string),
			NewEvent:     make(chan models.Event),
		}
	})

	return eventSocketService
}

package repositories

import "github.com/erikyvanov/pruebas-demos/models"

type DataEventRepository interface {
	GetEntityName() string
	GetEventName() string
	AddEvent(event models.Event)
	GetAllEvents() []models.Event
}

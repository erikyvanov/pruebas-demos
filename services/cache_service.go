package services

import (
	"sync"

	"github.com/erikyvanov/pruebas-demos/models"
	"github.com/erikyvanov/pruebas-demos/repositories"
)

type CacheService struct {
	repositories []repositories.DataEventRepository
}

func (e *CacheService) GetCache() []models.Event {
	events := make([]models.Event, 0)

	for _, repo := range e.repositories {
		events = append(events, repo.GetAllEvents()...)
	}

	return events
}

func (e *CacheService) EventToCache(event models.Event) {
	for _, repo := range e.repositories {
		if event.EventName == repo.GetEventName() {
			repo.AddEvent(event)
			return
		}
	}
}

var cacheService *CacheService
var onceCacheService sync.Once

func GetCacheService() *CacheService {
	onceCacheService.Do(func() {
		cacheService = &CacheService{
			repositories: []repositories.DataEventRepository{repositories.GetInvitationRepository()},
		}
	})

	return cacheService
}

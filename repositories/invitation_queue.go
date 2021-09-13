package repositories

import (
	"sync"

	"github.com/erikyvanov/pruebas-demos/models"
)

type InvitationRepository struct {
	entityName  string
	eventName   string
	invitations []models.Event
	lock        sync.RWMutex
}

func (q *InvitationRepository) AddEvent(event models.Event) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.invitations = append(q.invitations, event)
}

func (q *InvitationRepository) DequeueEvent() *models.Event {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.invitations) == 0 {
		return nil
	}

	event := q.invitations[0]
	q.invitations = q.invitations[1:]

	return &event
}

func (q *InvitationRepository) GetAllEvents() []models.Event {
	q.lock.Lock()
	defer q.lock.Unlock()

	invitations := q.invitations
	q.invitations = make([]models.Event, 0)

	return invitations
}

func (q *InvitationRepository) GetEntityName() string {
	return q.entityName
}

func (q *InvitationRepository) GetEventName() string {
	return q.eventName
}

var onceInvitationQueue sync.Once
var invitationRepository *InvitationRepository

func GetInvitationRepository() *InvitationRepository {
	onceInvitationQueue.Do(func() {
		invitationRepository = &InvitationRepository{
			entityName:  "user_space",
			eventName:   "invitation",
			invitations: make([]models.Event, 0),
		}
	})

	return invitationRepository
}

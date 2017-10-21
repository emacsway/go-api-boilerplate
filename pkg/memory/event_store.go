package memory

import (
	"app/pkg/domain"
	"sync"

	"github.com/google/uuid"
)

type eventStore struct {
	mtx    sync.RWMutex
	events map[string]*domain.Event
}

func (s *eventStore) Store(events []*domain.Event) error {
	if len(events) == 0 {
		return nil
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// todo check event version
	for _, e := range events {
		s.events[e.Id.String()] = e
	}

	return nil
}

func (s *eventStore) Get(id uuid.UUID) (*domain.Event, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if val, ok := s.events[id.String()]; ok {
		return val, nil
	}
	return nil, ErrEventNotFound
}

func (s *eventStore) FindAll() []*domain.Event {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	es := make([]*domain.Event, 0, len(s.events))
	for _, val := range s.events {
		es = append(es, val)
	}
	return es
}

func (s *eventStore) GetStream(streamId uuid.UUID, streamName string) []*domain.Event {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	e := make([]*domain.Event, 0, 0)
	for _, val := range s.events {
		if val.Metadata.StreamName == streamName && val.Metadata.StreamId == streamId {
			e = append(e, val)
		}
	}
	return e
}

func NewEventStore() domain.EventStore {
	return &eventStore{
		events: make(map[string]*domain.Event),
	}
}
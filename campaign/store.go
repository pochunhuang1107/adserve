package campaign

import (
	"errors"
	"sync"
)

type CampaignID uint64

type Store struct {
	mu        sync.Mutex
	campaigns map[CampaignID]Campaign
	nextID    CampaignID
}

func NewStore() *Store {
	return &Store{
		campaigns: make(map[CampaignID]Campaign),
		nextID:    1,
	}
}

func (s *Store) Create(c *Campaign) CampaignID {
	s.mu.Lock()
	defer s.mu.Unlock()

	nextID := s.nextID
	s.campaigns[nextID] = *c
	s.nextID = s.nextID + 1
	return nextID
}

func (s *Store) Get(id CampaignID) (Campaign, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.campaigns[id]
	return c, ok
}

func (s *Store) Deactivate(id CampaignID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	c, ok := s.campaigns[id]
	if !ok {
		return ErrCampaignNotFound
	}

	c.Deactivate()

	s.campaigns[id] = c

	return nil
}

var ErrCampaignNotFound = errors.New("campaign not found")

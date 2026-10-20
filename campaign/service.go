package campaign

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateCampaign(name string, budgetCents int) (CampaignID, error) {
	c, err := NewCampaign(name, budgetCents)
	if err != nil {
		return 0, err
	}
	id := s.store.Create(c)
	return id, nil
}

func (s *Service) DeactivateCampaign(id CampaignID) error {
	return s.store.Deactivate(id)
}

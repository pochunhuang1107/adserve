package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	type testCase struct {
		name         string
		campaignName string
		budgetCents  int
		expectedErr  error
	}

	tests := []testCase{
		{
			name:         "valid campaign",
			campaignName: "fall-sale",
			budgetCents:  5000,
			expectedErr:  nil,
		},
		{
			name:         "invalid name",
			campaignName: "",
			budgetCents:  5000,
			expectedErr:  ErrInvalidName,
		},
		{
			name:         "invalid budget",
			campaignName: "fall-sale",
			budgetCents:  0,
			expectedErr:  ErrInvalidBudget,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewCampaign(tt.campaignName, tt.budgetCents)

			if err != tt.expectedErr {
				t.Fatalf("expected error %v, got error: %v", tt.expectedErr, err)
			}

			if tt.expectedErr == nil {
				if c == nil {
					t.Fatal("expected campaign, got nil")
				}

				if tt.campaignName != c.Name() {
					t.Fatalf("expected name=%s, but got %s instead", tt.campaignName, c.Name())
				}

				if tt.budgetCents != c.BudgetCents() {
					t.Fatalf("expected budget=%d, but got %d instead", tt.budgetCents, c.BudgetCents())
				}

				if !c.IsActive() {
					t.Fatal("expected new campaign to be active")
				}
			}
		})
	}
}

func TestCampaignDeactivate(t *testing.T) {
	c, err := NewCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create campaign: %v", err)
	}
	if !c.IsActive() {
		t.Fatal("expected new campaign to be active")
	}
	c.Deactivate()
	if c.IsActive() {
		t.Fatal("expected campaign to be inactive")
	}
}

func TestCampaignCanServe(t *testing.T) {
	c, err := NewCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create campaign: %v", err)
	}
	if !c.CanServe() {
		t.Fatal("expected new campaign can serve")
	}
	c.Deactivate()
	if c.CanServe() {
		t.Fatal("expected new campaign can't serve")
	}
}

func TestNewStore(t *testing.T) {
	s := NewStore()
	c, err := NewCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create campaign: %v", err)
	}

	id := s.Create(c)
	campaign, ok := s.Get(id)
	if !ok {
		t.Fatal("id cannot be found in the store")
	}

	if campaign.Name() != c.Name() {
		t.Fatalf(
			"expected name=%s, got %s",
			c.Name(),
			campaign.Name(),
		)
	}

	if campaign.BudgetCents() != c.BudgetCents() {
		t.Fatalf(
			"expected budget=%d, got %d",
			c.BudgetCents(),
			campaign.BudgetCents(),
		)
	}
}

func TestStoreGet_NotFound(t *testing.T) {
	s := NewStore()

	_, ok := s.Get(999)
	if ok {
		t.Fatal("expected campaign not to be found")
	}
}

func TestServiceValidInput(t *testing.T) {
	store := NewStore()
	service := NewService(store)
	id, err := service.CreateCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if id == 0 {
		t.Fatal("expected a nonzero campaign ID")
	}
	campaign, ok := store.Get(id)
	if !ok {
		t.Fatalf("couldn't retrieve campaign with id=%d", id)
	}
	if campaign.Name() != "fall-sale" {
		t.Fatalf("expected campaign name=fall-sale, but got name=%s", campaign.Name())
	}
	if campaign.BudgetCents() != 5000 {
		t.Fatalf("expected budget=5000, but got budget=%d", campaign.BudgetCents())
	}
}

func TestServiceInvalidInput(t *testing.T) {
	store := NewStore()
	service := NewService(store)
	id, err := service.CreateCampaign("", 0)
	if err == nil {
		t.Fatal("expected error")
	}
	if id != 0 {
		t.Fatal("expected a zero campaign ID")
	}
	if err != ErrInvalidName {
		t.Fatalf("expected ErrInvalidName, got %v", err)
	}
	_, ok := store.Get(1)
	if ok {
		t.Fatal("invalid campaign should not be stored")
	}
}

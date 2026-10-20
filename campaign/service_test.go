package campaign

import "testing"

func TestServiceDeactivateExistingCampaignTwice(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	id, err := service.CreateCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	err = service.DeactivateCampaign(id)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	c, ok := store.Get(id)
	if !ok {
		t.Fatalf("campaign %d not found", id)
	}

	if c.IsActive() {
		t.Fatal("expected campaign to be inactive")
	}

	err = service.DeactivateCampaign(id)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	c, ok = store.Get(id)
	if !ok {
		t.Fatalf("campaign %d not found", id)
	}

	if c.IsActive() {
		t.Fatal("expected campaign to be inactive")
	}
}

func TestServiceDeactivateNonExistingCampaign(t *testing.T) {
	store := NewStore()
	service := NewService(store)

	err := service.DeactivateCampaign(1)
	if err != ErrCampaignNotFound {
		t.Fatalf("expected ErrCampaignNotFound, got %v", err)
	}
}

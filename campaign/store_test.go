package campaign

import (
	"sync"
	"testing"
)

func TestStoreConcurrentCreate(t *testing.T) {
	store := NewStore()

	c1, err := NewCampaign("campaign1", 1000)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	c2, err := NewCampaign("campaign2", 2000)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	var id1, id2 CampaignID
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		id1 = store.Create(c1)
	}()

	go func() {
		defer wg.Done()
		id2 = store.Create(c2)
	}()

	wg.Wait()
	if id1 == id2 {
		t.Fatalf("expected unique campaign IDs, both received %d", id1)
	}
	if id1 == 0 || id2 == 0 {
		t.Fatalf("expected nonzero campaign IDs, got %d and %d", id1, id2)
	}

	c, ok := store.Get(id1)
	if !ok {
		t.Fatalf("error retrieving campaign with id=%v\n", id1)
	}
	if c.Name() != c1.Name() {
		t.Fatalf("expected name=%s, got %s\n", c1.Name(), c.Name())
	}
	if c.BudgetCents() != c1.BudgetCents() {
		t.Fatalf("expected budget=%d, got %d\n", c1.BudgetCents(), c.BudgetCents())
	}

	c, ok = store.Get(id2)
	if !ok {
		t.Fatalf("error retrieving campaign with id=%v\n", id2)
	}
	if c.Name() != c2.Name() {
		t.Fatalf("expected name=%s, got %s\n", c2.Name(), c.Name())
	}
	if c.BudgetCents() != c2.BudgetCents() {
		t.Fatalf("expected budget=%d, got %d\n", c2.BudgetCents(), c.BudgetCents())
	}

}

func TestStoreCopiesCampaignOnCreateAndGet(t *testing.T) {
	store := NewStore()

	original, err := NewCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create campaign: %v", err)
	}
	id := store.Create(original)

	retrieved, ok := store.Get(id)
	if !ok {
		t.Fatalf("campaign not found: id=%d", id)
	}
	retrieved.Deactivate()

	again, ok := store.Get(id)

	if !ok {
		t.Fatalf("campaign not found: id=%d", id)
	}

	if again.IsActive() != true {
		t.Fatal("campaign should stay active")
	}

	original.Deactivate()
	again, ok = store.Get(id)

	if !ok {
		t.Fatalf("campaign not found: id=%d", id)
	}
	if again.IsActive() != true {
		t.Fatal("campaign should stay active")
	}
}

func TestStoreDeactivateExistingCampaign(t *testing.T) {
	store := NewStore()
	c, err := NewCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create campaign: %v", err)
	}

	id := store.Create(c)

	err = store.Deactivate(id)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	retrieved, ok := store.Get(id)
	if !ok {
		t.Fatalf("error retrieving campaign with id=%v\n", id)
	}
	if retrieved.IsActive() {
		t.Fatal("expected active=false, but got true")
	}
	if c.Name() != retrieved.Name() {
		t.Fatalf("expected name=%s, got %s\n", c.Name(), retrieved.Name())
	}
	if c.BudgetCents() != retrieved.BudgetCents() {
		t.Fatalf("expected budget=%d, got %d\n", c.BudgetCents(), retrieved.BudgetCents())
	}
}

func TestStoreDeactivateNonExistingCampaign(t *testing.T) {
	store := NewStore()

	err := store.Deactivate(1)
	if err != ErrCampaignNotFound {
		t.Fatalf("expected ErrCampaignNotFound, got %v\n", err)
	}
}

func TestStoreDeactivateTwice(t *testing.T) {
	s := NewStore()

	c, err := NewCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create campaign: %v", err)
	}
	id := s.Create(c)

	err = s.Deactivate(id)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	err = s.Deactivate(id)
	if err != nil {
		t.Fatalf("error: %v\n", err)
	}

	retrieved, ok := s.Get(id)
	if !ok {
		t.Fatalf("error retrieving campaign with id=%v\n", id)
	}
	if retrieved.IsActive() {
		t.Fatal("expected active=false, but got true")
	}
}

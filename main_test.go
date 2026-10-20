package main

import (
	"fmt"
	"github.com/pochunhuang1107/adserve/campaign"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func createTestCampaign(t *testing.T, service *campaign.Service) campaign.CampaignID {
	t.Helper()

	id, err := service.CreateCampaign("fall-sale", 5000)
	if err != nil {
		t.Fatalf("failed to create test campaign: %v", err)
	}

	return id
}

func TestHealthEndpoint(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)

	handler := newHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	if recorder.Body.String() != "ok\n" {
		t.Fatalf("expected body %q, got %q", "ok\n", recorder.Body.String())
	}
}

func TestCreateCampaignEndpoint(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)

	handler := newHandler(service)

	body := strings.NewReader(`{"name":"fall-sale","budget_cents":5000}`)
	req := httptest.NewRequest(http.MethodPost, "/campaigns", body)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, recorder.Code)
	}

	c, ok := store.Get(1)
	if !ok {
		t.Fatal("campaign creation failed")
	}

	if c.Name() != "fall-sale" {
		t.Fatalf("campaign name should be %s, got %s instead", "fall-sale", c.Name())
	}

	if c.BudgetCents() != 5000 {
		t.Fatalf("campaign budget should be %d, got %d instead", 5000, c.BudgetCents())
	}

	if recorder.Body.String() != "created campaign id=1\n" {
		t.Fatalf("expected body %q, got %q", "created campaign id=1\n", recorder.Body.String())
	}
}

func TestCreateCampaignEndpoint_InvalidName(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)

	handler := newHandler(service)

	body := strings.NewReader(`{"name":"","budget_cents":5000}`)
	req := httptest.NewRequest(http.MethodPost, "/campaigns", body)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}

	_, ok := store.Get(1)
	if ok {
		t.Fatal("campaign shouldn't be created due to invalid input")
	}
}

func TestDeactivateCampaignEndpoint_ExistingActive(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)
	handler := newHandler(service)

	id := createTestCampaign(t, service)

	req := httptest.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/campaigns/%d/deactivate", id),
		nil,
	)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, recorder.Code)
	}

	c, ok := store.Get(id)
	if !ok {
		t.Fatalf("error retriving campaign with campaign id=%d", id)
	}
	if c.IsActive() {
		t.Fatal("expected campaign to be inactive")
	}
	if recorder.Body.Len() != 0 {
		t.Fatalf("expected empty body for 204 response, got %q", recorder.Body.String())
	}
}

func TestDeactivateCampaignEndpoint_ExistingTwice(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)
	handler := newHandler(service)

	id := createTestCampaign(t, service)

	for i := 0; i < 2; i++ {
		req := httptest.NewRequest(
			http.MethodPost,
			fmt.Sprintf("/campaigns/%d/deactivate", id),
			nil,
		)

		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusNoContent {
			t.Fatalf("attempt %d: expected 204, got %d", i+1, recorder.Code)
		}

		if recorder.Body.Len() != 0 {
			t.Fatalf("attempt %d: expected empty response body", i+1)
		}

	}

	c, ok := store.Get(id)
	if !ok {
		t.Fatalf("campaign %d not found", id)
	}

	if c.IsActive() {
		t.Fatal("expected campaign to remain inactive")
	}

}

func TestDeactivateCampaignEndpoint_NonExistent(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)
	handler := newHandler(service)

	req := httptest.NewRequest(
		http.MethodPost,
		"/campaigns/999/deactivate",
		nil,
	)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestDeactivateCampaignEndpoint_Malformed(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)
	handler := newHandler(service)

	req := httptest.NewRequest(
		http.MethodPost,
		"/campaigns/abc/deactivate",
		nil,
	)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestDeactivateCampaignEndpoint_GetRequest(t *testing.T) {
	store := campaign.NewStore()
	service := campaign.NewService(store)
	handler := newHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/campaigns/1/deactivate",
		nil,
	)

	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d, got %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/pochunhuang1107/adserve/campaign"
	"log"
	"net/http"
	"strconv"
)

func newHandler(service *campaign.Service) http.Handler {
	mux := http.NewServeMux()

	type createCampaignRequest struct {
		Name        string `json:"name"`
		BudgetCents int    `json:"budget_cents"`
	}

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		fmt.Fprintln(w, "ok")
	})

	mux.HandleFunc("/campaigns", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req createCampaignRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}

		id, err := service.CreateCampaign(req.Name, req.BudgetCents)
		if err != nil {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "created campaign id=%d\n", id)
	})

	mux.HandleFunc("POST /campaigns/{id}/deactivate", func(w http.ResponseWriter, r *http.Request) {
		rawID := r.PathValue("id")

		parsedID, err := strconv.ParseUint(rawID, 10, 64)
		if err != nil {
			http.Error(w, "invalid id type", http.StatusBadRequest)
			return
		}
		id := campaign.CampaignID(parsedID)

		err = service.DeactivateCampaign(id)
		switch err {
		case nil:
			w.WriteHeader(http.StatusNoContent)
		case campaign.ErrCampaignNotFound:
			http.Error(w, "campaign not found", http.StatusNotFound)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}

	})

	return mux
}

func main() {

	store := campaign.NewStore()
	service := campaign.NewService(store)

	handler := newHandler(service)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

package llmclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const testDaySessionID = "550e8400-e29b-41d4-a716-446655440000"

func TestClientPlanContract(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Fatalf("method = %s, want POST", r.Method)
		}
		if r.URL.Path != "/plan" {
			t.Fatalf("path = %s, want /plan", r.URL.Path)
		}
		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Fatalf("content type = %s, want application/json", got)
		}

		var request PlanRequest
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&request); err != nil {
			t.Fatalf("decode request: %v", err)
		}
		if request.DaySessionID != testDaySessionID {
			t.Fatalf("day_session_id = %s, want %s", request.DaySessionID, testDaySessionID)
		}

		arrival, err := time.Parse(
			time.RFC3339,
			"2026-08-13T09:00:00+05:30",
		)
		if err != nil {
			t.Fatal(err)
		}

		departure, err := time.Parse(
			time.RFC3339,
			"2026-08-13T10:30:00+05:30",
		)
		if err != nil {
			t.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(PlanResponse{
			Stops: []Stop{{
				Position:         1,
				Title:            "Bangalore Palace",
				CategoryLabel:    "Sightseeing",
				ImageURL:         "",
				PlannedArrival:   arrival,
				PlannedDeparture: departure,
				TravelMinutes:    20,
				StayMinutes:      90,
			}},
		})
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}

	got, err := client.Plan(context.Background(), PlanRequest{DaySessionID: testDaySessionID})
	if err != nil {
		t.Fatal(err)
	}
	if len(got.Stops) != 1 || got.Stops[0].Title != "Bangalore Palace" {
		t.Fatalf("unexpected response: %+v", got)
	}
}

func TestClientPlanNon2xx(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Plan(context.Background(), PlanRequest{DaySessionID: testDaySessionID})
	if err == nil || !strings.Contains(err.Error(), "status 400") {
		t.Fatalf("expected status error, got %v", err)
	}
}

func TestClientPlanTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: 10 * time.Millisecond})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Plan(context.Background(), PlanRequest{DaySessionID: testDaySessionID})
	if err == nil || !strings.Contains(err.Error(), "timed out") {
		t.Fatalf("expected timeout error, got %v", err)
	}
}

func TestClientPlanRejectsInvalidRequest(t *testing.T) {
	client, err := NewClient(Config{BaseURL: "http://localhost:8000", Timeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Plan(context.Background(), PlanRequest{DaySessionID: "not-a-uuid"})
	if err == nil || !strings.Contains(err.Error(), "valid UUID") {
		t.Fatalf("expected UUID validation error, got %v", err)
	}
}

func TestClientPlanRejectsUnknownResponseFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"stops":[{"position":1,"title":"Bangalore Palace","category_label":"Sightseeing","image_url":"","planned_arrival":"09:00","planned_departure":"10:30","travel_minutes":20,"stay_minutes":90,"unexpected":"field"}]}`))
	}))
	defer server.Close()

	client, err := NewClient(Config{BaseURL: server.URL, Timeout: time.Second})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Plan(context.Background(), PlanRequest{DaySessionID: testDaySessionID})
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("expected unknown field error, got %v", err)
	}
}

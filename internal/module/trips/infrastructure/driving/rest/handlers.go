package rest

import (
	"encoding/json"
	"net/http"
	"trip/application/command"
	_ "trip/application/dto"
	"trip/application/services"
)

type Handler struct {
	createTrip *services.CreateTripService
	listTrips  *services.ListTripService
}

func NewHandler(createTrip *services.CreateTripService, listTrips *services.ListTripService) *Handler {
	return &Handler{
		createTrip: createTrip,
		listTrips:  listTrips,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/trips" {
		writeError(w, r, http.StatusNotFound, "not_found", "not found")
		return
	}

	switch r.Method {
	case http.MethodPost:
		h.create(w, r)
	case http.MethodGet:
		h.list(w, r)
	default:
		writeError(w, r, http.StatusMethodNotAllowed, "bad_request", "method not allowed")
	}
}

type CreateTripRequest struct {
	// Destination city
	Destination string `json:"destination"`

	// Start Date
	StartDate string `json:"start_date"`

	// End Date
	EndDate string `json:"end_date"`

	// Number of Travelers
	TravelersCount int `json:"travelers_count"`
}

// CreateTrips godoc
// @Summary Create a Trip
// @Description Creates a new trip for the specified destination and dates.
// @Tags Trip
// @Accept json
// @Produce json
// @Param request body CreateTripRequest true "Trip details"
// @Success 201 {object} dto.TripDTO
// @Failure 400 {object} apiError
// @Failure 404 {object} apiError
// @Failure 500 {object} apiError
// @Router /api/trips [post]
func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateTripRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, r, http.StatusBadRequest, "bad_request", "invalid json body")
		return
	}

	trip, err := h.createTrip.CreateTrip(r.Context(), command.CreateTripCommand{
		Destination:    req.Destination,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		TravelersCount: req.TravelersCount,
	})
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(trip)
}

// ListTrips godoc
// @Summary List all the Trips
// @Description List all the trips created with trip id.
// @Tags Trip
// @Produce json
// @Success 200 {array} dto.Tripdto
// @Failure 500 {object} ErrorResponse
// @Router /api/trips [get]
func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	trips, err := h.listTrips.ListTrips(r.Context())
	if err != nil {
		writeDomainError(w, r, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(trips)
}

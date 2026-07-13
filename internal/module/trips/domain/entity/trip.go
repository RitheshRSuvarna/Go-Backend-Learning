package entity

import (
	"common"
	"time"
)

type Trip struct {
	id          common.TripID
	destination string
	startDate   string
	endDate     string
	travelCount int
	createdAt   common.Time
}

func NewTrip(destination, startDate, endDate string, travelCount int) (*Trip, error) {

	if destination == "" {
		return nil, common.NewValidationError("destination is required", nil)
	}

	if startDate == "" {
		return nil, common.NewValidationError("start_date is required", nil)
	}

	if endDate == "" {
		return nil, common.NewValidationError("end_date is required", nil)
	}

	if travelCount <= 0 {
		return nil, common.NewValidationError("traveler_count must be greater then zero", nil)
	}

	if _, err := time.Parse("2006-01-02", startDate); err != nil {
		return nil, common.NewValidationError("start_date must be YYYY-MM-DD", err)
	}

	if _, err := time.Parse("2006-01-02", endDate); err != nil {
		return nil, common.NewValidationError("end_date must be YYYY-MM-DD", err)
	}

	now := common.Now()

	return &Trip{
		destination: destination,
		startDate:   startDate,
		endDate:     endDate,
		travelCount: travelCount,
		createdAt:   now,
	}, nil

}

func (t *Trip) ID() common.TripID      { return t.id }
func (t *Trip) Destination() string    { return t.destination }
func (t *Trip) StartDate() string      { return t.startDate }
func (t *Trip) EndDate() string        { return t.endDate }
func (t *Trip) TravelersCount() int    { return t.travelCount }
func (t *Trip) CreatedAt() common.Time { return t.createdAt }

func (t *Trip) SetID(id common.TripID) {
	t.id = id
}

func (t *Trip) AssignPersistence(id common.TripID, createdAt common.Time) {
	t.id = id
	t.createdAt = createdAt
}

func RestoreTrip(
	id common.TripID,
	destination, startDate, endDate string,
	travelCount int,
	createdAt common.Time,
) *Trip {
	return &Trip{
		id:          id,
		destination: destination,
		startDate:   startDate,
		endDate:     endDate,
		travelCount: travelCount,
		createdAt:   createdAt,
	}
}

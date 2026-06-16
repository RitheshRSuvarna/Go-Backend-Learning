package entity

import (
	"common"
	"time"
)

type DaySession struct {
	id         common.DaySessionID
	tripID     common.TripID
	date       string
	starttime  string
	startlabel string
	createdAt  common.Time
}

func NewDaySession(TripID, date, starttime, startlabel string) (*DaySession, error) {
	if TripID == "" {
		return nil, common.NewValidationError("Trip is is required", nil)
	}
	if date == "" {
		return nil, common.NewValidationError("date is required", nil)
	}
	if starttime == "" {
		return nil, common.NewValidationError("time is required", nil)
	}
	if startlabel == "" {
		return nil, common.NewValidationError("start_balel is required", nil)
	}
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return nil, common.NewValidationError("start_date must be in YYYY-MM-DD format", err)
	}

	now := common.Now()

	return &DaySession{
		tripID:     common.TripID{},
		date:       date,
		starttime:  starttime,
		startlabel: startlabel,
		createdAt:  now,
	}, nil
}

func (d *DaySession) ID() common.DaySessionID { return d.id }
func (d *DaySession) TripID() common.TripID   { return d.tripID }
func (d *DaySession) Date() string            { return d.date }
func (d *DaySession) STime() string           { return d.starttime }
func (d *DaySession) Label() string           { return d.startlabel }
func (d *DaySession) CreatedAt() common.Time  { return d.createdAt }

func (d *DaySession) SetID(id common.DaySessionID) {
	d.id = id
}

func (d *DaySession) AssignPresistance(id common.DaySessionID, createdAt common.Time) {
	d.id = id
	d.createdAt = createdAt
}

func RestoreDaySession(
	id common.DaySessionID,
	tripID common.TripID,
	date, starttime, startlabel string,
	createdAt common.Time,
) *DaySession {
	return &DaySession{
		id:         id,
		tripID:     tripID,
		date:       date,
		starttime:  starttime,
		startlabel: startlabel,
		createdAt:  createdAt,
	}
}

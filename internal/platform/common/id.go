package common

import (
	"fmt"

	"github.com/google/uuid"
)

type ID struct {
	value string
}

func NewID(value string) (ID, error) {
	if value == "" {
		return ID{}, fmt.Errorf("value cannot be empty")
	}
	return ID{value: value}, nil
}

func GenerateID() ID {
	return ID{value: uuid.New().String()}
}

func (id ID) String() string {
	return id.value
}

func (id ID) Value() string {
	return id.value
}

func (id ID) IsZero() bool {
	return id.value == ""
}

type TripID ID

func NewTripID(value string) (TripID, error) {
	id, err := NewID(value)
	if err != nil {
		return TripID{}, err
	}
	return TripID(id), nil
}

func GenerateTripID() TripID {
	return TripID(GenerateID())
}

type DaySessionID ID

func NewDaySessionID(value string) (DaySessionID, error) {
	id, err := NewID(value)
	if err != nil {
		return DaySessionID{}, err
	}
	return DaySessionID(id), nil
}

func GenerateDaySessionID() DaySessionID {
	return DaySessionID(GenerateID())
}

type PlanVersionID ID

func NewPlanVersionID(value string) (PlanVersionID, error) {
	id, err := NewID(value)
	if err != nil {
		return PlanVersionID{}, err
	}
	return PlanVersionID(id), nil
}

func GeneratePlanVersionID() PlanVersionID {
	return PlanVersionID(GenerateID())
}

type PlanStopID ID

func NewPlanStopID(value string) (PlanStopID, error) {
	id, err := NewID(value)
	if err != nil {
		return PlanStopID{}, err
	}
	return PlanStopID(id), nil
}

func GeneratePlanStopID() PlanStopID {
	return PlanStopID(GenerateID())
}

func (id TripID) String() string { return ID(id).String() }
func (id TripID) Value() string  { return ID(id).Value() }
func (id TripID) IsZero() bool   { return ID(id).IsZero() }

func (id DaySessionID) String() string { return ID(id).String() }
func (id DaySessionID) Value() string  { return ID(id).Value() }
func (id DaySessionID) IsZero() bool   { return ID(id).IsZero() }

func (id PlanVersionID) String() string { return ID(id).String() }
func (id PlanVersionID) Value() string  { return ID(id).Value() }
func (id PlanVersionID) IsZero() bool   { return ID(id).IsZero() }

func (id PlanStopID) String() string { return ID(id).String() }
func (id PlanStopID) Value() string  { return ID(id).Value() }
func (id PlanStopID) IsZero() bool   { return ID(id).IsZero() }

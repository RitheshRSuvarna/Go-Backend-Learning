package common

import (
	"fmt"
	"github.com/google/uuid"
)

type ID struct{
	value string
}

func NewID(value string) (ID, error) {
	if value == "" {
		return ID{}, fmt.Errorf("value cannot be empty")
	}
	return ID{value: value}, nil
	}

func GenerateID() ID{
	return ID{value:uuid.New().String()}
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

func NewTripID(value string) (TripID, error){
	id, err :=NewID(value)
	if err != nil {
		return TripID{}, err
	}
	return TripID(id), nil
}

func GenerateTripID() TripID {
	return TripID(GenerateID())
}

func (id TripID) String() string{ return ID(id).String()}
func (id TripID) Value() string{ return ID(id).Value()}
func (id TripID) IsZero() bool{ return ID(id).IsZero()}

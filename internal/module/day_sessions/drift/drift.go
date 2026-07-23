package drift

import (
	"time"

	"common"
)

type Status struct {
	DriftMinutes int    `json:"drift_minutes"`
	DriftStatus  string `json:"drift_status"`
	Replan       bool   `json:"replan_recommendation"`
}

func Compute(plannedArrival time.Time, actualArrival common.Time) Status {
	diff := actualArrival.Time().Sub(plannedArrival)

	// Convert duration to minutes
	driftMinutes := int(diff.Minutes())

	status := Status{
		DriftMinutes: driftMinutes,
	}

	switch {
	case driftMinutes > 20:
		status.DriftStatus = "Delayed"
		status.Replan = true

	case driftMinutes < -20:
		status.DriftStatus = "Early"
		status.Replan = false

	default:
		status.DriftStatus = "OnTime"
		status.Replan = false
	}

	return status
}

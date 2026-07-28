package entity

import "encoding/json"

type Payload struct {
	PlanStopID string `json:"plan_stop_id"`
}

func (e *Events) PlanStopID() (string, error) {
    var payload Payload

    if err := json.Unmarshal(e.Payload(), &payload); err != nil {
        return "", err
    }

    return payload.PlanStopID, nil
}
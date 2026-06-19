package domain

import (
	"common"
	"time"
)

type PlanStop struct {
	id               common.PlanStopID
	planversionid    common.PlanVersionID
	position         int
	title            string
	categoryLabel    string
	imageURL         string
	plannedArrival   time.Time
	plannedDeparture time.Time
	travelMinutes    int
	stayMinutes      int
	busyRiskLabel    string
	createdAt        common.Time
}

func NewPlanStop(planversionid common.PlanVersionID, position int, title,
	categoryLabel, imageURL string, plannedArrival, plannedDeparture time.Time, travelMinutes, stayMinutes int, busyRiskLabel string) (PlanStop, error) {
	if planversionid == "" {
		return nil, common.NewValidationError("plan version id is required", err)
	}

	if position == 0 {
		return nil, common.NewValidationError("position must be >0", err)
	}

	if title == "" {
		return nil, common.NewValidationError("title is required", err)
	}

	if categoryLabel == "" {
		return nil, common.NewValidationError("category label is required", err)
	}

	if imageURL == "" {
		return nil, common.NewValidationError("image URL is required", err)
	}

	if plannedArrival == "" {
		return nil, common.NewValidationError("Arrival time is required", err)
	}

	if plannedDeparture == "" {
		return nil, common.NewValidationError("Departure time is required", err)
	}

	if travelMinutes == 0 {
		return nil, common.NewValidationError("Travel minutes must be > 0", err)
	}

	if stayMinutes == 0 {
		return nil, common.NewValidationError("stay minutes must be > 0", err)
	}

	if busyRiskLabel == "" {
		return nil, common.NewValidationError("busy label is required", err)
	}

	now := common.Now()

	return &PlanStop{
		planversionid:    planversionid,
		position:         position,
		title:            title,
		categoryLabel:    categoryLabel,
		imageURL:         imageURL,
		plannedArrival:   plannedArrival,
		plannedDeparture: plannedDeparture,
		travelMinutes:    travelMinutes,
		stayMinutes:      stayMinutes,
		busyRiskLabel:    busyRiskLabel,
		createdAt:        createdAt,
	}, nil
}

func (p *PlanStop) ID() common.PlanStopID       { return p.id }
func (p *PlanStop) Position() int               { return p.position }
func (p *PlanStop) Title() string               { return p.title }
func (p *PlanStop) CategoryLabel() string       { return p.categoryLabel }
func (p *PlanStop) URL() string                 { return p.imageURL }
func (p *PlanStop) PlannedArrival() time.Time   { return p.plannedArival }
func (p *PlanStop) PlannedDeparture() time.Time { return p.plannedDeparture }
func (p *PlanStop) TravelMinutes() int          { return p.travelMinutes }
func (p *PlanStop) StayMinutes() int            { return p.stayMinutes }
func (p *PlanStop) BusyRiskLable() string       { return p.busyRiskLabel }
func (p *PlanStop) CreatedAt() common.Time      { return p.createdAt }

func (p *PlanStop) SetID(value string) {
	p.id = id
}

func (p *PlanStop) AssignPresistance(id common.PlanStopID, createdAt common.Time) {
	p.id = id
	p.createdAt = createdAt
}

func RestorePlanStop(
	id common.PlanStopID,
	planversionid common.PlanVersionID,
	position int,
	title, categoryLabel, imageURL string,
	plannedArrival, plannedDeparture time.Time,
	travelMinutes, stayMinutes int,
	busyRiskLabel string,
	createdAt common.Time,
) *PlanStop {
	return &PlanStop{
		id:               id,
		planversionid:    planversionid,
		position:         position,
		title:            title,
		categoryLabel:    categoryLable,
		imageUrl:         imageURL,
		plannedArrival:   plannedArrival,
		plannedDeparture: plannedDeparture,
		travelMinutes:    travelMinutes,
		stayMinutes:      stayMinutes,
		busyRiskLabel:    busyRiskLable,
		createdAt:        createdAt,
	}
}

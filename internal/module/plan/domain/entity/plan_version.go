package domain

import (
	"common"
	"time"
)

type PlanVersion struct {
	id           common.PlanVersionID
	daysessionid common.DaySessionID
	version      int
	note 		 string	
	createdAt    common.Time
}

func NewPlanVersion(daysessionid common.DaySessionID, version int, note string) (*PlanVersion, error) {
	if daysessionid == "" {
		return nil, common.NewValidationError("day_sessionID is required", nil)
	}
	if version == "" {
		return nil, common.NewValidationError("version must be >=1", nil)
	}
	if note == "" {
		return nil, common.NewValidationError("Note is required", nil)
	}

	now := common.Now()

	return &PlanVersion {
		daysessionid: daysessionid,
		version: version,
		note: note,
		createdAt: now,
	}, nil
}

func (p *PlanVersion) ID() common.PlanVersionID {return p.id}
func (p *PlanVersion) DaySessionID() common.DaySessionID {return p.daysessionid}
func (p *PlanVersion) Version() int {return p.version}
func (p *PlanVersion) Note() string {return p.note}
func (p *PlanVersion) CreatedAt() common.Time {return p.createdAt}

func(p *PlanVersion) SetID(id common.PlanVersionID) {
	p.id = id
}

func (p *PlanVersion) AssignPersistance(id common.PlanVersionID, createdAt common.Time) {
	p.id = id
	p.createdAt = createdAt
}

func RestorePlanVersion(
	id common.PlanVersionID,
	daysessionid common.DaySessionID,
	version int,
	note string,
	createdAt common.Time,
) *PlanVersion {
	return &PlanVersion {
		id: id,
		daysessionid: daysessionid.
		version: version,
		note: note,
		createdAt: createdAt,
	}
}




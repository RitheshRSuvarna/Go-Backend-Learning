package repository

import (
	"common"
	"fmt"
	"plan/domain/entity"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func rowToDomainPlanStop(
	id, planversionid pgtype.UUID,
	position int,
	title, categoryLabel,imageURL string,
	plannedArrival,plannedDeparture time.Time,
	travelMinutes,stayMinutes int,
	busyRiskLabel string,
	createdAt pgtype.Timestamptz,
) (*entity.PlanStop, error) {
	planstopid, err := common.NewPlanStopID(pgTypeToString(id))
	if err != nil {
		return nil, fmt.Errorf("Invalid plan stop id: %v", err)
	}

	planversionId, err := common.NewPlanVersionID(pgTypeToString(planversionid))
	if err != nil {
		return nil, fmt.Errorf("Invalid planversion id: %v", err)
	}

	if !createdAt.Valid {
		return nil, fmt.Errorf("Invalid createdAt")
	}

	return entity.RestorePlanStop(
		planstopid,
		planversionId,
		position,
		title,
		categoryLabel,
		imageURL,
		plannedArrival,
		plannedDeparture,
		travelMinutes,
		stayMinutes,
		busyRiskLabel,
		common.NewTime(createdAt.Time),
	), nil


}
package repository

import (
	"common"
	"fmt"
	"plan/domain/entity"

	"github.com/jackc/pgx/v5/pgtype"
)

func rowToDomainPlanVersion(
	id, daysessionID pgtype.UUID,
	version int,
	note string,
	createdAt pgtype.Timestamptz,
) (*entity.PlanVersion, error) {
	planversionID, err := common.NewPlanVersionID(pgTypeToString(id))
	if err != nil {
		return nil, fmt.Errorf("Invalid plan version id: %v", err)
	}
	daysessioniD, err := common.NewDaySessionID(pgTypeToString(daysessionID))
	if err != nil {
		return nil, fmt.Errorf("Invalid daysession id: %v", err)
	}

	if !createdAt.Valid {
		return nil, fmt.Errorf("Invalid created_At")
	}

	return entity.RestorePlanVersion(
		planversionID,
		daysessioniD,
		version,
		note,
		common.NewTime(createdAt.Time),
	), nil
}

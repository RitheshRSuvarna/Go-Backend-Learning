package repository

import (
	"common"
	"context"
	"fmt"
)

func (r *PostgresPlanVersionRepository) GetLatestVersion(ctx context.Context, id common.DaySessionID) (int, error) {
	pgid, err := uuidStringToPgUUID(id.String())
	if err != nil {
		return 0, err
	}

	version, err := r.getQueries(ctx).GetLatestPlanVersion(ctx, pgid)
	if err != nil {
		return 0, fmt.Errorf("Failed to get latest plan version: %w", err)
	}

	return int(version), nil
}

package services

import (
	"common"
	"context"
	"fmt"
	"plans/domain/repository"
)

type GetLatestVersionService struct {
	planVersionRepo repository.PlanVersionRepository
}

func NewGetLatestVersion(planVersionRepo repository.PlanVersionRepository) *GetLatestVersionService {
	return &GetLatestVersionService{
		planVersionRepo: planVersionRepo,
	}
}

func (s *GetLatestVersionService) GetLatestVersion(ctx context.Context, id common.DaySessionID) (int, error) {
	latestVersion, err := s.planVersionRepo.GetLatestVersion(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest plan version: %w", err)
	}
	return latestVersion, nil
}
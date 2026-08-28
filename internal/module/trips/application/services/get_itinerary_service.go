package services

// import (
// 	"fmt"
// 	"context"
// 	"common"
// 	"trip/application/dto"
// 	"trip/domain/repository"
// 	daysession "day_session/domain/repository"
// 	planrepo "plans/domain/repository"
// )

// type GetItineraryService struct {
// 	tripRepo        repository.TripRepository
// 	daySessionRepo  daysession.DaySessionRepository
// 	planVersionRepo planrepo.PlanVersionRepository
// 	planStopRepo    planrepo.PlanStopRepository
// }

// func NewGetItineraryService(
// 	tripRepo repository.TripRepository,
// 	daySessionRepo daysession.DaySessionRepository,
// 	planVersionRepo planrepo.PlanVersionRepository,
// 	planStopRepo planrepo.PlanStopRepository,
// ) *GetItineraryService {
// 	return &GetItineraryService{
// 		tripRepo:        tripRepo,
// 		daySessionRepo:  daySessionRepo,
// 		planVersionRepo: planVersionRepo,
// 		planStopRepo:    planStopRepo,
// 	}
// }

// func (s *GetItineraryService) GetItinerary(
// 	ctx context.Context,
// 	tripID common.TripID,
// ) (dto.ItineraryDTO, error) {

// 	// Get trip
// 	trip, err := s.tripRepo.GetByID(ctx, tripID)
// 	if err != nil {
// 		return dto.ItineraryDTO{}, err
// 	}

// 	// Get ALL day sessions for this trip
// 	daySessions, err := s.daySessionRepo.ListDaySession(ctx, tripID)
// 	if err != nil {
// 		return dto.ItineraryDTO{}, err
// 	}

// 	// Print all day session IDs
// 	for _, daySession := range daySessions {
// 		fmt.Println("Day Session ID:", daySession.ID().String())
// 	}

// 	// Prepare response
// 	out := dto.ItineraryDTO{
// 		Trip:   trip.ID().String(),
// 		TripName: trip.Name(),
// 	}

// 	// Process EVERY day session
// 	for _, daySession := range daySessions {

// 		// Get active plan for THIS day session
// 		planVersion, err := s.planVersionRepo.GetActivePlan(
// 			ctx,
// 			daySession.ID(),
// 		)
// 		if err != nil {
// 			return dto.ItineraryDTO{}, err
// 		}

// 		// Get plan stops for THIS plan version
// 		planStops, err := s.planStopRepo.ListStop(
// 			ctx,
// 			planVersion.ID(),
// 		)
// 		if err != nil {
// 			return dto.ItineraryDTO{}, err
// 		}

// 		// Print stops
// 		fmt.Println(
// 			"Day Session:",
// 			daySession.ID().String(),
// 			"Plan Version:",
// 			planVersion.ID().String(),
// 		)

// 		for _, stop := range planStops {
// 			fmt.Println(
// 				"Plan Stop:",
// 				stop.ID().String(),
// 				stop.Name(),
// 			)
// 		}

// 		// Add stops to response
// 		for _, stop := range planStops {
// 			out.PlanStops = append(out.PlanStops, dto.PlanStopDTO{
// 				ID:          stop.ID().String(),
// 				Name:        stop.Name(),
// 				Description: stop.Description(),
// 			})
// 		}
// 	}

// 	return out, nil
// }

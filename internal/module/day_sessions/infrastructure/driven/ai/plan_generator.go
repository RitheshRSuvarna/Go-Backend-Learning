package ai

import (
	"context"
	"fmt"
	"time"

	"day_session/domain/port"
	"llmclient"
)

type PlanGenerator struct {
	client *llmclient.Client
}

func NewPlanGenerator(client *llmclient.Client) *PlanGenerator {
	return &PlanGenerator{
		client: client,
	}
}

func (g *PlanGenerator) GeneratePlan(
	ctx context.Context,
	request port.PlanRequest,
) (port.PlanResponse, error) {

	response, err := g.client.Plan(ctx, llmclient.PlanRequest{
		DaySessionID: request.DaySessionID,
		Destination:  request.Destination,
		Date:         request.Date,
		StartTime:    request.StartTime,
		StartLabel:   request.StartLabel,
	})
	if err != nil {
		return port.PlanResponse{}, err
	}

	stops := make([]port.PlanStop, 0, len(response.Stops))

	for _, stop := range response.Stops {
		stops = append(stops, port.PlanStop{
			Position:         stop.Position,
			Title:             stop.Title,
			CategoryLabel:    stop.CategoryLabel,
			ImageURL:         stop.ImageURL,
			PlannedArrival:   stop.PlannedArrival,
			PlannedDeparture: stop.PlannedDeparture,
			TravelMinutes:    stop.TravelMinutes,
			StayMinutes:      stop.StayMinutes,
		})
	}

	return port.PlanResponse{
		Stops: stops,
	}, nil
}

func (g *PlanGenerator) Replan(
	ctx context.Context,
	request port.RePlanRequest,
) (port.PlanResponse, error) {

	existingStops := make([]llmclient.Stop, 0, len(request.ExistingPlanStops))

	for _, stop := range request.ExistingPlanStops {
		existingStops = append(existingStops, llmclient.Stop{
			Position:         stop.Position,
			Title:             stop.Title,
			CategoryLabel:    stop.CategoryLabel,
			ImageURL:         stop.ImageURL,
			PlannedArrival:   stop.PlannedArrival,
			PlannedDeparture: stop.PlannedDeparture,
			TravelMinutes:    stop.TravelMinutes,
			StayMinutes:      stop.StayMinutes,
		})
	}

	fmt.Printf("REPLAN: sending %d existing stops to LLM\n", len(existingStops))

	start := time.Now()

	response, err := g.client.Replan(
		ctx,
		llmclient.ReplanRequest{
			DaySessionID:      request.DaySessionID,
			Destination:       request.Destination,
			Date:              request.Date,
			StartTime:         request.StartTime,
			StartLabel:        request.StartLabel,
			ExistingPlanStops: existingStops,
		},
	)

	fmt.Printf("REPLAN: LLM call took %v\n", time.Since(start))

	if err != nil {
		return port.PlanResponse{}, err
	}

	stops := make([]port.PlanStop, 0, len(response.Stops))

	for _, stop := range response.Stops {
		stops = append(stops, port.PlanStop{
			Position:         stop.Position,
			Title:             stop.Title,
			CategoryLabel:    stop.CategoryLabel,
			ImageURL:         stop.ImageURL,
			PlannedArrival:   stop.PlannedArrival,
			PlannedDeparture: stop.PlannedDeparture,
			TravelMinutes:    stop.TravelMinutes,
			StayMinutes:      stop.StayMinutes,
		})
	}

	return port.PlanResponse{
		Stops: stops,
	}, nil
}
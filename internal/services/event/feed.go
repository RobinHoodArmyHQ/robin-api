package event

import (
	"context"
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
)

func (i *impl) GetEventFeed(ctx context.Context, req *GetEventFeedRequest) (*GetEventFeedResponse, error) {
	ev := env.FromContext(ctx)
	// Get events from repository
	resp, err := ev.EventRepository.GetEventsByLocation(&event.GetEventsByLocationRequest{
		CityId: req.CityId,
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		return nil, err
	}

	return &GetEventFeedResponse{Events: resp.Events}, nil
}

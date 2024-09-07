package event

import (
	"context"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type Service interface {
	GetEventFeed(ctx context.Context, req *GetEventFeedRequest) (*GetEventFeedResponse, error)
	MarkEventInterested(ctx context.Context, eventID nanoid.NanoID) error
	GetEventParticipants(ctx context.Context, eventID nanoid.NanoID) (*GetEventParticipantsResponse, error)
}

type impl struct {
}

func New() Service {
	return &impl{}
}

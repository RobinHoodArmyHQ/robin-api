package event

import (
	"context"
)

type Service interface {
	GetEventFeed(ctx context.Context, req *GetEventFeedRequest) (*GetEventFeedResponse, error)
}

type impl struct {
}

func New() Service {
	return &impl{}
}

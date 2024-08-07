package participants

import (
	"context"
	"github.com/RobinHoodArmyHQ/robin-api/models"
)

type ParticipantsRepository interface {
	GetEventParticipant(ctx context.Context, req *GetEventParticipantRequest) (*GetEventParticipantResponse, error)
	GetEventParticipants(ctx context.Context, eventID int64) (*GetEventParticipantsResponse, error)
	CreateParticipant(ctx context.Context, req *CreateParticipantsRequest) error
}

type GetEventParticipantRequest struct {
	EventID int64
	UserID  uint64
}

type GetEventParticipantResponse struct {
	Participant *models.Participant
}

type GetEventParticipantsResponse struct {
	Participants []*models.Participant
}

type CreateParticipantsRequest struct {
	Participant *models.Participant
}

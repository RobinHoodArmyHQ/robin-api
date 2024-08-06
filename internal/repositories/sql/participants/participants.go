package participants

import (
	"context"
	"errors"
	"fmt"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/participants"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ParticipantsRepository is the repository for event participants
type ParticipantsRepository struct {
	logger *zap.Logger
	db     *database.SqlDB
}

// NewParticipantsRepository creates a new ParticipantsRepository
func NewParticipantsRepository(logger *zap.Logger, db *database.SqlDB) *ParticipantsRepository {
	return &ParticipantsRepository{logger, db}
}

func (r *ParticipantsRepository) CreateParticipant(ctx context.Context, req *participants.CreateParticipantsRequest) error {
	exec := r.db.Master().Create(req.Participant)
	if exec.Error != nil {
		return fmt.Errorf("failed to create participant: %v", exec.Error)
	}

	return nil
}

func (r *ParticipantsRepository) GetEventParticipant(ctx context.Context, req *participants.GetEventParticipantRequest) (*participants.GetEventParticipantResponse, error) {
	p := &models.Participant{}
	resp := &participants.GetEventParticipantResponse{}
	exec := r.db.Master().First(p, "event_id = ? AND user_id = ?", req.EventID, req.UserID)
	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		err := fmt.Errorf("event not found: %v", req.EventID)
		r.logger.Error(err.Error())
		return resp, nil
	}

	if exec.Error != nil {
		return nil, fmt.Errorf("failed to get events: %v", exec.Error)
	}

	resp.Participant = p
	return resp, nil
}

func (r *ParticipantsRepository) GetEventParticipants(ctx context.Context, eventID int64) (*participants.GetEventParticipantsResponse, error) {
	resp := &participants.GetEventParticipantsResponse{}
	var participants []*models.Participant
	exec := r.db.Master().
		Preload("User").
		Where("event_id = ?", eventID).
		Find(&participants)

	if errors.Is(exec.Error, gorm.ErrRecordNotFound) {
		err := fmt.Errorf("no participants found for event: %v", eventID)
		r.logger.Error(err.Error())
		return resp, nil
	}

	if exec.Error != nil {
		r.logger.Error(exec.Error.Error())
		return nil, fmt.Errorf("failed to get participants: %v", exec.Error)
	}

	resp.Participants = participants

	return resp, nil
}

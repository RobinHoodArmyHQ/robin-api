package event

import (
	"context"
	"errors"
	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/participants"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/user"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/ctxmeta"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

func (i *impl) MarkEventInterested(ctx context.Context, eventID nanoid.NanoID) error {
	userID := ctxmeta.GetUser(ctx)
	ev := env.FromContext(ctx)

	userResp, err := ev.UserRepository.GetUser(&user.GetUserRequest{UserID: userID})
	if err != nil {
		return err
	}

	if userResp == nil || userResp.User == nil {
		return errors.New("user not found")
	}

	eventResp, err := ev.EventRepository.GetEvent(&event.GetEventRequest{EventID: eventID})
	if err != nil {
		return err
	}

	if eventResp == nil || eventResp.Event == nil {
		return errors.New("event not found")
	}

	// Check if user is already interested in the event
	interestResp, err := ev.ParticipantsRepository.GetEventParticipant(ctx, &participants.GetEventParticipantRequest{
		EventID: eventResp.Event.ID,
		UserID:  userResp.User.ID,
	})
	if err != nil {
		return err
	}

	if interestResp.Participant != nil {
		return errors.New("user is already interested in this event")
	}

	// Create participant
	participant := &models.Participant{
		EventID: eventResp.Event.ID,
		UserID:  userResp.User.ID,
		Status:  models.StatusInterested,
		Role:    models.RoleVolunteer,
	}

	if err = ev.ParticipantsRepository.CreateParticipant(ctx, &participants.CreateParticipantsRequest{
		Participant: participant,
	}); err != nil {
		return err
	}

	return nil
}

func (i *impl) GetEventParticipants(ctx context.Context, eventID nanoid.NanoID) (*GetEventParticipantsResponse, error) {
	ev := env.FromContext(ctx)

	eventResp, err := ev.EventRepository.GetEvent(&event.GetEventRequest{EventID: eventID})
	if err != nil {
		return nil, err
	}

	if eventResp == nil || eventResp.Event == nil {
		return nil, errors.New("event not found")
	}

	participantsResp, err := ev.ParticipantsRepository.GetEventParticipants(ctx, eventResp.Event.ID)
	if err != nil {
		return nil, err
	}

	return &GetEventParticipantsResponse{
		Participants: participantsResp.Participants,
	}, nil
}

package event

import (
	"errors"
	"fmt"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/models"
)

const (
	nameMaxLength        = 100
	descriptionMaxLength = 500

	timestampMin = 1 * time.Hour
	timestampMax = 30 * 24 * time.Hour
)

func validateCreateEventRequest(req *models.Event) error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if len(req.Name) > nameMaxLength {
		return fmt.Errorf("name should be less than %d characters", nameMaxLength)
	}

	if req.Description == "" {
		return errors.New("description is required")
	}
	if len(req.Description) > descriptionMaxLength {
		return fmt.Errorf("description should be less than %d characters", descriptionMaxLength)
	}

	if req.Timestamp == 0 {
		return errors.New("timestamp is required")
	}
	if time.Unix(req.Timestamp, 0).Before(time.Now().Add(timestampMin)) {
		return fmt.Errorf("timestamp should be at least %s from now", timestampMin)
	}
	if time.Unix(req.Timestamp, 0).After(time.Now().Add(timestampMax)) {
		return fmt.Errorf("timestamp should be at most %s from now", timestampMax)
	}

	if req.EventType == 0 {
		return errors.New("eventType is required")
	}

	// TODO: validations on remaining fields
	return nil
}

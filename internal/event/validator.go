package event

import (
	"errors"
	"fmt"
	"time"

	"github.com/RobinHoodArmyHQ/robin-api/models"
)

const (
	nameMaxLength         = 100
	descriptionMaxLength  = 500
	locationNameMaxLength = 100

	maxRobins uint8 = 200

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

	if req.StartTime.IsZero() {
		return errors.New("start time is required")
	}
	if req.StartTime.Before(time.Now().Add(timestampMin)) {
		return fmt.Errorf("start time should be at least %s from now", timestampMin)
	}
	if req.StartTime.After(time.Now().Add(timestampMax)) {
		return fmt.Errorf("start time should be at most %s from now", timestampMax)
	}

	if req.EventType == 0 || !models.ValidEventTypes[req.EventType] {
		return errors.New("invalid event type provided")
	}

	if req.MinRobins == 0 || req.MinRobins > maxRobins {
		return fmt.Errorf("min robins should be less than or equal to %d", maxRobins)
	}
	if req.MaxRobins == 0 || req.MaxRobins > maxRobins {
		return fmt.Errorf("max robins should be less than or equal to %d", maxRobins)
	}
	if req.MaxRobins < req.MinRobins {
		return errors.New("max robins should be greater than or equal to min robins")
	}

	if req.EventLocation == nil {
		return errors.New("event location is required")
	}
	if req.EventLocation.Name == "" {
		return errors.New("event location name is required")
	}
	if len(req.EventLocation.Name) > locationNameMaxLength {
		return fmt.Errorf("event location name should be less than %d characters", locationNameMaxLength)
	}
	if req.EventLocation.Latitude == 0 {
		return errors.New("event location latitude is required")
	}
	if req.EventLocation.Longitude == 0 {
		return errors.New("event location longitude is required")
	}

	return nil
}

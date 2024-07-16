package repositories

import "github.com/RobinHoodArmyHQ/robin-api/models"

type LocationRepository interface {
	GetCities(req *GetCitiesRequest) (*GetCitiesResponse, error)
}

type GetCitiesRequest struct {
}

type GetCitiesResponse struct {
	Cities []*models.City
}

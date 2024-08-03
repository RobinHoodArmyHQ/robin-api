package repositories

import (
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
)

type LocationRepository interface {
	GetCities(req *GetCitiesRequest) (*GetCitiesResponse, error)
	GetCityByID(req *GetCityByCityIDRequest) (*GetCityByCityIDResponse, error)
}

type GetCitiesRequest struct {
}

type GetCitiesResponse struct {
	Cities []*models.City
}

type GetCityByCityIDRequest struct {
	CityID nanoid.NanoID
}

type GetCityByCityIDResponse struct {
	City *models.City
}

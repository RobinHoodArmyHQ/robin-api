package sql

import (
	"fmt"

	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"go.uber.org/zap"
)

// LocationRepository is the repository for locations
type LocationRepository struct {
	logger *zap.Logger
	db     *database.SqlDB
}

// NewEventRepository creates a new EventRepository
func NewLocationRepository(logger *zap.Logger, db *database.SqlDB) *LocationRepository {
	return &LocationRepository{logger, db}
}

func (r *LocationRepository) GetCities(req *repositories.GetCitiesRequest) (*repositories.GetCitiesResponse, error) {
	cities := make([]*models.City, 0)

	exec := r.db.Master().Preload("Country").Find(&cities)
	if exec.Error != nil {
		return nil, fmt.Errorf("failed to get cities: %v", exec.Error)
	}

	return &repositories.GetCitiesResponse{
		Cities: cities,
	}, nil
}

func (r *LocationRepository) GetCityByID(req *repositories.GetCityByCityIDRequest) (*repositories.GetCityByCityIDResponse, error) {
	city := &models.City{}

	exec := r.db.Master().Preload("Country").First(city, "city_id = ?", req.CityID)
	if exec.Error != nil {
		return nil, fmt.Errorf("failed to get city: %v", exec.Error)
	}

	return &repositories.GetCityByCityIDResponse{
		City: city,
	}, nil
}

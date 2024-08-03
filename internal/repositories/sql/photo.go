package sql

import (
	"fmt"

	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/database"
	"go.uber.org/zap"
)

type PhotoRepository struct {
	logger *zap.Logger
	db     *database.SqlDB
}

func NewPhotoRepository(logger *zap.Logger, db *database.SqlDB) *PhotoRepository {
	return &PhotoRepository{logger, db}
}

// CreateEvent creates a new event within a transaction
func (r *PhotoRepository) CreatePhotoUpload(req *repositories.PhotoUploadsRequest) error {
	err := r.db.Master().Table("photo_uploads").Create(req).Error
	if err != nil {
		return fmt.Errorf("database query failed: %v", err)
	}

	return nil
}

package repositories

import (
	"time"
)

type PhotoUploadsRequest struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement"`
	S3Prefix  string     `gorm:"column:s3_prefix"`
	S3Key     string     `gorm:"column:s3_key"`
	UserId    int64      `gorm:"column:user_id;index"`
	EventId   int64      `gorm:"column:event_id"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

type PhotoRepository interface {
	CreatePhotoUpload(req *PhotoUploadsRequest) error
}

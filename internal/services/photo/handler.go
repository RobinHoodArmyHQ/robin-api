package photo

import (
	"mime/multipart"
	"net/http"

	"github.com/RobinHoodArmyHQ/robin-api/internal/env"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories"
	"github.com/RobinHoodArmyHQ/robin-api/internal/repositories/event"
	"github.com/RobinHoodArmyHQ/robin-api/models"
	"github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type PhotoUploadRequest struct {
	File    *multipart.FileHeader `form:"file" binding:"required"`
	EventId nanoid.NanoID         `form:"event_id"`
}
type PhotoUploadResponse struct {
	Status *models.Status `json:"status"`
	Name   string         `json:"name"`
}

func PhotoUploadHandler(c *gin.Context) {
	ev := env.FromContext(c)

	req := &PhotoUploadRequest{}
	if err := c.ShouldBind(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed(err.Error()))
		return
	}

	// get the numeric event id
	event, err := ev.EventRepository.GetEvent(&event.GetEventRequest{EventID: req.EventId})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, models.StatusFailed(err.Error()))
		return
	}

	// generate a new key to use as name for the uploaded file
	key, err := nanoid.GetID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	file, err := req.File.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	// upload file to S3
	_, err = ev.S3Service().PutObject(&s3.PutObjectInput{
		Bucket: aws.String(viper.GetString("s3.bucket_photos")),
		Key:    aws.String(S3Prefix_CheckinPhotos + key.String()),
		Body:   file,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	// save the photo upload record in the database
	err = ev.PhotoRepository().CreatePhotoUpload(&repositories.PhotoUploadsRequest{
		S3Prefix: S3Prefix_CheckinPhotos,
		S3Key:    key.String(),
		UserId:   c.GetInt64("user_id_numeric"), // TODO: get the user id from the context
		EventId:  event.Event.ID,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, models.StatusFailed(err.Error()))
		return
	}

	c.JSON(http.StatusOK, &PhotoUploadResponse{
		Status: models.StatusSuccess(),
		Name:   key.String(),
	})
}

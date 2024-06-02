package models

import "github.com/google/uuid"

type User struct {
	UserId       uuid.UUID    `json:"user_id,omitempty"`
	FirstName    string       `json:"first_name,omitempty"`
	LastName     string       `json:"last_name,omitempty"`
	AvatarURL    string       `json:"avatar_url,omitempty"`
	MobileNumber MobileNumber `json:"mobile_number,omitempty"`
	EmailId      string       `json:"email_id,omitempty"`
	FacebookId   string       `json:"facebook_id,omitempty"`
	TwitterId    string       `json:"twitter_id,omitempty"`
	InstagramId  string       `json:"instagram_id,omitempty"`
	Level        Level        `json:"level,omitempty"`
	NumDrives    uint         `json:"num_drives,omitempty"`
	DefaultCity  City         `json:"default_city,omitempty"`
}

type Level struct {
	Name          string `json:"name,omitempty"`
	BadgeImageURL string `json:"badge_image_url,omitempty"`
}

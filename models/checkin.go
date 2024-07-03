package models

type CheckIn struct {
	CheckInID         uint64   `json:"check_in_id,omitempty" gorm:"primaryKey"`
	UserID            uint64   `json:"user_id,omitempty"`
	EventID           uint64   `json:"event_id,omitempty"`
	PhotoURLs         []string `json:"photo_urls,omitempty" gorm:"serializer:json"`
	Description       string   `json:"description,omitempty"`
	NoOfPeopleServed  uint64   `json:"no_of_people_served,omitempty"`
	NoOfStudentTaught uint64   `json:"no_of_student_taught,omitempty"`
	CreatedAt         string   `json:"created_at,omitempty"`
	UpdatedAt         string   `json:"updated_at,omitempty"`
}

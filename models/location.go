package models

type Location struct {
	LocationId uint64  `json:"location_id,omitempty" gorm:"primaryKey"`
	Name       string  `json:"name,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Longitude  float64 `json:"longitude,omitempty"`
}

type City struct {
	CityId  string  `json:"city_id,omitempty"`
	Name    string  `json:"name,omitempty"`
	Country Country `json:"country,omitempty"`
}

type Country struct {
	CountryId string `json:"country_id,omitempty"`
	Name      string `json:"name,omitempty"`
}

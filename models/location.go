package models

import "github.com/RobinHoodArmyHQ/robin-api/pkg/nanoid"

type Location struct {
	ID            uint64  `json:"-" gorm:"primaryKey;auto_increment"`
	Name          string  `json:"name,omitempty"`
	Latitude      float64 `json:"latitude,omitempty"`
	Longitude     float64 `json:"longitude,omitempty"`
	GooglePlaceID string  `json:"google_place_id,omitempty"`
}

type City struct {
	ID        int32         `json:"-" gorm:"primaryKey;auto_increment"`
	CityId    nanoid.NanoID `json:"city_id,omitempty"`
	Name      string        `json:"name,omitempty"`
	CountryId int8          `json:"-"`
	Country   Country       `json:"country,omitempty" gorm:"foreignKey:ID;references:CountryId"`
}

type Country struct {
	ID        int8          `json:"-" gorm:"primaryKey;auto_increment"`
	CountryId nanoid.NanoID `json:"country_id,omitempty"`
	Name      string        `json:"name,omitempty"`
}

type GetCitiesRequest struct {
}

type GetCitiesResponse struct {
	Status *Status `json:"status,omitempty"`
	Cities []*City `json:"cities,omitempty"`
}

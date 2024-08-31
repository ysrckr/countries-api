package models

type Maps struct {
	GoogleMaps     string `json:"googleMaps,omitempty" bson:"googleMaps,omitempty"`
	OpenStreetMaps string `json:"openStreetMaps,omitempty" bson:"openStreetMaps,omitempty"`
}

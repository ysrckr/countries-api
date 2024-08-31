package models

type CapitalInfo struct {
	Latlng []float64 `json:"latlng,omitempty" bson:"latlng,truncate,omitempty"`
}

package models

type CoatOfArms struct {
	Svg string `json:"svg,omitempty" bson:"svg,omitempty"`
	Png string `json:"png,omitempty" bson:"png,omitempty"`
}

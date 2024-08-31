package models

type Flags struct {
	Svg string `json:"svg,omitempty" bson:"svg,omitempty"`
	Png string `json:"png,omitempty" bson:"png,omitempty"`
	Alt string `json:"alt,omitempty" bson:"alt,omitempty"`
}

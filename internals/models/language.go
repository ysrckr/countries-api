package models

type Language struct {
	Language string `json:"language,omitempty" bson:"language,omitempty"`
	Code     string `json:"code,omitempty" bson:"code,omitempty"`
}

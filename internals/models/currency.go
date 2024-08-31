package models

type Currency struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Symbol string `json:"symbol,omitempty" bson:"symbol,omitempty"`
	Code   string `json:"code,omitempty" bson:"code,omitempty"`
}

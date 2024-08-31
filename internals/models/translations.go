package models

type Translation struct {
	Official string `json:"official,omitempty" bson:"official,omitempty"`
	Common   string `json:"common,omitempty" bson:"common,omitempty"`
	Code     string `json:"code,omitempty" bson:"code,omitempty"`
}

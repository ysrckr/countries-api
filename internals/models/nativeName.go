package models

type NativeName struct {
	Nld struct {
		Official string `json:"official,omitempty" bson:"official,omitempty"`
		Common   string `json:"common,omitempty" bson:"common,omitempty"`
	} `json:"nld,omitempty" bson:"nld,omitempty"`
	Pap struct {
		Official string `json:"official,omitempty" bson:"official,omitempty"`
		Common   string `json:"common,omitempty" bson:"common,omitempty"`
	} `json:"pap,omitempty" bson:"pap,omitempty"`
}

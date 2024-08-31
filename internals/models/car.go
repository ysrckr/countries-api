package models

type Car struct {
	Signs []any  `json:"signs,omitempty" bson:"signs,omitempty"`
	Side  string `json:"side,omitempty" bson:"side,omitempty"`
}

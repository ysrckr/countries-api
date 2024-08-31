package models

type PostalCode struct {
	Format any `json:"format,omitempty" bson:"format,omitempty"`
	Regex  any `json:"regex,omitempty" bson:"regex,omitempty"`
}

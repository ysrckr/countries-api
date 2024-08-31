package models

type Idd struct {
	Root     string   `json:"root,omitempty" bson:"root,omitempty"`
	Suffixes []string `json:"suffixes,omitempty" bson:"suffixes,omitempty"`
}

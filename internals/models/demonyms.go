package models

type Demonyms struct {
	Eng struct {
		F string `json:"f,omitempty" bson:"f,omitempty"`
		M string `json:"m,omitempty" bson:"m,omitempty"`
	} `json:"eng,omitempty" bson:"eng,omitempty"`
	Fra struct {
		F string `json:"f,omitempty" bson:"f,omitempty"`
		M string `json:"m,omitempty" bson:"m,omitempty"`
	} `json:"fra,omitempty" bson:"fra,omitempty"`
}

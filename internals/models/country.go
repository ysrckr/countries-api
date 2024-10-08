package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name struct {
		Common     string      `json:"common" bson:"common,omitempty"`
		Official   string      `json:"official,omitempty" bson:"official,omitempty"`
		NativeName *NativeName `json:"nativeName,omitempty" bson:"nativeName,omitempty"`
	} `json:"name,omitempty" bson:"name,omitempty"`
	Tld          []string               `json:"tld,omitempty" bson:"tld,omitempty"`
	Cca2         string                 `json:"cca2" bson:"cca2"`
	Ccn3         string                 `json:"ccn3,omitempty" bson:"ccn3,omitempty"`
	Cca3         string                 `json:"cca3,omitempty" bson:"cca3,omitempty"`
	Cioc         string                 `json:"cioc,omitempty" bson:"cioc,omitempty"`
	Fifa         string                 `json:"fifa,omitempty" bson:"fifa,omitempty"`
	Independent  bool                   `json:"independent,omitempty" bson:"independent,omitempty"`
	Status       string                 `json:"status,omitempty" bson:"status,omitempty"`
	UnMember     bool                   `json:"unMember,omitempty" bson:"unMember,omitempty"`
	Idd          *Idd                   `json:"idd,omitempty" bson:"idd,omitempty"`
	Capital      []string               `json:"capital,omitempty" bson:"capital,omitempty"`
	CapitalInfo  *CapitalInfo           `json:"capitalInfo,omitempty" bson:"capitalInfo,omitempty"`
	AltSpellings []string               `json:"altSpellings,omitempty" bson:"altSpellings,omitempty"`
	Region       string                 `json:"region,omitempty" bson:"region,omitempty"`
	Subregion    string                 `json:"subregion,omitempty" bson:"subregion,omitempty"`
	Continents   []string               `json:"continents,omitempty" bson:"continents,omitempty"`
	Language     map[string]string      `json:"languages,omitempty" bson:"languages,omitempty"`
	Translations map[string]Translation `json:"translations,omitempty" bson:"translations,omitempty"`
	Latlng       []float64              `json:"latlng,omitempty" bson:"latlng,truncate,omitempty"`
	Landlocked   bool                   `json:"landlocked,omitempty" bson:"landlocked,omitempty"`
	Borders      []any                  `json:"borders,omitempty" bson:"borders,omitempty"`
	Area         int                    `json:"area,omitempty" bson:"area,truncate,omitempty"`
	Flag         string                 `json:"flag,omitempty" bson:"flag,omitempty"`
	Flags        *Flags                 `json:"flags,omitempty" bson:"flags,omitempty"`
	CoatOfArms   *CoatOfArms            `json:"coatOfArms,omitempty" bson:"coatOfArms,omitempty"`
	Population   int                    `json:"population,omitempty" bson:"population,truncate,omitempty"`
	Maps         *Maps                  `json:"maps,omitempty" bson:"maps,omitempty"`
	Gini         *struct{}              `json:"gini,omitempty" bson:"gini,omitempty"`
	Car          *Car                   `json:"car,omitempty" bson:"car,omitempty"`
	PostalCode   *PostalCode            `json:"postalCode,omitempty" bson:"postalCode,omitempty"`
	StartOfWeek  string                 `json:"startOfWeek,omitempty" bson:"startOfWeek,omitempty"`
	Timezones    []string               `json:"timezones,omitempty" bson:"timezones,omitempty"`
	Currencies   map[string]Currency    `json:"currencies,omitempty"`
	Demonyms     *Demonyms              `json:"demonyms,omitempty" bson:"demonyms,omitempty"`
}

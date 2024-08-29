package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name struct {
		Common     string `bson:"common,omitempty"`
		Official   string `bson:"official,omitempty"`
		NativeName struct {
			Nld struct {
				Official string `bson:"official,omitempty"`
				Common   string `bson:"common,omitempty"`
			} `bson:"nld,omitempty"`
			Pap struct {
				Official string `bson:"official,omitempty"`
				Common   string `bson:"common,omitempty"`
			} `bson:"pap,omitempty"`
		} `bson:"nativeName,omitempty"`
	} `bson:"name,omitempty"`
	Tld         []string `json:"tld" bson:"tld,omitempty"`
	Cca2        string   `bson:"cca2"`
	Ccn3        string   `bson:"ccn3,omitempty"`
	Cca3        string   `bson:"cca3,omitempty"`
	Cioc        string   `bson:"cioc,omitempty"`
	Fifa        string   `bson:"fifa,omitempty"`
	Independent bool     `bson:"independent,omitempty"`
	Status      string   `bson:"status,omitempty"`
	UnMember    bool     `bson:"unMember,omitempty"`
	Idd         struct {
		Root     string   `bson:"root,omitempty"`
		Suffixes []string `bson:"suffixes,omitempty"`
	} `bson:"idd,omitempty"`
	Capital     []string `bson:"capital,omitempty"`
	CapitalInfo struct {
		Latlng []float64 `bson:"latlng,truncate,omitempty"`
	} `bson:"capitalInfo,omitempty"`
	AltSpellings []string `bson:"altSpellings,omitempty"`
	Region       string   `bson:"region,omitempty"`
	Subregion    string   `bson:"subregion,omitempty"`
	Continents   []string `bson:"continents,omitempty"`
	Languages    struct {
		Nld string `bson:"nld,omitempty"`
		Pap string `bson:"pap,omitempty"`
	} `bson:"languages,omitempty"`
	Translations struct {
		Ara struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"ara"`
		Bre struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"bre"`
		Ces struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"ces"`
		Cym struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"cym"`
		Deu struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"deu"`
		Est struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"est"`
		Fin struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"fin"`
		Fra struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"fra"`
		Hrv struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"hrv"`
		Hun struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"hun"`
		Ita struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"ita"`
		Jpn struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"jpn"`
		Kor struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"kor"`
		Nld struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"nld"`
		Per struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"per"`
		Pol struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"pol"`
		Por struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"por"`
		Rus struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"rus"`
		Slk struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"slk"`
		Spa struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"spa"`
		Srp struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"srp"`
		Swe struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"swe"`
		Tur struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"tur"`
		Urd struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"urd"`
		Zho struct {
			Official string `bson:"official,omitempty"`
			Common   string `bson:"common"`
		} `bson:"zho"`
	} `bson:"translations"`
	Latlng     []float64 `bson:"latlng,truncate,omitempty"`
	Landlocked bool      `bson:"landlocked,omitempty"`
	Borders    []any     `bson:"borders,omitempty"`
	Area       int       `bson:"area,truncate,omitempty"`
	Flag       string    `bson:"flag"`
	Demonyms   struct {
		Eng struct {
			F string `bson:"f,omitempty"`
			M string `bson:"m,omitempty"`
		} `bson:"eng,omitempty"`
		Fra struct {
			F string `bson:"f,omitempty"`
			M string `bson:"m,omitempty"`
		} `bson:"fra,omitempty"`
	} `bson:"demonyms,omitempty"`
	Flags struct {
		Svg string `bson:"svg"`
		Png string `bson:"png"`
		Alt string `bson:"alt"`
	} `bson:"flags"`
	CoatOfArms struct {
		Svg string `bson:"svg,omitempty"`
		Png string `bson:"png,omitempty"`
	} `bson:"coatOfArms,omitempty"`
	Population int `bson:"population,truncate,omitempty"`
	Maps       struct {
		GoogleMaps     string `bson:"googleMaps"`
		OpenStreetMaps string `bson:"openStreetMaps"`
	} `bson:"maps"`
	Gini struct{} `bson:"gini,omitempty"`
	Car  struct {
		Signs []any  `bson:"signs,omitempty"`
		Side  string `bson:"side,omitempty"`
	} `bson:"car,omitempty"`
	PostalCode struct {
		Format any `bson:"format,omitempty"`
		Regex  any `bson:"regex,omitempty"`
	} `bson:"postalCode,omitempty"`
	StartOfWeek string   `bson:"startOfWeek,omitempty"`
	Timezones   []string `bson:"timezones,omitempty"`
	Currencies  struct {
		Awg struct {
			Name   string `json:"name,omitempty"`
			Symbol string `json:"symbol,omitempty"`
		} `json:"AWG,omitempty"`
	} `json:"currencies,omitempty"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Country struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name struct {
		Common     string `jsson:"common" bson:"common"`
		Official   string `jsson:"official" bson:"official"`
		NativeName struct {
			Nld struct {
				Official string `json:"official" bson:"official"`
				Common   string `json:"common" bson:"common"`
			} `json:"nld" bson:"nld"`
			Pap struct {
				Official string `json:"official" bson:"official"`
				Common   string `bson:"common"`
			} `bson:"pap"`
		} `bson:"nativeName"`
	} `bson:"name"`
	Tld         []string `bson:"tld"`
	Cca2        string   `json:"cca2" bson:"cca2"`
	Ccn3        string   `bson:"ccn3"`
	Cca3        string   `bson:"cca3"`
	Cioc        string   `bson:"cioc"`
	Fifa        string   `bson:"fifa"`
	Independent bool     `bson:"independent"`
	Status      string   `bson:"status"`
	UnMember    bool     `bson:"unMember"`
	Idd         struct {
		Root     string   `bson:"root"`
		Suffixes []string `bson:"suffixes"`
	} `bson:"idd"`
	Capital     []string `bson:"capital"`
	CapitalInfo struct {
		Latlng []float64 `bson:"latlng,truncate,omitempty"`
	} `bson:"capitalInfo"`
	AltSpellings []string `bson:"altSpellings"`
	Region       string   `bson:"region"`
	Subregion    string   `bson:"subregion"`
	Continents   []string `bson:"continents"`
	Languages    struct {
		Nld string `bson:"nld"`
		Pap string `bson:"pap"`
	} `bson:"languages"`
	Translations struct {
		Ara struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"ara"`
		Bre struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"bre"`
		Ces struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"ces"`
		Cym struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"cym"`
		Deu struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"deu"`
		Est struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"est"`
		Fin struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"fin"`
		Fra struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"fra"`
		Hrv struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"hrv"`
		Hun struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"hun"`
		Ita struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"ita"`
		Jpn struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"jpn"`
		Kor struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"kor"`
		Nld struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"nld"`
		Per struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"per"`
		Pol struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"pol"`
		Por struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"por"`
		Rus struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"rus"`
		Slk struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"slk"`
		Spa struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"spa"`
		Srp struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"srp"`
		Swe struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"swe"`
		Tur struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"tur"`
		Urd struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"urd"`
		Zho struct {
			Official string `bson:"official"`
			Common   string `bson:"common"`
		} `bson:"zho"`
	} `bson:"translations"`
	Latlng     []float64 `bson:"latlng,truncate,omitempty"`
	Landlocked bool      `bson:"landlocked"`
	Borders    []any     `bson:"borders"`
	Area       int       `bson:"area,truncate"`
	Flag       string    `bson:"flag"`
	Demonyms   struct {
		Eng struct {
			F string `bson:"f"`
			M string `bson:"m"`
		} `bson:"eng"`
		Fra struct {
			F string `bson:"f"`
			M string `bson:"m"`
		} `bson:"fra"`
	} `bson:"demonyms"`
	Flags struct {
		Svg string `bson:"svg"`
		Png string `bson:"png"`
		Alt string `bson:"alt"`
	} `bson:"flags"`
	CoatOfArms struct {
		Svg string `bson:"svg"`
		Png string `bson:"png"`
	} `bson:"coatOfArms"`
	Population int `bson:"population,truncate"`
	Maps       struct {
		GoogleMaps     string `bson:"googleMaps"`
		OpenStreetMaps string `bson:"openStreetMaps"`
	} `bson:"maps"`
	Gini struct{} `bson:"gini"`
	Car  struct {
		Signs []any  `bson:"signs"`
		Side  string `bson:"side"`
	} `bson:"car"`
	PostalCode struct {
		Format any `bson:"format"`
		Regex  any `bson:"regex"`
	} `bson:"postalCode"`
	StartOfWeek string   `bson:"startOfWeek"`
	Timezones   []string `bson:"timezones"`
	Currencies  struct {
		Awg struct {
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"AWG"`
	} `json:"currencies"`
}

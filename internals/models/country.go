package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type QueryCountry struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name struct {
		Common   string `json:"common" bson:"common,omitempty"`
		Official string `json:"official,omitempty" bson:"official,omitempty"`
	} `json:"name,omitempty" bson:"name,omitempty"`
	Cca2  string `json:"cca2" bson:"cca2"`
	Flag  string `json:"flag,omitempty" bson:"flag,omitempty"`
	Flags struct {
		Svg string `json:"svg,omitempty" bson:"svg,omitempty"`
		Png string `json:"png,omitempty" bson:"png,omitempty"`
		Alt string `json:"alt,omitempty" bson:"alt,omitempty"`
	} `json:"flags,omitempty" bson:"flags,omitempty"`
	Translations struct {
		Tur struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"tur,omitempty" bson:"tur,omitempty"`
	} `json:"translations" bson:"translations"`
}

type Country struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name struct {
		Common     string `json:"common" bson:"common,omitempty"`
		Official   string `json:"official,omitempty" bson:"official,omitempty"`
		NativeName struct {
			Nld struct {
				Official string `json:"official,omitempty" bson:"official,omitempty"`
				Common   string `json:"common,omitempty" bson:"common,omitempty"`
			} `json:"nld,omitempty" bson:"nld,omitempty"`
			Pap struct {
				Official string `json:"official,omitempty" bson:"official,omitempty"`
				Common   string `json:"common,omitempty" bson:"common,omitempty"`
			} `json:"pap,omitempty" bson:"pap,omitempty"`
		} `json:"nativeName,omitempty" bson:"nativeName,omitempty"`
	} `json:"name,omitempty" bson:"name,omitempty"`
	Tld         []string `json:"tld,omitempty" bson:"tld,omitempty"`
	Cca2        string   `json:"cca2" bson:"cca2"`
	Ccn3        string   `json:"ccn3,omitempty" bson:"ccn3,omitempty"`
	Cca3        string   `json:"cca3,omitempty" bson:"cca3,omitempty"`
	Cioc        string   `json:"cioc,omitempty" bson:"cioc,omitempty"`
	Fifa        string   `json:"fifa,omitempty" bson:"fifa,omitempty"`
	Independent bool     `json:"independent,omitempty" bson:"independent,omitempty"`
	Status      string   `json:"status,omitempty" bson:"status,omitempty"`
	UnMember    bool     `json:"unMember,omitempty" bson:"unMember,omitempty"`
	Idd         struct {
		Root     string   `json:"root,omitempty" bson:"root,omitempty"`
		Suffixes []string `json:"suffixes,omitempty" bson:"suffixes,omitempty"`
	} `json:"idd,omitempty" bson:"idd,omitempty"`
	Capital     []string `json:"capital,omitempty" bson:"capital,omitempty"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng,omitempty" bson:"latlng,truncate,omitempty"`
	} `json:"capitalInfo,omitempty" bson:"capitalInfo,omitempty"`
	AltSpellings []string `json:"altSpellings,omitempty" bson:"altSpellings,omitempty"`
	Region       string   `json:"region,omitempty" bson:"region,omitempty"`
	Subregion    string   `json:"subregion,omitempty" bson:"subregion,omitempty"`
	Continents   []string `json:"continents,omitempty" bson:"continents,omitempty"`
	Languages    struct {
		Nld string `json:"nld,omitempty" bson:"nld,omitempty"`
		Pap string `json:"pap,omitempty" bson:"pap,omitempty"`
	} `json:"languages,omitempty" bson:"languages,omitempty"`
	Translations struct {
		Ara struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"ara,omitempty" bson:"ara,omitempty"`
		Bre struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"bre,omitempty" bson:"bre,omitempty"`
		Ces struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"ces,omitempty" bson:"ces,omitempty"`
		Cym struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"cym,omitempty" bson:"cym,omitempty"`
		Deu struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"deu,omitempty" bson:"deu,omitempty"`
		Est struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"est,omitempty" bson:"est,omitempty"`
		Fin struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"fin,omitempty" bson:"fin,omitempty"`
		Fra struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"fra,omitempty" bson:"fra,omitempty"`
		Hrv struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"hrv,omitempty" bson:"hrv,omitempty"`
		Hun struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"hun,omitempty" bson:"hun,omitempty"`
		Ita struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"ita,omitempty" bson:"ita,omitempty"`
		Jpn struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"jpn,omitempty" bson:"jpn,omitempty"`
		Kor struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"kor,omitempty" bson:"kor,omitempty"`
		Nld struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"nld,omitempty" bson:"nld,omitempty"`
		Per struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"per,omitempty" bson:"per,omitempty"`
		Pol struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"pol,omitempty" bson:"pol,omitempty"`
		Por struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"por,omitempty" bson:"por,omitempty"`
		Rus struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"rus,omitempty" bson:"rus,omitempty"`
		Slk struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"slk,omitempty" bson:"slk,omitempty"`
		Spa struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"spa,omitempty" bson:"spa,omitempty"`
		Srp struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"srp,omitempty" bson:"srp,omitempty"`
		Swe struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"swe,omitempty" bson:"swe,omitempty"`
		Tur struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"tur,omitempty" bson:"tur,omitempty"`
		Urd struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"urd,omitempty" bson:"urd,omitempty"`
		Zho struct {
			Official string `json:"official,omitempty" bson:"official,omitempty"`
			Common   string `json:"common,omitempty" bson:"common,omitempty"`
		} `json:"zho,omitempty" bson:"zho,omitempty"`
	} `json:"translations" bson:"translations"`
	Latlng     []float64 `json:"latlng,omitempty" bson:"latlng,truncate,omitempty"`
	Landlocked bool      `json:"landlocked,omitempty" bson:"landlocked,omitempty"`
	Borders    []any     `json:"borders,omitempty" bson:"borders,omitempty"`
	Area       int       `json:"area,omitempty" bson:"area,truncate,omitempty"`
	Flag       string    `json:"flag,omitempty" bson:"flag,omitempty"`
	Demonyms   struct {
		Eng struct {
			F string `json:"f,omitempty" bson:"f,omitempty"`
			M string `json:"m,omitempty" bson:"m,omitempty"`
		} `json:"eng,omitempty" bson:"eng,omitempty"`
		Fra struct {
			F string `json:"f,omitempty" bson:"f,omitempty"`
			M string `json:"m,omitempty" bson:"m,omitempty"`
		} `json:"fra,omitempty" bson:"fra,omitempty"`
	} `json:"demonyms,omitempty" bson:"demonyms,omitempty"`
	Flags struct {
		Svg string `json:"svg,omitempty" bson:"svg,omitempty"`
		Png string `json:"png,omitempty" bson:"png,omitempty"`
		Alt string `json:"alt,omitempty" bson:"alt,omitempty"`
	} `json:"flags,omitempty" bson:"flags,omitempty"`
	CoatOfArms struct {
		Svg string `json:"svg,omitempty" bson:"svg,omitempty"`
		Png string `json:"png,omitempty" bson:"png,omitempty"`
	} `json:"coatOfArms,omitempty" bson:"coatOfArms,omitempty"`
	Population int `json:"population,omitempty" bson:"population,truncate,omitempty"`
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

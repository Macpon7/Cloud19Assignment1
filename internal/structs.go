package internal

type ResultSpeciesGBIF struct {
	ResultArray		[]SpeciesGBIF	`json:"results"`
}

type SpeciesGBIF struct {
	Key				int			`json:"speciesKey"`
	Kingdom			string		`json:"kingdom"`
	Phylum			string		`json:"phylum"`
	Order			string		`json:"order"`
	Family			string		`json:"family"`
	Genus			string		`json:"genus"`
}

type NameGBIF struct {
	ScientificName	string		`json:"scientificName"`
	CanonicalName	string		`json:"canonicalName"`
	Year			string		`json:"bracketYear"`
}

type Specie struct {
	Key				int			`json:"key"`
	Kingdom			string		`json:"kingdom"`
	Phylum			string		`json:"phylum"`
	Order			string		`json:"order"`
	Family			string		`json:"family"`
	Genus			string		`json:"genus"`
	ScientificName	string		`json:"scientificName"`
	CanonicalName	string		`json:"canonicalName"`
	Year			string		`json:"year"`
}

type ResultCountryGBIF struct {
	ResultArray		[]CountryGBIF	`json:"results"`
}

type CountryGBIF struct {
	Species 		string		`json:"species"`
	SpeciesKey		int			`json:"speciesKey"`
}

type CountryRCEU struct {
	Code 			string		`json:"alpha2Code"`
	CountryName 	string		`json:"name"`
	CountryFlag 	string		`json:"flag"`
}

type Country struct {
	Code 			string		`json:"code"`
	CountryName 	string		`json:"countryname"`
	CountryFlag 	string		`json:"countryflag"`
	Species 		[]string	`json:"species"`
	SpeciesKey 		[]int		`json:"speciesKey"`
}

type Diagnostics struct {
	Gbif			int			`json:"gbif"`
	Restcountries	int			`json:"restcountries"`
	Version			string		`json:"version"`
	Uptime			string		`json:"uptime"`
}
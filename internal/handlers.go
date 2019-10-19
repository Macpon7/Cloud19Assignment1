package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//This is the default handler for invalid requests
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default handler: invalid request received.")
	http.Error(w, "Invalid Request", http.StatusBadRequest)
}

/*
returns a list of species based on a country code and limit parameters
country/{:country_identifier}{?limit={:limit}}
*/
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	//Separates the url to get the country code
	parts := strings.Split(r.URL.Path, "/")
	countryCode := parts[len(parts)-1]

	//GETs information about the country from RCEU
	resp, err := http.Get("https://restcountries.eu/rest/v2/alpha/" + countryCode)
	if err != nil {
		http.Error(w, "Could not get country, Bad request", 400)
		return
	}

	dec := json.NewDecoder(resp.Body)
	var countryInfo CountryRCEU

	err = dec.Decode(&countryInfo)
	if err != nil {
		http.Error(w, "Could not decode json", 400)
		return
	}

	//GETs the species that have occurred recently in the country from GBIF
	resp, err = http.Get("http://api.gbif.org/v1/occurrence/search?country=" + countryCode)
	if err !=nil {
		http.Error(w,"Bad request", 400)
		return
	}

	dec = json.NewDecoder(resp.Body)
	var speciesInfo ResultCountryGBIF

	err = dec.Decode(&speciesInfo)
	if err != nil  {
		http.Error(w, "Could not decode json", 400)
		return
	}

	//Combines the two results from the get requests into one output object
	countryOut := CountryFmt(speciesInfo, countryInfo)

	//Adds JSON to the header and encodes the output object to JSON
	http.Header.Add(w.Header(), "content-type", "application/json")
	err = json.NewEncoder(w).Encode(countryOut)
	if err != nil {
		http.Error(w,"Could not encode", 400)
	}
}

/*
provides information on specific species
species/{:speciesKey}
*/
func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	//Separates the URL to get the species key
	parts := strings.Split(r.URL.Path, "/")
	speciesKey := parts[len(parts)-1]

	//GETs the species genetic information
	resp, err := http.Get("http://api.gbif.org/v1/occurrence/search?speciesKey=" + speciesKey)
	if err != nil {
		http.Error(w, "Bad request species", 400)
		return
	}
	dec1 := json.NewDecoder(resp.Body)
	var responsesGBIF ResultSpeciesGBIF

	err = dec1.Decode(&responsesGBIF)
	if err != nil {
		http.Error(w, "Could not decode species json", 400)
		return
	}

	//GETs the species name and year information
	resp, err = http.Get("http://api.gbif.org/v1/species/" + speciesKey + "/name")
	if err != nil {
		http.Error(w,"Bad Request name", 400)
		return
	}

	dec := json.NewDecoder(resp.Body)
	var namesGBIF NameGBIF

	err = dec.Decode(&namesGBIF)
	if err != nil {
		http.Error(w, "Could not decode names json", 400)
		return
	}

	//Combines the results from the two GET requests into one output object
	speciesOut := SpeciesFmt(responsesGBIF, namesGBIF)

	//Adds JSON to the header and encodes the output object to JSON
	http.Header.Add(w.Header(), "content-type", "application/json")
	err = json.NewEncoder(w).Encode(speciesOut)
	if err!= nil {
		http.Error(w, "Could not encode json", 400)
		return
	}
}




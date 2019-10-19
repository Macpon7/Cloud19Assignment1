package main

import (
	"encoding/json"
	"time"
	"fmt"
	"log"
	"net/http"
	"os"
	"CO1/internal"
)

var Start time.Time
var Version string = "0.1"

func main () {
	Start = time.Now()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", internal.HandlerNil)
	http.HandleFunc("/conservation/v1/country/", internal.HandlerCountry)
	http.HandleFunc("/conservation/v1/species/", internal.HandlerSpecies)
	http.HandleFunc("/conservation/v1/diag/", HandlerDiag)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

/*
Indicates the availability of the services this service depends on
*/
func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	var diagOut internal.Diagnostics
	resp, err := http.Get("http://api.gbif.org/v1/")
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	diagOut.Gbif = resp.StatusCode

	resp, err = http.Get("http://restcountries.eu/")
	if err != nil {
		http.Error(w, "Bad request", 400)
		return
	}
	diagOut.Restcountries = resp.StatusCode

	diagOut.Version = Version

	elapsed := time.Since(Start)
	diagOut.Uptime = elapsed.String()

	http.Header.Add(w.Header(), "content-type", "application/json")
	err = json.NewEncoder(w).Encode(diagOut)
	if err!= nil {
		http.Error(w, "Could not encode json", 400)
		return
	}
}
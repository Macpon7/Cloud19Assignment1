package internal

import (
	//"fmt"
	//"encoding/json"
	//"net/http"
	//"log"
)

//Combines the variables from the two get requests
func SpeciesFmt (result ResultSpeciesGBIF, name NameGBIF) Specie {
	var out Specie
	out.Key = result.ResultArray[0].Key
	out.Kingdom = result.ResultArray[0].Kingdom
	out.Phylum = result.ResultArray[0].Phylum
	out.Order = result.ResultArray[0].Order
	out.Family = result.ResultArray[0].Family
	out.Genus = result.ResultArray[0].Genus
	out.ScientificName = name.ScientificName
	out.CanonicalName = name.CanonicalName
	out.Year = name.Year

	return out
}

//Combines the variables from the two get requests
func CountryFmt (result ResultCountryGBIF, country CountryRCEU) Country {
	var out Country
	out.Code = country.Code
	out.CountryName = country.CountryName
	out.CountryFlag = country.CountryFlag

	//Checks if the species and species key are duplicate before appending them
	speciesDuplicate := false
	speciesKeyDuplicate := false
	for i:=0; i<len(result.ResultArray); i++ {
		speciesDuplicate = false
		speciesKeyDuplicate = false
		if i>0 {
			for j:=0; j<len(out.Species); j++ {
				if out.Species[j]==result.ResultArray[i].Species {speciesDuplicate = true}
			}
			for j:=0; j<len(out.SpeciesKey); j++ {
				if out.SpeciesKey[j]==result.ResultArray[i].SpeciesKey {speciesKeyDuplicate = true}
			}
		}
		if !speciesDuplicate {out.Species = append(out.Species, result.ResultArray[i].Species)}
		if !speciesKeyDuplicate {out.SpeciesKey = append(out.SpeciesKey, result.ResultArray[i].SpeciesKey)}
	}

	return out
}
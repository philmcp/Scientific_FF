package main

import (
	//	"fmt"
	"log"
	"strings"
)

// Teams: Roto to DK
func teamRoto2DK(name string) string {
	teamMapping := map[string]string{
		"TOT": "spurs",
		"WHU": "west ham",
		"ARS": "arsenal",
		"BOU": "bournemouth",
		"NOR": "norwich",
		"NEW": "newcastle",
		"SOU": "southampton",
		"SWA": "swansea",
		"CHE": "chelsea",
		"WAT": "watford",
		"EVE": "everton",
		"MCI": "man city",
		"WBA": "west brom",
		"STK": "stoke",
		"CRY": "crystal palace",
		"MU":  "man utd",
		"AVL": "aston villa",
		"LEI": "leicester",
		"LIV": "liverpool",
		"SUN": "sunderland",
		"MID": "middlesborough",
		"HUL": "hull",
		"BUR": "burnley",
	}

	if _, exist := teamMapping[strings.ToUpper(name)]; !exist {
		log.Fatal("There is a problem with the mapping for " + name)
	}

	return strings.ToLower(strings.Replace(teamMapping[strings.ToUpper(name)], " ", "-", -1))
}

func teamPL2DK(name string) string {
	teamMapping := map[string]string{
		"spurs":          "TOT",
		"west ham":       "WHU",
		"arsenal":        "ARS",
		"bournemouth":    "BOU",
		"norwich":        "NOR",
		"newcastle":      "NEW",
		"southampton":    "SOU",
		"swansea":        "SWA",
		"chelsea":        "CHE",
		"watford":        "WAT",
		"everton":        "EVE",
		"man city":       "MCI",
		"west brom":      "WBA",
		"stoke":          "STK",
		"crystal palace": "CRY",
		"man utd":        "MU",
		"aston villa":    "AVL",
		"leicester":      "LEI",
		"liverpool":      "LIV",
		"sunderland":     "SUN",
		"burnley":        "BUR",
		"middlesborough": "MID",
		"hull":           "HUL",
	}

	if _, exist := teamMapping[strings.ToLower(name)]; !exist {
		log.Fatal("There is a problem with the mapping for " + name)
	}

	return strings.ToLower(teamMapping[strings.ToLower(name)])
}

func teamFFS2DK(name string) string {
	teamMapping := map[string]string{"sto": "stk", "sot": "sou", "mun": "mu"}

	if _, exist := teamMapping[name]; exist {
		//	fmt.Println("Mapping " + name + " to " + teamMapping[name])
		return teamMapping[name]
	} else {
		return name
	}
}

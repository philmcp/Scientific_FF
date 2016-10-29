package scrape

import (
	//	"fmt"
	"log"
	"strings"
)

var DKTeams = map[string]bool{
	"WHU": true,
	"CHE": true,
	"SUN": true,
	"WAT": true,
	"MID": true,
	"WBA": true,
	"HUL": true,
	"ARS": true,
	"BOU": true,
	"LEI": true,
	"BUR": true,
	"MU":  true,
	"SWA": true,
	"EVE": true,
	"TOT": true,
	"STK": true,
	"MCI": true,
	"CRY": true,
	"SOT": true,
	"LIV": true,
}

func teamRoto2DK(name string) string {
	name = strings.ToLower(name)
	teamMapping := map[string]string{
		"tottenham hotspur":    "TOT",
		"west ham united":      "WHU",
		"arsenal":              "ARS",
		"afc bournemouth":      "BOU",
		"norwich":              "NOR",
		"newcastle":            "NEW",
		"southampton":          "SOT",
		"swansea city":         "SWA",
		"chelsea":              "CHE",
		"watford":              "WAT",
		"everton":              "EVE",
		"manchester city":      "MCI",
		"west bromwich albion": "WBA",
		"stoke city":           "STK",
		"crystal palace":       "CRY",
		"manchester united":    "MU",
		"middlesbrough":        "MID",
		"aston villa":          "AVL",
		"leicester city":       "LEI",
		"liverpool":            "LIV",
		"sunderland":           "SUN",
		"burnley":              "BUR",
		"middlesborough":       "MID",
		"hull city":            "HUL",
	}

	if _, exist := teamMapping[name]; !exist {
		log.Fatal("Roto team name doesnt exist: " + name)
	}

	out := strings.ToLower(teamMapping[name])

	if _, exist := DKTeams[strings.ToUpper(out)]; !exist {
		log.Fatal("DK team name doesnt exist: " + strings.ToUpper(out))
	}

	return out
}

func teamFPL2DK(name string) string {
	teamMapping := map[string]string{
		"sto": "stk",
		"sou": "sot",
		"whu": "whu",
		"mun": "mu",
		"cry": "cry",
		"bou": "bou",
		"ars": "ars",
		"lei": "lei",
		"stk": "stk",
		"wba": "wba",
		"mci": "mci",
		"swa": "swa",
		"hul": "hul",
		"che": "che",
		"tot": "tot",
		"eve": "eve",
		"liv": "liv",
		"sun": "sun",
		"mid": "mid",
		"wat": "wat",
		"bur": "bur",
	}

	if _, exist := teamMapping[name]; !exist {
		log.Fatal("FPL team name doesnt exist: " + name)
	}

	out := strings.ToLower(teamMapping[name])

	if _, exist := DKTeams[strings.ToUpper(out)]; !exist {
		log.Fatal("DK team name doesnt exist: " + strings.ToUpper(out))
	}

	return out
}

func teamFFS2DK(name string) string {
	teamMapping := map[string]string{
		"sto": "stk",
		"sot": "sot",
		"whm": "whu",
		"mun": "mu",
		"cry": "cry",
		"bou": "bou",
		"ars": "ars",
		"lei": "lei",
		"stk": "stk",
		"wba": "wba",
		"mci": "mci",
		"swa": "swa",
		"hul": "hul",
		"che": "che",
		"tot": "tot",
		"eve": "eve",
		"liv": "liv",
		"sun": "sun",
		"mid": "mid",
		"wat": "wat",
		"bur": "bur",
	}

	if _, exist := teamMapping[name]; !exist {
		log.Fatal("FFS team name doesnt exist: " + name)
	}

	out := strings.ToLower(teamMapping[name])

	if _, exist := DKTeams[strings.ToUpper(out)]; !exist {
		log.Fatal("DK team name doesnt exist: " + strings.ToUpper(out))
	}

	return out
}

func mapDuplicateNames(name string) string {
	playerMapping := map[string]string{
		"christian benteke": "Benteke",
		"jonathan benteke":  "JBenteke",
		"pau l<f3>pez":      "Pau Lopez",
		"lewis cook":        "Cook",
		"steve cook":        "S_Cook",
		"adam smith":        "A_Smith",
		"bradley smith":     "B_Smith",
		"mark wilson":       "M_Wilson",
		"callum wilson":     "C_Wilson",
		"josh robson":       "J_Robson",
		"thomas robson":     "T_Robson",
	}

	if _, exist := playerMapping[strings.ToLower(name)]; exist {
		return playerMapping[strings.ToLower(name)]
	} else {
		return name
	}
}

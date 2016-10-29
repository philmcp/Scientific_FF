package models

import (
	"fmt"
	"strings"
)

type Lineup struct {
	Team    string
	OppTeam string
	Players []string
}

func NewLineup(text string) Lineup {
	fmt.Println(text)
	temp := strings.Split(text, ":")
	teams := strings.Split(temp[0], " XI")
	if len(temp) < 2 {
		return Lineup{}
	}

	players := strings.FieldsFunc(temp[1], func(r rune) bool {
		return r == ',' || r == ';'
	})

	p := []string{}
	for _, player := range players {
		p = append(p, strings.Trim(strings.Replace(player, ".", "", -1), " "))
	}

	opp := ""
	if len(teams) > 1 && teams[0] != "" {
		opp = strings.TrimSpace(strings.Replace(teams[1], "vs ", "", -1))

	}
	ret := Lineup{Team: strings.TrimSpace(teams[0]), OppTeam: opp, Players: p}
	return ret
}

package models

import (
	"fmt"
	"log"
	"strings"
)

// Player pool

type PlayerPool map[string]PlayerList

// Pick a random fantasy football team
func (pool PlayerPool) RandomLineup(minValue float64, formation map[string]int) Generated {
	team := PlayerPool{}
	projection := 0.0
	wage := 0.0

	used := map[string]bool{}
	usedTeams := map[string]bool{}

	for pos, num := range formation {
		isUtility := pos == "u"
		lineup := pool[pos].RandomPosition(used, minValue, num, isUtility)
		team[pos] = lineup
		for _, p := range lineup {
			projection += p.GetScore()
			wage += p.Wage
			used[p.getFullName()] = true
			usedTeams[p.Team] = true
		}
	}

	return Generated{Team: team, Wage: wage, Projection: projection, NumTeams: len(usedTeams)}
}

// I.e. cant have a striker against the teams goalkeaper
func (players PlayerPool) IsBalanced() bool {
	attackTeams := map[string]bool{}
	for _, player := range players["f"] {
		attackTeams[player.Team] = true
	}
	for _, player := range players["m"] {
		if player.PositionRaw == "m/f" {
			attackTeams[player.Team] = true
		}
	}
	for _, player := range players["u"] {
		if player.PositionRaw == "f" || player.PositionRaw == "m/f" {
			attackTeams[player.Team] = true
		}
	}

	for _, player := range players["d"] {
		if _, exists := attackTeams[player.OppTeam]; exists {
			return false
		}
	}
	for _, player := range players["gk"] {
		if _, exists := attackTeams[player.OppTeam]; exists {
			return false
		}
	}
	return true
}

func (p PlayerPool) Print() {
	log.Println("======= Start player pool ======= \n")
	for _, players := range p {
		for _, player := range players {
			player.print()
		}
	}
	log.Println("======= End player pool ======= ")
}

func (p PlayerPool) PrintTeamCounts() string {
	teams := map[string]int{}
	for _, players := range p {
		for _, player := range players {
			if _, exist := teams[player.Team]; exist {
				teams[player.Team]++
			} else {
				teams[player.Team] = 1
			}
		}
	}
	out := ""
	for team, cur := range teams {
		out += fmt.Sprintf("%d %s, ", cur, teamDK2Twitter(team))
	}
	out = strings.Trim(out, ", ")

	return out
}

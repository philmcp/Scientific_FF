package models

import (
	"log"
)

// Player pool

type PlayerPool map[string]PlayerList

// Pick a random fantasy football team
func (pool PlayerPool) RandomLineup(minValue float64, formation map[string]int) Lineup {
	team := PlayerPool{}
	projection := 0.0
	wage := 0.0

	used := map[string]bool{}
	usedTeams := map[string]bool{}

	for pos, num := range formation {
		lineup := pool[pos].RandomPosition(used, minValue, num)
		team[pos] = lineup
		for _, p := range lineup {
			projection += p.getScore()
			wage += p.Wage
			used[p.getFullName()] = true
			usedTeams[p.Team] = true
		}
	}

	return Lineup{Team: team, Wage: wage, Projection: projection, NumTeams: len(usedTeams)}
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

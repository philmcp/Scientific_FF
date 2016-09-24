package models

import (
	"fmt"
)

// Player pool
type PlayerPool map[string]PlayerList

// Pick a random fantasy football team
func (pool PlayerPool) RandomLineup(minValue float64, formation map[string]int) {
	team := PlayerPool{}
	projection := 0.0
	wage := 0.0

	used := map[string]bool{}
	usedTeams := map[string]bool{}

	for pos, num := range formation {
		lineup := pool[pos].randomPosition(used, minValue, num)
		team[pos] = lineup
		for _, p := range lineup {
			projection += p.getScore()
			wage += p.Wage
			used[p.getFullName()] = true
			usedTeams[p.Team] = true
		}
	}
}

func (p PlayerPool) Print() {
	fmt.Println("======= Start player pool ======= \n")
	for _, players := range p {
		for _, player := range players {
			player.print()
		}
	}
	fmt.Println("======= End player pool ======= ")

}

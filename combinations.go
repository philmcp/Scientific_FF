package main

import (
	"fmt"
	//	"os"
)

// Pick a random fantasy football team
func (pool PlayerPool) randTeamLineup(minValue float64) {
	team := PlayerPool{}
	projection := 0.0
	wage := 0.0

	used := map[string]bool{}

	for pos, num := range formation {
		lineup := pool[pos].randPosLineup(used, minValue, num)
		team[pos] = lineup
		for _, p := range lineup {
			projection += p.Projection
			wage += p.Wage
			used[p.getFullName()] = true
		}
	}

	//team.print()

	if wage <= MAX_WAGE && projection > BEST.Projection { //} && sufficientTeams(team) {
		fmt.Printf("\n***** New high score: %f Wage: $%f MinValue: %f\n", projection, wage, minValue)
		BEST.Wage = wage
		BEST.Projection = projection
		BEST.Team = team

		//	temp := create(team)
		//	temp.drawTeam()

		team.print()
	}
}

// Pick a random fantasy football lineup for a sepcific position (e.g. Defender)
func (players PlayerList) randPosLineup(ignore map[string]bool, minValue float64, num int) PlayerList {
	out := PlayerList{}

	for {
		rand := players[Random(0, len(players))]
		fullName := rand.getFullName()
		if _, exist := ignore[fullName]; !exist {

			out = append(out, rand)
			ignore[fullName] = true
		}
		if len(out) >= num {
			break
		}
	}

	return out
}

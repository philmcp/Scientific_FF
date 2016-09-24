package main

import (
	"fmt"
	//	"os"
)

// Pick a random fantasy football team
func (pool PlayerPool) randomLineup(minValue float64) {
	team := PlayerPool{}
	projection := 0.0
	wage := 0.0

	used := map[string]bool{}
	usedTeams := map[string]bool{}

	for pos, num := range conf.Formation {
		lineup := pool[pos].randPosLineup(used, minValue, num)
		team[pos] = lineup
		for _, p := range lineup {
			projection += p.getScore()
			wage += p.Wage
			used[p.getFullName()] = true
			usedTeams[p.Team] = true
		}
	}

	if wage <= conf.MaxWage && projection > bestLineup.Projection && len(usedTeams) >= conf.MinNumTeams {
		fmt.Printf("\n***** New high score: %f Wage: $%f MinValue: %f\n", projection, wage, minValue)
		bestLineup.Wage = wage
		bestLineup.Projection = projection
		bestLineup.Team = team

		team.drawTeam()

		team.print()
	}
}

// Pick a random fantasy football lineup for a sepcific position (e.g. Defender)
func (players PlayerList) randPosLineup(ignore map[string]bool, minValue float64, num int) PlayerList {
	out := PlayerList{}

	for {

		rand := players[Random(0, len(players))]
		if minValue > rand.Selected {
			continue
		}

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

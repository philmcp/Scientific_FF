package main

import (
	"fmt"
	//	"os"
)

// Pick a random fantasy football team
func (l *League) randTeamLineup(formation map[string]int, minValue float64) {
	team := []Player{}
	value := 0.0
	Wage := 0.0
	for pos, _ := range formation {
		if pos == "u" {
			continue
		}
		lineup := l.randPosLineup(pos, []Player{}, minValue)
		team = append(team, lineup...)
		for _, p := range lineup {
			value += p.Value
			Wage += p.Wage
		}
	}

	// Process utility players
	if _, ok := formation["u"]; ok {
		lineup := l.randPosLineup("u", team, minValue)
		team = append(team, lineup...)
		for _, p := range lineup {
			value += p.Value
			Wage += p.Wage

			if Wage > maxWage {
				return
			}
		}
	}

	if value > highestValue && sufficientTeams(team) {
		fmt.Printf("\n***** New high score: %f Wage: $%f MinValue: %f\n", value, Wage, minValue)
		highestValue = value
		highestValueWage = Wage
		highestValueTeam = team

		temp := create(team)
		temp.drawTeam()

		for _, player := range team {
			fmt.Printf("%+v\n", player)
		}
	}
}

// Pick a random fantasy football lineup for a sepcific position (e.g. Defender)
func (l *League) randPosLineup(pos string, remove []Player, minValue float64) []Player {
	out := []Player{}
	slice := l.FinalPlayers[pos]

	if len(slice) == 0 {
		return slice
	}

	removed := 0

	// TODO: this is slow as we loop over all players, fix
	temp := []Player{}

	// If utility Players, remove those already in the lineup
	if len(remove) > 0 {
		for _, v := range slice {
			pos := playerInLineup(&v, remove)
			if pos == -1 && v.Value > minValue {
				temp = append(temp, v)
				removed++
			}
		}
		slice = temp
	}

	/*if removed > 0 {
		fmt.Printf("Removed %d players\n", removed)
	}*/

	already := map[int]bool{}
	for i := 0; i < formation[pos]; {

		rand := random(0, len(slice)-1)

		if slice[rand].Value < minValue {
			continue
		}

		_, exist := already[rand]

		if !exist {
			slice[rand].Position = pos
			out = append(out, slice[rand])
			already[rand] = true
			i++
		}
	}

	return out
}

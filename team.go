package main

import (
	"fmt"
	"strings"
)

// Team
type Team struct {
	Name    string
	Players map[string]map[string]Player // Pos -> Name -> Player
}

func create(players []Player) Team {
	out := Team{}
	out.Players = map[string]map[string]Player{}
	for _, v := range players {
		if _, exist := out.Players[v.Position]; !exist {
			out.Players[v.Position] = map[string]Player{}
		}
		out.Players[v.Position][v.Name] = v
	}
	return out
}

func sufficientTeams(players []Player) bool {
	teams := 0
	got := map[string]bool{}
	for _, v := range players {
		if _, exist := got[v.Team]; !exist {
			got[v.Team] = true
			teams++
		}

		if teams >= minNumTeams {
			fmt.Println("Success: number of different teams is more than 3")
			return true
		}
	}
	return false

}

func (t *Team) print() {
	fmt.Printf("\n================================================================================\n======================================" + strings.ToUpper(t.Name) + "=======================================\n================================================================================\n")
	for pos, _ := range formation {
		fmt.Printf("\nPOSITION ("+strings.ToUpper(pos)+") - %d players\n", len(t.Players[pos]))
		for _, v := range t.Players[pos] {

			fmt.Printf(v.print() + ", ")
		}
		fmt.Println("")
	}
}

package main

import (
	"errors"
	"fmt"
)

// League
type League struct {
	Name  string
	Teams map[string]Team

	// Map between pos -> Player
	FinalPlayers map[string][]Player
}

// Add a player to the correct team in the league, also add it to the list of all players in the league
func (l *League) addPlayer(p *Player) {

	// Add to team
	temp, ex := l.Teams[p.Team]
	temp.Name = p.Team
	if !ex {
		temp.Players = map[string]map[string]Player{}
	}

	_, ex2 := l.Teams[p.Team].Players[p.Position]
	if !ex2 {
		temp.Players[p.Position] = map[string]Player{}
	}
	temp.Players[p.Position][p.Name] = *p

	// Utility player
	if _, ok := temp.Players["u"]; !ok {
		temp.Players["u"] = map[string]Player{}
	}

	if p.Position != "gk" {
		temp.Players["u"][p.Name] = *p
	}
	l.Teams[p.Team] = temp

}

// Get a player from the league
func (l *League) getPlayer(p Player) (Player, error) {
	curP := Player{}
	temp, ex := l.Teams[p.Team]
	temp.Name = p.Team
	if !ex {
		return curP, errors.New("Cant find the team " + p.Team + " - " + l.Name)
	}

	curP, exists := temp.Players[p.Position][p.Name]

	// Try to look in other positions
	if !exists {

		for pos, _ := range formation {
			otherExist := false
			curP, otherExist = temp.Players[pos][p.Name]
			if otherExist {
				exists = true
				break
			}
		}

		// Still doesnt exist? Give up
		if !exists {
			fmt.Printf("Can't find player: %s %s\n", p.Team, p.Name)
			return curP, errors.New(fmt.Sprintf("Can't find player: %s %s", p.Team, p.Name) + " - " + l.Name)
		}
	}

	return curP, nil

}

// Number of players
func (l *League) getSize() int {
	i := 0
	for _, v := range l.Teams {
		i += len(v.Players["u"])

	}
	return i
}

// Return (Exists, Boost Value)
func (l *League) isPlayerPlaying(p Player) (bool, bool) {
	if _, exist := l.Teams[p.Team]; !exist || len(l.Teams) == 0 {
		//	fmt.Println("(" + p.Team + ") The league is empty, so keeping all players")
		return false, false
	}

	_, err := l.getPlayer(p)
	if err != nil {
		return false, false
	} else {
		return true, true
	}
}

func (l *League) print() {
	fmt.Printf("\n======> League: %s (%d teams) \n", l.Name, len(l.Teams))
	for _, v := range l.Teams {
		v.print()
	}

}

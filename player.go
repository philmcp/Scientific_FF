package main

import (
	"fmt"
	"strings"
)

// Player
type Player struct {
	Name string
	Wage float64
	// g, d, m, f
	Position  string
	Team      string
	Value     float64
	IsPlaying string // Playing, Predicted to play, Not playing, Don't know
}

// Is the player already in the lineup?
func playerInLineup(p2 *Player, lineup []Player) int {
	for k, v := range lineup {
		if p2.equals(&v) {
			return k
		}
	}
	return -1
}

// Equality check for 2 players from 2 seperate data sources
func (p1 *Player) equals(p2 *Player) bool {
	out := p1.Name == p2.Name && p1.Team == p2.Team //&& p1.Position == p2.Position
	return out
}

func (p *Player) getDisplayName() string {
	return strings.ToUpper(p.Name[:1]) + p.Name[1:]
}

func (p *Player) print() string {
	return fmt.Sprintf("%s (%s) $%.0f Value:%.2f ", strings.ToUpper(p.Name[:1])+p.Name[1:], strings.ToUpper(p.Team), p.Wage, p.Value)

}

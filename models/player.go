package models

import (
	"fmt"
	"strings"
)

type Player struct {
	Name                string
	Team                string
	Position            string
	Status              string
	ReturningFromInjury bool
	Wage                float64
	Cost                float64
	Selected            float64
	Projection          float64
	Form                float64
	Points              float64 // Is this reliable? Aguero could have been out for half the season, come back and play sunderland at home
	ValueForm           float64 // What if the player was injured the last 3 games?
	TransfersInEvent    float64
	TransfersOutEvent   float64
	TransferRatio       float64
}

func (p *Player) GetDisplayName() string {
	return strings.ToUpper(p.Name[:1]) + p.Name[1:]
}

func (p *Player) getFullName() string {
	return p.Name + "_" + p.Team
}

func (p *Player) print() {
	fmt.Printf("%s (%s - %s) \tWage: %.0f \tProjection: %.2f \t Points: %.2f \t Selected: %.2f \t Form: %.2f \t ValueForm: %.2f \t TransferRatio: %.2f \t Score: %.2f\n", p.Name, p.Team, p.Status[:3], p.Wage, p.Projection, p.Points, p.Selected, p.Form, p.ValueForm, p.TransferRatio, p.getScore())
}

func (p *Player) getScore() float64 {
	//return p.Projection * p.Form * p.Points * p.Selected * p.ValueForm
	//	fmt.Println(maxTransferRatio)

	feat := p.Selected * p.Projection // .Form * p.Points * p. p.TransferRatio

	return feat

}

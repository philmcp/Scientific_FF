package models

import (
	"github.com/philmcp/Scientific_FF/utils"
	"log"
)

type Roto struct {
	Status              string
	ReturningFromInjury bool
}

type FPL struct {
	NowCost                 float64
	ValueForm               float64
	ValueSeason             float64
	CostChangeStart         float64
	CostChangeEvent         float64
	CostChangeStartFall     float64
	CostChangeEventFall     float64
	SelectedByPercent       float64
	Form                    float64
	TransfersOut            float64
	TransfersIn             float64
	TransfersOutEvent       float64
	TransfersInEvent        float64
	TotalPoints             float64
	EventPoints             float64
	PointsPerGame           float64
	Minutes                 float64
	GoalsScored             float64
	Assists                 float64
	CleanSheets             float64
	GoalsConceded           float64
	OwnGoals                float64
	PenaltiesSaved          float64
	PenaltiesMissed         float64
	YellowCards             float64
	RedCards                float64
	Saves                   float64
	Bonus                   float64
	Bps                     float64
	Influence               float64
	Creativity              float64
	Threat                  float64
	IctIndex                float64
	EaIndex                 float64
	TeamStrength            float64
	GameTeamStrengthOverall float64
	GameTeamStrengthAttack  float64
	GameTeamStrengthDefence float64
	GameOppStrengthOverall  float64
	GameOppStrengthAttack   float64
	GameOppStrengthDefence  float64
}

type Player struct {
	Name       string
	Team       string
	Position   string
	Cost       float64
	Wage       float64
	Projection float64

	Roto Roto
	FPL  FPL
}

func (p *Player) GetDisplayName() string {
	return utils.GetDisplayName(p.Name)
}

func (p *Player) getFullName() string {
	return p.Name + "_" + p.Team
}

func (p *Player) print() {
	log.Printf("%s (%s - %s) \tWage: %.0f \tProjection: %.2f \t Points: %.2f \t Selected: %.2f \t Form: %.2f \t ValueForm: %.2f \t Score: %.2f\n", p.Name, p.Team, p.Roto.Status[:3], p.Wage, p.Projection, p.FPL.TotalPoints, p.FPL.SelectedByPercent, p.FPL.Form, p.FPL.ValueForm, p.getScore())
}

func (p *Player) getScore() float64 {
	feat := (p.Projection * p.FPL.Form) / (1 + p.FPL.SelectedByPercent)
	//feat := p.FPL.SelectedByPercent
	return feat
}

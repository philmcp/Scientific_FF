package models

import (
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"strings"
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
	Name             string
	Team             string
	OppTeam          string
	Position         string
	PositionRaw      string
	Cost             float64
	Wage             float64
	Projection       float64
	IsHome           bool
	AvgPointsPerGame float64 // DK

	Roto Roto
	FPL  FPL
}

func (p *Player) GetDisplayName() string {
	return utils.GetDisplayName(p.Name)
}

func (p *Player) getFullName() string {
	return p.Name + "_" + p.Team
}

func (p *Player) IsGoodUtility() bool {
	return p.PositionRaw == "f" || p.PositionRaw == "m/f"
}

func (p *Player) GetGame() string {

	if p.IsHome {
		return strings.ToUpper(p.Team) + "* v " + strings.ToUpper(p.OppTeam)
	} else {
		return strings.ToUpper(p.OppTeam) + " v " + strings.ToUpper(p.Team) + "*"
	}
}

func (p *Player) print() {

	log.Printf("(%s %s) %s %.0f - %s\tProjection: %.2f\tPoints: %.2f\tSelected: %.2f\tPPG: %.2f\tAvgPPG: %.2f\tScore: %.2f\n", p.PositionRaw, p.Roto.Status[:3], p.Name, p.Wage, p.GetGame(), p.Projection, p.FPL.TotalPoints, p.FPL.SelectedByPercent, p.FPL.PointsPerGame, p.AvgPointsPerGame, p.GetScore())
}

func (p *Player) GetScore() float64 {
	feat := (0.9 * p.Projection) // + (p.AvgPointsPerGame * 0.15) // p.Projection //* p.FPL.PointsPerGame //) / (1 + (p.FPL.SelectedByPercent * 0.25))
	//feat := p.FPL.SelectedByPercent
	return feat
}

func teamDK2Twitter(name string) string {
	teamMapping := map[string]string{
		"stk": "#SCFC",
		"sot": "#SaintsFC",
		"whu": "#WHUFC",
		"mu":  "#MUFC",
		"cry": "#CPFC",
		"bou": "#AFCB",
		"ars": "#Arsenal",
		"lei": "#LCFC",
		"wba": "#WBA",
		"mci": "#MCFC",
		"swa": "#Swans",
		"hul": "#HCFC",
		"che": "#CFC",
		"tot": "#COYS",
		"eve": "#EFC",
		"liv": "#LFC",
		"sun": "#SAFC",
		"mid": "#BORO",
		"wat": "#WatfordFC",
		"bur": "#Burnley",
	}

	if _, exist := teamMapping[name]; !exist {
		log.Fatal("DK team name doesnt exist: " + name)
	}

	return teamMapping[name]
}

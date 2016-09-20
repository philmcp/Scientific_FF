package main

import (
	"fmt"
)

type CSVData struct {
	Data [][]string
}

type Configuration struct {
	Database struct {
		Host     string
		Port     int
		DB       string
		User     string
		Password string
	}
	FFPassword string
	Buffer     struct {
		AccessToken string
		TwitterID   string
	}
}

type Player struct {
	Name string
	Team string
	//	Pos                 string
	Status              string
	ReturningFromInjury bool

	Wage float64

	Projection float64
}

func (p *Player) getFullName() string {
	return p.Name + "_" + p.Team
}

func (p *Player) print() {
	fmt.Printf("%s (Wage: %f, Projection: %f)\n", p.Name, p.Wage, p.Projection)
}

type BestLineup struct {
	Team       PlayerPool
	Projection float64
	Wage       float64
}

// A full lineup or entire player pool
type PlayerPool map[string]PlayerList

func (p PlayerPool) print() {
	fmt.Println("======= Start player pool ======= \n")
	for _, players := range p {
		for _, player := range players {
			player.print()
		}
	}
	fmt.Println("======= End player pool ======= ")

}

// A list of players
type PlayerList []Player

package models

import (
	"fmt"
)

// Player pool
type PlayerPool map[string]PlayerList
type PlayerList []Player

func (p PlayerPool) print() {
	fmt.Println("======= Start player pool ======= \n")
	for _, players := range p {
		for _, player := range players {
			player.print()
		}
	}
	fmt.Println("======= End player pool ======= ")

}

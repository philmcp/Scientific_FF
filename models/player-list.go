package models

import (
	//	"fmt"
	"github.com/philmcp/Scientific_FF/utils"
)

type PlayerList []Player

// Pick a random fantasy football lineup for a sepcific position (e.g. Defender)
func (players PlayerList) RandomPosition(ignore map[string]bool, minValue float64, num int, isUtility bool) PlayerList {
	out := PlayerList{}
	for {
		rand := players[utils.Random(0, len(players))]

		//	fmt.Println(rand.OppTeam)
		if minValue > rand.GetScore() {
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

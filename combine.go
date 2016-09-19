package main

/*
import (
	"fmt"
	//	"os"
)

// Given a fantasy league and a stats league (i.e. FFS)
// &leagueFFS, &leaguePL, &leagueFFSP
func (fantasy *League) combine(stats *League, playing *League, predicted *League) {
	fantasy.FinalPlayers = map[string][]Player{}

	weights := fantasy.removeNonLineup(playing, predicted)

	for v, _ := range formation {
		fantasy.FinalPlayers[v] = []Player{}
	}

	// Go over the second league's players and find the closest match
	fmt.Println("Combining")
	for k1, team := range fantasy.Teams {

		for pos, _ := range formation {
			for k2, player := range team.Players[pos] {

				cur, err := stats.getPlayer(player)

				if err == nil {
					temp := player
					temp.Value = cur.Value

					if weights[team.Name][pos][player.Name] == "Yes" {
						temp.IsPlaying = "Yes"
						temp.Value *= 1.25

					} else if weights[team.Name][pos][player.Name] == "No" {
						temp.IsPlaying = "No"
						temp.Value *= 0.0
					} else if weights[team.Name][pos][player.Name] == "Predicted" {
						temp.IsPlaying = "Predicted"
						temp.Value *= 1.0
					} else {
						temp.IsPlaying = "Unknown"
						temp.Value *= 0.0
					}

					fantasy.Teams[k1].Players[pos][k2] = temp
					//		stats.Teams[k1].Players[pos][k2] = temp

					fantasy.FinalPlayers[pos] = append(fantasy.FinalPlayers[pos], temp)

				}
			}

		}

	}
	fmt.Println("End Combining")

}

func (fantasy *League) removeNonLineup(lineups *League, predicted *League) map[string]map[string]map[string]string {
	output := map[string]map[string]map[string]string{}
	fmt.Println("Finding players that are playing...")

	for _, team := range fantasy.Teams {
		//	fmt.Printf("%+v\n", team)
		// Copy as a temp
		curPlayers := team.Players
		output[team.Name] = map[string]map[string]string{}
		// Positions
		for pos, players := range curPlayers {

			output[team.Name][pos] = map[string]string{}
			for name, player := range players {

				playing, boost := lineups.isPlayerPlaying(player)
				predicted, _ := predicted.isPlayerPlaying(player)

				//fmt.Printf("Player: "+player.Name+" is Predicted: %+v Playing: %+v Boost: %+v\n", predicted, playing, boost)

				if playing && boost {
					output[team.Name][pos][name] = "Yes"
				} else if playing && !boost {
					output[team.Name][pos][name] = "Unknown"
				} else if !playing && predicted {
					output[team.Name][pos][name] = "Predicted"
				} else {
					delete(team.Players[pos], name)
					output[team.Name][pos][name] = "No"
				}

			}

		}
	}
	return output

}
*/

package main

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"golang.org/x/net/context"
	"log"
	"time"
)

var (
	best           = models.Generated{Projection: 0}
	maxPlayerScore = 0.0
	lineupText     = []string{"Its #GW%d time! Here is our #EPL #Draftkings lineup for %s",
		"This week's #EPL #Draftkings lineup! #GW%d %s - Good luck all!",
		"Here's our lineup for today's #GW%d %s #Draftkings contest",
		"Algorithm driven #EPL fantasy football lineups - here's today's #GW%d %s selection",
		"...and here it is, good luck all! #EPL #DraftKings #GW%d %s",
		"Here is our #AI selected #GW%d %s #DraftKings lineup - Good luck!",
		"Our algorithm driven #EPL #DraftKings lineups for #GW%d %s is here! Good luck!"}
)

/* Generate a new optimal linup */
func generate() {
	log.Println("Generating a new optimal line up")
	log.Printf("Running for Week: %d Season: %d Game: %d\n", config.Week, config.Season, config.DKID)

	loadAllData()

	pool, maxScore := db.GetPlayers()
	maxPlayerScore = maxScore

	//	pool.Print()

	getBestLineup(&pool)
	drawer.DrawTeam(&best)

	msg := fmt.Sprintf("Our EPL #GW%d %s @DraftKings lineup: %s", config.Week, config.DKName, best.Team.PrintTeamCounts())
	buffer.Post(msg, config.RemoteLoc+"lineup.png")
}

func getBestLineup(pool *models.PlayerPool) {
	log.Println("============ Started generating lineups =============")
	//	Spawn some threads to select random team combinations

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Duration+5))

	for k := 3.0; k <= (config.Threads); k++ {
		go thread(ctx, int(k), pool, maxPlayerScore/k)
	}

	select {
	case <-ctx.Done():
		return
	case <-time.After(time.Second * time.Duration(config.Duration)):
		//Here I'm actually ending it earlier than the timeout with cancel().
		cancel()
	}

	log.Println("============ Finished generating lineups =============")
}

// Thread to select random team combinations
func thread(ctx context.Context, num int, pool *models.PlayerPool, minValue float64) {
	printGap := 1000
	log.Printf("Searching for lineups with a min value of %f\n", minValue)
	i := 0
	for {

		select {
		case <-ctx.Done():
			log.Println("Done ", num, "working")
			return
		default:
			cur := pool.RandomLineup(minValue, config.Formation)

			if cur.Wage <= config.MaxWage && cur.Projection > best.Projection && cur.NumTeams >= config.MinNumTeams && cur.Team.IsBalanced() {
				log.Printf("\n\n***** New high score: %f Wage: $%f MinValue: %f Teams: %d Iter: %d\n", cur.Projection, cur.Wage, minValue, cur.NumTeams, iter)
				best.Wage = cur.Wage
				best.Projection = cur.Projection
				best.Team = cur.Team
				best.Team.Print()
			}

			if i%printGap == 0 {
				iter += printGap
				if iter%(printGap*int(config.Threads)) == 0 {

					speed := float64(iter) / float64(time.Now().Sub(start).Seconds())
					log.Printf("Thread: %d Iter: %d Speed: %f per second\n", num, iter, speed)
				}

			}
			i++

		}

	}
}

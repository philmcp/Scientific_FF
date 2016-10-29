package main

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"log"
	"strings"
	"time"
)

func lineups() {
	day := time.Now().Weekday().String()
	if day == "Saturday" || day == "Sunday" || day == "Monday" {
		tweets := twitter.GetLineupNews()

		log.Println("============ Getting lineup news! ============ ")

		for _, tweet := range tweets {

			hasBeenPosted, _ := db.Conn.Query("SELECT id FROM tweets WHERE id = $1", tweet.ID)
			if !hasBeenPosted.Next() {
				lineup := models.NewLineup(tweet.Text)
				db.InsertTweet(&tweet)
				log.Printf("New lineup: %+v\n", lineup)
				drawer.DrawLineup(&lineup, tweet.ID)

				if lineup.OppTeam != "" {
					lineup.OppTeam = " vs " + lineup.OppTeam
				}

				post := fmt.Sprintf(lineup.Team + lineup.OppTeam + ": ")
				for _, player := range lineup.Players {
					post += player + ", "
				}
				post = strings.Trim(post, ", ") + " #FPL #EPL"

				img := fmt.Sprintf(config.RemoteLoc+"lineups/%d.png", tweet.ID)
				buffer.Post(post, img)
			} else {
				log.Printf("%+v already posted\n", tweet)
			}

		}
	}

}

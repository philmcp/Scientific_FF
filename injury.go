package main

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"log"

	"strings"
)

func injuries() {
	tweets := twitter.GetInjuryNews()

	log.Println("============ Getting injury news!! ============ ")

	for _, tweet := range tweets {
		if !strings.Contains(tweet.Text, "#FPL") {
			continue
		}

		check := fmt.Sprintf(config.OutputFolder+"injuries/%d.png", tweet.ID)
		log.Println("Checking to see if exists: " + check)
		if !utils.FileExists(check) {

			log.Printf("Tweet %d hasnt been seen before, posting injury\n", tweet.ID)
			inj := models.NewInjury(tweet.Text)

			// Has this player played this season?
			name := utils.GetLastName(inj.Name)
			isWorthyPlayer, _ := db.Conn.Query("SELECT name FROM dk WHERE season = $1 AND name = $2", config.Season, name)
			defer isWorthyPlayer.Close()
			if isWorthyPlayer.Next() {
				log.Println(name + " IS worthy...")
				drawer.DrawInjury(&inj, tweet.ID)
				buffer.Post(fmt.Sprintf("#INJURY - %s (#%s) - %s - %s #FPL #EPL", inj.Name, strings.ToUpper(inj.Team), inj.Injury, inj.Returns), fmt.Sprintf("%sinjuries/%d.png", config.RemoteLoc, tweet.ID))
			} else {
				log.Println(name + " IS NOT worthy...")
			}
		}

	}
}

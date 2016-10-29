package main

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"time"
)

var (
	lastPosted = ""
	random1    = []string{"The best fantasy football lineups are those selected by algorithms. When it comes to picking your @DraftKings lineups - leave it to the machine.",
		"We generate 1M different #EPL lineups to select the team with the highest projected fantasy football score.",
		"Follow us for algorithm driven #EPL fantasy football lineups.",
		"Using artificial intelligence to pick the 'optimal' weekly fantasy football lineup. Stop losing, start winning.",
		"You may not use our line-ups, but you should at least be aware of which lineup 'the computer' would pick #AI"}

	random2 = []string{"Follow us for team news, injuries and algorithmic driven @DraftKings linueups.",
		"Which sites do we support? Currently: @dkuk #EPL Soon: @FanduelUK",
		"#FPL There's too much data to humanly process, are you still doing this manually? Leave it to the machine.",
		"Our algorithms run every #EPL @dkuk game week - crunching numbers half an hour before games and posted 15 minutes before kick-off.",
		"Forgot to swap out an injured #FPL player? Follow us for up the minute player injury news"}

	random3 = []string{"We consider significant #EPL data when training our algorithms e.g. goals scored, #FPL player 'selected by' % etc",
		"Follow @Scientific_FF for team news, player injuries and data driven @DraftKings line-ups.",
	}
)

func fanduel(day string) {
	fmt.Println(day + " " + time.Now().Weekday().String())
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING fanduel ======= ")
	buffer.Post("Fancy having your initial deposit DOUBLED (up to £400) at FanDuel? 💰\n\nSign up NOW ➡️ http://bit.ly/2eXZkfj", "https://www.scoopanalytics.com/ff/images/fanduel.png")
}

func ladbrokes(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING ladbrokes ======= ")
	buffer.Post("Fancy a FREE £50 bet at Ladbrokes? 💰\n\nGet your #freebet NOW ➡️ http://bit.ly/2ewkUYr", "https://www.scoopanalytics.com/ff/images/ladbrokes.png")
}

func skybet(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING skybet ======= ")
	buffer.Post("Want a FREE £20 bet at SkyBet? 💰\n\nGet your #freebet NOW ➡️ http://bit.ly/2eXWN4R", "https://www.scoopanalytics.com/ff/images/skybet.png")
}

func betfair(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING betfair ======= ")
	buffer.Post("Fancy a FREE £30 bet at Betfair? 💰\n\nGet your #freebet NOW ➡️ http://bit.ly/2f74mFH", "https://www.scoopanalytics.com/ff/images/betfair.png")
}

func netbet(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING netbet ======= ")
	buffer.Post("Fancy a FREE $50 bet at NetBet? 💰\n\nJoin with promo code 'WELCOME50' NOW ➡️ http://bit.ly/2ewjbCs", "https://www.scoopanalytics.com/ff/images/netbet.png")
}

func randomPost1(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	id := utils.Random(0, len(random1)-1)
	log.Println("======= POSTING RANDOM1 ======= ")
	buffer.Post(random1[id], "")
}

func randomPost2(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	id := utils.Random(0, len(random2)-1)
	log.Println("======= POSTING RANDOM2 ======= ")
	buffer.Post(random2[id], "")
}

func gameday(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING gameday ======= ")

	buffer.Post(fmt.Sprintf("#FPL It's #GW%d - Our #AI driven #DraftKings lineups will be posted later today, check back soon...", config.Week), "https://www.scoopanalytics.com/ff/images/gameday.jpg")

}

func crunching(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	log.Println("======= POSTING crunching ======= ")
	buffer.Post(fmt.Sprintf("#FPL It's almost time for #GW%d - Our #DraftKings lineups will be posted 15 minutes before kick-off", config.Week), "https://www.scoopanalytics.com/ff/images/crunching.jpg")

}

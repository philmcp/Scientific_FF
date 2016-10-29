package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/philmcp/Scientific_FF/api"
	"github.com/philmcp/Scientific_FF/conf"
	"github.com/philmcp/Scientific_FF/draw"
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/orm"
	"github.com/philmcp/Scientific_FF/scrape"
	"log"
	"os"
	"runtime"
	"time"
)

// Config
var (
	config  = &models.Configuration{}
	db      *orm.ORM
	buffer  *api.BufferAPI
	twitter *api.TwitterAPI
	scraper *scrape.Scrape
	drawer  *draw.Draw

	inputFolder  string
	outputFolder string

	// Thread stuff
	iter  = 0
	start = time.Now()
)

// TODO:
// 			Set up cron jobs so that it posted the tweets 1hr, 2hr, 30mins before etc
// Sort injury images so they appear better ("returns in hours also"")
// Inlcude player team counts in buffer
// NOTE WEEK 16 - games are midweek
func main() {
	runtime.GOMAXPROCS(12)

	confFile := "assets/conf/local.json"
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}

	// Load config
	if err := conf.LoadConfig(confFile, config); err != nil {
		log.Fatal(err)
	}

	db = orm.NewORM(config)
	twitter = api.NewTwitter(config)
	buffer = api.NewBuffer(config)
	scraper = scrape.NewScraper(config)
	drawer = draw.NewDraw(config)
	//injuries()
	//	crunching("Saturday")
	//gocron.Every(1).Day().At("14:31").Do(generate)
	generate()
	//	generate()

	// Random
	gocron.Every(1).Day().At("18:18").Do(randomPost1, "Tuesday")
	gocron.Every(1).Day().At("18:18").Do(randomPost2, "Tuesday")

	// Analytics
	gocron.Every(1).Day().At("18:18").Do(analyticsInForm, "Wednesday")
	gocron.Every(1).Day().At("18:18").Do(analyticsTransfersOut, "Thursday")
	gocron.Every(1).Day().At("18:18").Do(analyticsTransfersIn, "Friday")

	// Affiliates
	gocron.Every(1).Day().At("14:13").Do(skybet, "Monday")
	gocron.Every(1).Day().At("14:13").Do(fanduel, "Tuesday")
	gocron.Every(1).Day().At("14:13").Do(ladbrokes, "Wednesday")
	gocron.Every(1).Day().At("14:13").Do(betfair, "Thursday")
	gocron.Every(1).Day().At("14:13").Do(netbet, "Friday")

	// Gameday
	gocron.Every(1).Day().At("10:15").Do(gameday, "Saturday")
	//gocron.Every(1).Day().At("10:15").Do(gameday, "Sunday")

	// Crunching
	gocron.Every(1).Day().At("14:00").Do(crunching, "Saturday")
	//	gocron.Every(1).Day().At("14:00").Do(crunching, "Saturday")

	//	gocron.Every(1).Day().At("11:45").Do(crunching, "Sunday")

	// Lineups
	gocron.Every(25).Minutes().Do(lineups)

	// Injuries
	gocron.Every(30).Minutes().Do(injuries)

	<-gocron.Start()
}

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
	start = time.Now().UnixNano()
)

// TODO:
// 			Create DK upload scripts
// 			Set up cron jobs so that it posted the tweets 1hr, 2hr, 30mins before etc

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

	generate()
	//	injuries()
	//	conf.Twitter.getInjuryNews()

	gocron.Every(30).Minutes().Do(injuries)
	// function Start start all the pending jobs
	<-gocron.Start()
}

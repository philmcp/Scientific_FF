package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/philmcp/Scientific_FF/models"
	"log"
	"os"
	"runtime"
	"time"
)

// Config
var (
	conf = &Configuration{}
	db   *DB

	inputFolder  string
	outputFolder string

	bestLineup = BestLineup{Wage: 0, Projection: 0}

	// Thread stuff
	iter  = 0
	start = time.Now().UnixNano()
)

// TODO:
// 			Create DK upload scripts
// 			Set up cron jobs so that it posted the tweets 1hr, 2hr, 30mins before etc

func main() {
	runtime.GOMAXPROCS(12)
	if len(os.Args) < 2 {
		log.Fatal("Missing arguements...")
	}

	confFile := "conf/local.json"
	if len(os.Args) > 1 {
		confFile = os.Args[1]
	}

	// Load config
	if err := LoadConfig(confFile, conf); err != nil {
		log.Fatal(err)
	}

	db = Connect(conf.Database.Host, conf.Database.Port, conf.Database.DB, conf.Database.User, conf.Database.Password)

	//	generate()

	conf.Twitter.getInjuryNews()

	//gocron.Every(5).Seconds().Do(conf.Twitter.getInjuryNews)
	// function Start start all the pending jobs
	<-gocron.Start()
}

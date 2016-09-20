package main

import (
	"fmt"
	//	"os"
	"runtime"
	"time"
)

// Fantasy football Team Selection - Parameters
var (
	SEASON = 2016
	WEEK   = 5
	DKID   = 10920
	DKNAME = "18/9 12:00pm"

	conf = &Configuration{}

	db *DB

	inputFolder  = fmt.Sprintf("input/%d/%d/", SEASON, WEEK)
	outputFolder = fmt.Sprintf("output/%d/%d/", SEASON, WEEK)

	formation = map[string]int{"gk": 1, "d": 2, "m": 2, "f": 2, "u": 1}

	BEST = BestLineup{Wage: 0, Projection: 0}

	iter          = 0
	start         = time.Now().UnixNano()
	printGap      = 100000
	threads       = 10.0
	valueJump     = 0.25
	minValueStart = 0.75

	minNumTeams = 3

	MAX_WAGE = 50000.0

	// 'Source' -> 'Team' -> 'From' -> 'To'
	rename = map[string]map[string]map[string]string{}
)

// TODO:
// 			Create DK upload scripts
// 			Set up cron jobs so that it posted the tweets 1hr, 2hr, 30mins before etc
// 			Show unique players who r missing (shorter)
// 			Upload code to github

func main() {
	fmt.Printf("Running for Week: %d Season: %d Game: %d\n", WEEK, SEASON, DKID)
	time.Sleep(time.Second * 1)
	runtime.GOMAXPROCS(12)

	// Load config
	if err := LoadConfig("conf.json", conf); err != nil {
		fmt.Println(err)
	}

	db = Connect(conf.Database.Host, conf.Database.Port, conf.Database.DB, conf.Database.User, conf.Database.Password)

	pool := loadData()
	pool.getBestLineup()
	//postToBuffer()

}

// Scrape the data and loads it into the database
func loadData() PlayerPool {
	scrapeFFS()
	scrapeDK()
	scrapeRoto()

	// FFS
	fileFFS := inputFolder + "ffs-" + fmt.Sprintf("%d", DKID) + ".csv"
	csvFFS := loadCSV(fileFFS)
	db.loadFFS(&csvFFS)

	// DK
	fileDK := inputFolder + "dk-" + fmt.Sprintf("%d", DKID) + ".csv"
	csvDK := loadCSV(fileDK)
	db.loadDK(&csvDK)

	// Roto
	fileRotoP := inputFolder + "roto-players.csv"
	csvRotoP := loadCSV(fileRotoP)
	db.loadRoto(&csvRotoP)

	return db.getPlayers()
}

func (pool *PlayerPool) getBestLineup() {

	//	Spawn some threads to select random team combinations
	for k := 0.0; k < (threads * 2.0); k++ {
		go thread(pool, minValueStart+(k*valueJump))
	}

	// Run for 10 mins
	time.Sleep(90 * time.Second)
}

// Thread to select random team combinations
func thread(pool *PlayerPool, minValue float64) {
	fmt.Printf("Searching for lineups with a min value of %f\n", minValue)
	i := 0
	for {
		pool.randTeamLineup(minValue)
		i++
		if i%printGap == 0 {
			iter += i
			i = 0
			if iter%(printGap*int(threads)) == 0 {
				now := time.Now().UnixNano()
				elapsed := float64(now-start) / 1000000000.0
				speed := float64(iter) / elapsed
				fmt.Printf("Thread: %d after %fs - %f/s\n", iter, elapsed, speed)
			}

		}
	}
}

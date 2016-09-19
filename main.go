package main

import (
	"fmt"
	//	"os"
	"runtime"
	"time"
)

// Fantasy football Team Selection - Parameters
var (
	season   = 2016
	gameWeek = 6
	DKID     = 10921
	DKName   = "18/9 12:00pm"

	conf = &Configuration{}

	db *DB

	inputFolder  = fmt.Sprintf("input/%d/%d/", season, gameWeek)
	outputFolder = fmt.Sprintf("output/%d/%d/", season, gameWeek)

	league    = League{Teams: map[string]Team{}}
	formation = map[string]int{"gk": 1, "d": 2, "m": 2, "f": 2, "u": 1}

	highestValue     = 0.0
	highestValueWage = 0.0
	highestValueTeam = []Player{}

	iter          = 0
	start         = time.Now().UnixNano()
	printGap      = 100000
	threads       = 10.0
	valueJump     = 0.25
	minValueStart = 0.75

	minNumTeams = 3

	maxWage = 50000.0

	// 'Source' -> 'Team' -> 'From' -> 'To'
	rename = map[string]map[string]map[string]string{}
)

// TODO:
// 			Create DK upload scripts
// 			Set up cron jobs so that it posted the tweets 1hr, 2hr, 30mins before etc
// 			Show unique players who r missing (shorter)
// 			Upload code to github

func main() {
	fmt.Printf("Running for Week: %d Season: %d Game: %d\n", gameWeek, season, DKID)
	time.Sleep(time.Second * 1)
	runtime.GOMAXPROCS(12)

	// Load config
	if err := LoadConfig("conf.json", conf); err != nil {
		fmt.Println(err)
	}

	db = Connect(conf.Database.Host, conf.Database.Port, conf.Database.DB, conf.Database.User, conf.Database.Password)

	leagueFFS, leagueDK, leaguePL, leagueFFSP := loadData()
	getBestLineup(leagueFFS, leagueDK, leaguePL, leagueFFSP)
	postToBuffer()

}

// Load all data
func loadData() (leagueFFS League, leagueDK League, leaguePL League, leagueFFSP League) {
	// Scrape the data first
	scrapeFFS()
	scrapeDK()
	scrapeRoto()

	// Fantasy Football Scout Data
	fileFFS := inputFolder + "ffs-" + fmt.Sprintf("%d", DKID) + ".csv"
	csvFFS := loadCSV(fileFFS)
	db.parseFFSDB(&csvFFS)
	leagueFFS = csvFFS.parseFFS()

	// Draft Kings Data
	fileDK := inputFolder + "dk-" + fmt.Sprintf("%d", DKID) + ".csv"
	csvDK := loadCSV(fileDK)
	db.parseDKDB(&csvDK)
	leagueDK = csvDK.parseDK()
	//	leagueDK.print()

	// Roto players
	fileRotoP := inputFolder + "roto-players.csv"
	fmt.Println(fileRotoP)
	csvRotoP := loadCSV(fileRotoP)
	db.parseRotoPlayersDB(&csvRotoP)

	return leagueFFS, leagueDK, leaguePL, leagueFFSP
}

func getBestLineup(leagueFFS League, leagueDK League, leaguePL League, leagueFFSP League) {

	leagueDK.combine(&leagueFFS, &leaguePL, &leagueFFSP)
	//	Spawn some threads to select random team combinations
	for k := 0.0; k < (threads * 2.0); k++ {
		go thread(&leagueDK, formation, minValueStart+(k*valueJump))
	}

	// Run for 10 mins
	time.Sleep(90 * time.Second)
}

// Thread to select random team combinations
func thread(l *League, formation map[string]int, minValue float64) {
	fmt.Printf("Searching for lineups with a min value of %f\n", minValue)
	i := 0
	for {
		l.randTeamLineup(formation, minValue)
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

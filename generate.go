package main

import (
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"os"
	"time"
)

var (
	best       = models.Lineup{Projection: 0}
	allThreads = make(chan bool)
)

/* Generate a new optimal linup */
func generate() {
	log.Println("Generating a new optimal line up")
	log.Printf("Running for Week: %d Season: %d Game: %d\n", config.Week, config.Season, config.DKID)

	scrapeData()
	data := parseData()
	loadData(data)

	pool := db.GetPlayers()
	//	pool.Print()

	getBestLineup(&pool)
	drawer.DrawTeam(&best)

	//buffer.Post()
}

func scrapeData() {
	log.Println("\n============= Scraping data =============")

	if !utils.FileExists(scraper.Folder) {
		log.Println(scraper.Folder + " doesnt exist, creating...")
		err := os.Mkdir(scraper.Folder, 0755)
		if err != nil {
			log.Println(err)
		}
	}

	scraper.ScrapeFFS()
	scraper.ScrapeRoto()
	scraper.ScrapeFPL()
	scraper.ScrapeDK()
}

func parseData() *models.Data {
	log.Println("\n============= Parsing data =============")
	out := models.Data{}

	out.Roto = scraper.ParseRoto()
	out.FFS = scraper.ParseFFS()
	out.FPL = scraper.ParseFPL()
	out.DK = scraper.ParseDK()

	return &out

}

func loadData(data *models.Data) {
	log.Println("\n============= Loading data =============")

	db.LoadFFS(data.FFS)
	db.LoadDK(data.DK)
	db.LoadFPL(data.FPL)

	db.LoadRoto(data.Roto)
}

func getBestLineup(pool *models.PlayerPool) {
	log.Println("============ Started generating lineups =============")
	//	Spawn some threads to select random team combinations
	for k := 0.0; k < (config.Threads); k++ {
		go thread(pool, config.MinValue+(k*config.ValueJump))
	}

	// Run for X seconds
	time.Sleep(100 * time.Second)
	log.Println("============ Finished generating lineups =============")
}

// Thread to select random team combinations
func thread(pool *models.PlayerPool, minValue float64) {
	printGap := 1000
	log.Printf("Searching for lineups with a min value of %f\n", minValue)

	for i := 0; i < 10000; i++ {
		cur := pool.RandomLineup(minValue, config.Formation)

		if cur.Wage <= config.MaxWage && cur.Projection > best.Projection && cur.NumTeams >= config.MinNumTeams {
			log.Printf("\n***** New high score: %f Wage: $%f MinValue: %f Teams: %d\n", cur.Projection, cur.Wage, minValue, cur.NumTeams)
			best.Wage = cur.Wage
			best.Projection = cur.Projection
			best.Team = cur.Team
			best.Team.Print()
		}

		i++
		if i%printGap == 0 {
			iter += i
			i = 0
			if iter%(printGap*int(config.Threads)) == 0 {
				now := time.Now().UnixNano()
				elapsed := float64(now-start) / 1000000000.0
				speed := float64(iter) / elapsed
				log.Printf("Thread: %d after %fs - %f/s\n", iter, elapsed, speed)
			}

		}
	}
}

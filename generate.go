package main

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"time"
)

var best = models.Lineup{Projection: 0}

/* Generate a new optimal linup */
func generate() {
	fmt.Println("Generating a new optimal line up")
	fmt.Printf("Running for Week: %d Season: %d Game: %d\n", config.Week, config.Season, config.DKID)

	//scrapeData()
	data := parseData()
	loadData(data)

	pool := db.GetPlayers()
	pool.Print()

	getBestLineup(&pool)
	drawer.DrawTeam(&best)

	//postToBuffer()
}

func scrapeData() {
	fmt.Println("\n============= Scraping data =============")

	scraper.ScrapeFFS()
	scraper.ScrapeRoto()
	scraper.ScrapeFPL("value_form")
	scraper.ScrapeFPL("transfers_in_event")
	scraper.ScrapeFPL("transfers_out_event")
	scraper.ScrapeDK()
}

func parseData() *models.Data {
	fmt.Println("\n============= Parsing data =============")
	out := models.Data{}

	out.Roto = scraper.ParseRoto()
	out.FFS = scraper.ParseFFS()
	out.FPL.ValueForm = scraper.ParseFPL("value_form")
	out.FPL.TransfersInEvent = scraper.ParseFPL("transfers_in_event")
	out.FPL.TransfersOutEvent = scraper.ParseFPL("transfers_out_event")
	out.DK = scraper.ParseDK()

	return &out

}

func loadData(data *models.Data) {
	fmt.Println("\n============= Loading data =============")

	db.LoadFFS(data.FFS)
	db.LoadDK(data.DK)

	db.LoadFPL(data.FPL.ValueForm, "value_form")
	db.LoadFPL(data.FPL.TransfersInEvent, "transfers_in_event")
	db.LoadFPL(data.FPL.TransfersOutEvent, "transfers_out_event")

	db.LoadRoto(data.Roto)
}

func getBestLineup(pool *models.PlayerPool) {
	fmt.Println("============ Started generating lineups =============")
	//	Spawn some threads to select random team combinations
	for k := 0.0; k < (config.Threads); k++ {
		go thread(pool, config.MinValue+(k*config.ValueJump))
	}

	// Run for X seconds
	time.Sleep(10 * time.Second)
	fmt.Println("============ Finished generating lineups =============")
}

// Thread to select random team combinations
func thread(pool *models.PlayerPool, minValue float64) {
	printGap := 1000
	fmt.Printf("Searching for lineups with a min value of %f\n", minValue)

	for i := 0; i < 10000; i++ {
		cur := pool.RandomLineup(minValue, config.Formation)

		if cur.Wage <= config.MaxWage && cur.Projection > best.Projection && cur.NumTeams >= config.MinNumTeams {
			fmt.Printf("\n***** New high score: %f Wage: $%f MinValue: %f Teams: %d\n", cur.Projection, cur.Wage, minValue, cur.NumTeams)
			best.Wage = cur.Wage
			best.Projection = cur.Projection
			best.Team = cur.Team

			//			team.drawTeam()

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
				fmt.Printf("Thread: %d after %fs - %f/s\n", iter, elapsed, speed)
			}

		}
	}
}

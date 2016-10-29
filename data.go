package main

import (
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"os"
)

func loadAllData() {
	scrapeData()
	data := parseData()
	loadData(data)
}

func loadFPLData() {
	scraper.ScrapeFPL()
	out := models.Data{}
	out.FPL = scraper.ParseFPL()
	db.LoadFPL(out.FPL)
}

func scrapeData() {
	log.Printf("\n============= Scraping data for GW%d =============\n", config.Week)

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

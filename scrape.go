package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Scrape the data and loads it into the database
func scrapeData() PlayerPool {
	fplTypes := []string{"value_form", "transfers_in_event", "transfers_out_event"}

	//for _, cur := range fplTypes {
	//	scrapeFPL(cur)
	//	}

	scrapeFFS()
	scrapeDK()
	scrapeRoto()

	// FPL

	for _, cur := range fplTypes {
		fileFPL := inputFolder + "fpl-" + cur + ".csv"
		csvFPL := loadCSV(fileFPL)
		db.loadFPL(&csvFPL, cur)

	}

	// FFS
	fileFFS := inputFolder + "ffs-" + fmt.Sprintf("%d", conf.DKID) + ".csv"
	csvFFS := loadCSV(fileFFS)
	db.loadFFS(&csvFFS)

	// DK
	fileDK := inputFolder + "dk-" + fmt.Sprintf("%d", conf.DKID) + ".csv"
	csvDK := loadCSV(fileDK)
	db.loadDK(&csvDK)

	// Roto
	fileRotoP := inputFolder + "roto-players.csv"
	csvRotoP := loadCSV(fileRotoP)
	db.loadRoto(&csvRotoP)

	return db.getPlayers()
}

func scrapeFFS() {
	fmt.Println("Scraping FFS data to " + inputFolder)
	output, err := exec.Command("php", "scrape/ffs.php", inputFolder, fmt.Sprintf("%d", conf.DKID), conf.FFS.Username, conf.FFS.Password).Output()

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(output), "error") {
		log.Fatal("FSS didn't crawl properly")
	}

	fmt.Printf("FFS scrape finished\n\n")
}

func scrapeRoto() {
	fmt.Println("Scraping Roto data to " + inputFolder)
	output, err := exec.Command("php", "scrape/roto.php", inputFolder, fmt.Sprintf("%d", conf.DKID)).Output()
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(output), "error") {
		log.Fatal("Roto didn't crawl properly")
	}

	fmt.Printf("Roto Scrape finished\n\n")
}

// page = transfers_out_event, transfers_out_event, value_form
func scrapeFPL(page string) {
	fmt.Println("Scraping FPL data to " + inputFolder)
	out, err := exec.Command("phantomjs", "scrape/fpl.js", page, inputFolder).Output()

	fmt.Println(string(out))

	if err != nil {
		log.Fatal(err)
	}
	// TODO: check these files actually contain data

	fmt.Printf("FPL Scrape finished\n\n")
}

func scrapeDK() {

	url := fmt.Sprintf("https://www.draftkings.co.uk/lineup/getavailableplayerscsv?contestTypeId=27&draftGroupId=%d", conf.DKID)
	out := inputFolder + fmt.Sprintf("dk-%d.csv", conf.DKID)
	fmt.Println("Scraping DK data to " + inputFolder)

	Wget(url, out)
	fmt.Printf("DK Scrape finished\n\n")
}

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func scrapeFFS() {
	fmt.Println("Scraping FFS data to " + inputFolder)
	_, err := exec.Command("php", "scrape/ffs.php", inputFolder, fmt.Sprintf("%d", DKID)).Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FFS scrape finished\n\n")
}

func scrapeRoto() {
	fmt.Println("Scraping Roto data to " + inputFolder)
	_, err := exec.Command("php", "scrape/roto.php", inputFolder, fmt.Sprintf("%d", DKID)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Roto Scrape finished\n\n")
}

func scrapeDK() {

	url := fmt.Sprintf("https://www.draftkings.co.uk/lineup/getavailableplayerscsv?contestTypeId=27&draftGroupId=%d", DKID)
	out := inputFolder + fmt.Sprintf("dk-%d.csv", DKID)
	fmt.Println("Scraping DK data to " + inputFolder)

	Wget(url, out)
	fmt.Printf("DK Scrape finished\n\n")
}

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func scrapeFFS() {
	fmt.Println("Scraping FFS data to " + inputFolder + " " + fmt.Sprintf("%d", DKID))
	out, err := exec.Command("php", "scrape/ffs.php", inputFolder, fmt.Sprintf("%d", DKID)).Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output from FFS Scrape: %s\n", out)
}

func scrapeRoto() {
	fmt.Println("Scraping Roto data to " + inputFolder)
	out, err := exec.Command("php", "scrape/roto.php", inputFolder, fmt.Sprintf("%d", DKID)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output from Roto Lineup Scrape: %s\n", out)
}

func scrapeDK() {
	url := fmt.Sprintf("https://www.draftkings.co.uk/lineup/getavailableplayerscsv?contestTypeId=27&draftGroupId=%d", DKID)
	out := inputFolder + fmt.Sprintf("dk-%d.csv", DKID)
	fmt.Println("Downloading from " + url + " to " + out)
	wget(url, out)
}

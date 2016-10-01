package scrape

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"os/exec"
	"strings"
)

/* Does the actual scraping of data from the websites - Returns PlayerLists */

type Scrape struct {
	Config *models.Configuration
	Folder string
}

func NewScraper(config *models.Configuration) *Scrape {
	return &Scrape{
		Config: config,
		Folder: fmt.Sprintf("assets/data/input/%d/%d/", config.Season, config.Week),
	}
}

// FFS
func (s *Scrape) ScrapeFFS() {
	log.Println("Scraping FFS data to " + s.Folder)

	output, err := exec.Command("php", "assets/scrape/ffs.php", s.Folder, fmt.Sprintf("%d", s.Config.DKID), s.Config.FFScout.Username, s.Config.FFScout.Password).Output()
	log.Println(string(output))
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(string(output), "error") {
		log.Fatal("FSS didn't crawl properly")
	}
}

// Roto
func (s *Scrape) ScrapeRoto() {
	log.Println("Scraping Roto data to " + s.Folder + " " + fmt.Sprintf("%d", s.Config.DKID))
	output, err := exec.Command("php", "assets/scrape/roto.php", s.Folder, fmt.Sprintf("%d", s.Config.DKID)).Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(output))

	if strings.Contains(string(output), "error") {
		log.Fatal("Roto didn't crawl properly")
	}

}

// FPL
// page = transfers_out_event, transfers_out_event, value_form
func (s *Scrape) ScrapeFPL() {
	log.Printf("Scraping FPL data to " + s.Folder + "\n")

	out, err := exec.Command("php", "assets/scrape/fpl.php", s.Folder).Output()
	fmt.Println(string(out))
	if err != nil {
		log.Fatal(err)
	}

}

// DK
func (s *Scrape) ScrapeDK() {

	url := fmt.Sprintf("https://www.draftkings.co.uk/lineup/getavailableplayerscsv?contestTypeId=27&draftGroupId=%d", s.Config.DKID)
	out := s.Folder + fmt.Sprintf("dk-%d.csv", s.Config.DKID)
	log.Println("Scraping DK data to " + s.Folder)

	utils.WGet(url, out)
}

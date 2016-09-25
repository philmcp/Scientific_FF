package scrape

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"strconv"
	"strings"
)

/* Takes a CSVData array and produces PlayerLists */

// FFS
func (s *Scrape) ParseFFS() models.PlayerList {

	fmt.Println("Parsing FFS")

	fileFFS := s.Folder + "ffs-" + fmt.Sprintf("%d", s.Config.DKID) + ".csv"
	csv := LoadCSV(fileFFS)

	posMapping := map[string]string{"gk": "gk", "def": "d", "mid": "m", "fwd": "f"}

	index := -1

	out := models.PlayerList{}
	for row, v := range csv.Data {

		for k, val := range v {
			search := fmt.Sprintf("gw%d", s.Config.Week)
			if index == -1 && strings.Contains(val, search) {
				index = k
				fmt.Printf("FFS Index is %d\n", index)
			}
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		if row == 0 {
			continue
		}

		i, _ := strconv.ParseFloat(v[index], 64)
		team := teamFFS2DK(v[1])

		name := utils.GetLastName(mapDuplicateNames(v[0]))

		cur := models.Player{Name: name, Team: team, Projection: i, Position: posMapping[v[2]]}

		out = append(out, cur)

	}
	return out
}

// DK
func (s *Scrape) ParseDK() models.PlayerList {

	fileDK := s.Folder + "dk-" + fmt.Sprintf("%d", s.Config.DKID) + ".csv"
	csv := LoadCSV(fileDK)

	out := models.PlayerList{}
	fmt.Println("Parsing DK")

	for i, v := range csv.Data {
		if i == 0 {
			continue
		}
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		i, _ := strconv.ParseFloat(strings.Replace(v[2], ".", "", -1), 64)

		name := utils.GetLastName(mapDuplicateNames(v[1]))

		cur := models.Player{Name: name, Team: v[5], Wage: i, Position: v[0]}

		out = append(out, cur)

	}

	return out
}

// FPL
func (s *Scrape) ParseFPL(page string) models.PlayerList {

	fileFPL := s.Folder + "fpl-" + page + ".csv"
	csv := LoadCSV(fileFPL)

	fmt.Println("Parsing FPL")
	out := models.PlayerList{}
	for _, v := range csv.Data {

		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		name := utils.GetLastName(mapDuplicateNames(v[0]))

		cost, _ := strconv.ParseFloat(v[3], 64)
		selected, _ := strconv.ParseFloat(v[4], 64)
		form, _ := strconv.ParseFloat(v[5], 64)
		points, _ := strconv.ParseFloat(v[6], 64)
		data, _ := strconv.ParseFloat(v[7], 64)

		cur := models.Player{Name: name, Team: v[1], Position: v[2], Cost: cost, Selected: selected, Form: form, Points: points}

		if page == "value_form" {
			cur.ValueForm = data
		} else if page == "transfers_out_event" {
			cur.TransfersOutEvent = data
		} else if page == "transfers_in_event" {
			cur.TransfersInEvent = data
		}

		out = append(out, cur)

	}
	return out

}

// Roto
func (s *Scrape) ParseRoto() models.PlayerList {
	fmt.Println("Parsing Roto")
	fileRoto := s.Folder + "roto-players.csv"
	csv := LoadCSV(fileRoto)

	out := models.PlayerList{}

	for _, v := range csv.Data {
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		name := utils.GetLastName(v[1])
		team := teamRoto2DK(v[0])
		returning := v[4] == "true"
		cur := models.Player{Name: name, Team: team, Position: v[2], Status: v[3], ReturningFromInjury: returning}
		out = append(out, cur)
	}

	return out
}

package scrape

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"reflect"
	"strconv"
	"strings"
)

/* Takes a CSVData array and produces PlayerLists */

// FFS
func (s *Scrape) ParseFFS() models.PlayerList {

	log.Println("Parsing FFS")

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
				log.Printf("FFS Index is %d\n", index)
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
	log.Println("Parsing DK")

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
func (s *Scrape) ParseFPL() models.PlayerList {

	fileFPL := s.Folder + "fpl.csv"
	csv := LoadCSV(fileFPL)

	log.Println("Parsing FPL")
	out := models.PlayerList{}

	fmt.Sprintf("%+v\n", csv.Data)
	for i, v := range csv.Data {

		if i == 0 {
			continue
		}

		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		cur := models.Player{Name: utils.GetLastName(mapDuplicateNames(v[0])), Team: v[35]}

		temp := strings.Split("name,now_cost,value_form,value_season,cost_change_start,cost_change_event,cost_change_start_fall,cost_change_event_fall,selected_by_percent,form,transfers_out,transfers_in,transfers_out_event,transfers_in_event,total_points,event_points,points_per_game,minutes,goals_scored,assists,clean_sheets,goals_conceded,own_goals,penalties_saved,penalties_missed,yellow_cards,red_cards,saves,bonus,bps,influence,creativity,threat,ict_index,ea_index,team,team_strength,game_team_strength_overall,game_team_strength_attack,game_team_strength_defence,game_opp_strength_overall,game_opp_strength_attack,game_opp_strength_defence", ",")

		fpl := models.FPL{}
		curCol := 1

		for j := 1; j < len(temp); j++ {
			if j == 35 {
				curCol++
				continue
			}

			val, err := strconv.ParseFloat(v[curCol], 64)
			if err != nil {
				log.Fatal(err)
			}

			id := utils.DBToStructName(temp[j])

			field := reflect.ValueOf(&fpl).Elem().FieldByName(id)

			//	fmt.Printf("Setting %s (%d) to %f\n", id, j, val)

			if field.CanSet() {
				field.SetFloat(val)
			} else {
				fmt.Println("Cant set " + id)
			}
			curCol++
		}

		cur.FPL = fpl

		out = append(out, cur)

	}

	return out

}

// Roto
func (s *Scrape) ParseRoto() models.PlayerList {
	log.Println("Parsing Roto")
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
		cur := models.Player{Name: name, Team: team, Position: v[2],
			Roto: models.Roto{
				Status:              v[3],
				ReturningFromInjury: returning,
			},
		}
		out = append(out, cur)
	}

	return out
}

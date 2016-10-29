package orm

import (
	"github.com/philmcp/Scientific_FF/models"

	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"reflect"
	// Postgres Driver
	//	"fmt"
	_ "github.com/lib/pq"
	"strings"
)

// Fantasy Football Scout
func (db *ORM) LoadFFS(players models.PlayerList) {
	log.Println("Loading FFS to db")
	db.Conn.Exec("DELETE FROM ffs WHERE season = $1 AND week = $2", db.Config.Season, db.Config.Week)

	for _, player := range players {
		db.Conn.Exec("INSERT INTO ffs (name, team, projection, pos, season, week) VALUES ($1, $2, $3, $4, $5, $6)", player.Name, player.Team, player.Projection, player.Position, db.Config.Season, db.Config.Week)
	}
	log.Println("Finshed loading FFS to db")
}

// Draft Kings
func (db *ORM) LoadDK(players models.PlayerList) {
	log.Println("Loading DK to db")
	db.Conn.Exec("DELETE FROM dk WHERE season = $1 AND week = $2 AND dkid = $3", db.Config.Season, db.Config.Week, db.Config.DKID)
	for _, player := range players {
		db.Conn.Exec("INSERT INTO dk (name, team, wage, pos, season, week, dkid, avg_points_per_game) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", player.Name, player.Team, player.Wage, player.Position, db.Config.Season, db.Config.Week, db.Config.DKID, player.AvgPointsPerGame)
	}
	log.Println("Finshed loading DK to db")
}

// FPL
func (db *ORM) LoadFPL(players models.PlayerList) {
	log.Println("Loading FPL to db")
	db.Conn.Exec("DELETE FROM fpl WHERE season = $1 AND week = $2", db.Config.Season, db.Config.Week)

	cols := "now_cost,value_form,value_season,cost_change_start,cost_change_event,cost_change_start_fall,cost_change_event_fall,selected_by_percent,form,transfers_out,transfers_in,transfers_out_event,transfers_in_event,total_points,event_points,points_per_game,minutes,goals_scored,assists,clean_sheets,goals_conceded,own_goals,penalties_saved,penalties_missed,yellow_cards,red_cards,saves,bonus,bps,influence,creativity,threat,ict_index,ea_index,team_strength,game_team_strength_overall,game_team_strength_attack,game_team_strength_defence,game_opp_strength_overall,game_opp_strength_attack,game_opp_strength_defence"

	temp := strings.Split(cols, ",")

	for _, player := range players {

		_, err := db.Conn.Exec("INSERT INTO fpl (name, team, opp_team, is_home, season, week) VALUES ($1, $2, $3, $4, $5, $6)", player.Name, player.Team, player.OppTeam, player.IsHome, db.Config.Season, db.Config.Week)
		if err != nil {
			log.Printf("%+v\n", player)
			log.Println(err)
		}

		tx, err := db.Conn.Begin()
		if err != nil {
			log.Fatal(err)
		}
		defer tx.Rollback()

		// Update each column
		for i := 0; i < len(temp); i++ {
			field := reflect.ValueOf(&player.FPL).Elem().FieldByName(utils.DBToStructName(temp[i])).Float()
			//fmt.Printf("Setting %s to %f\n", temp[i], field)

			stmt, err := tx.Prepare("UPDATE fpl SET " + temp[i] + " = $1 WHERE name = $2 AND team = $3 AND season = $4 AND week = $5")
			if err != nil {
				log.Fatal(err)
			}

			_, err = stmt.Exec(field, player.Name, player.Team, db.Config.Season, db.Config.Week)

			if err != nil {
				log.Fatal(err)
			}

		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
		}

	}

	log.Println("Finshed loading FPL to db")

}

// Roto
func (db *ORM) LoadRoto(players models.PlayerList) {
	db.Conn.Exec("DELETE FROM roto_players WHERE season = $1 AND week = $2", db.Config.Season, db.Config.Week)

	for _, player := range players {
		db.Conn.Exec("INSERT INTO roto_players (team, name, pos, status, returning_from_injury, week, season) VALUES ($1, $2, $3, $4, $5, $6, $7)", player.Team, player.Name, player.Position, player.Roto.Status, player.Roto.ReturningFromInjury, db.Config.Week, db.Config.Season)
	}

}

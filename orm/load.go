package orm

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	// Postgres Driver
	_ "github.com/lib/pq"
)

// Fantasy Football Scout
func (db *ORM) LoadFFS(players models.PlayerList) {
	fmt.Println("Loading FFS to db")
	db.Conn.Exec("DELETE FROM ffs WHERE season = $1 AND week = $2", db.Config.Season, db.Config.Week)

	for _, player := range players {
		db.Conn.Exec("INSERT INTO ffs (name, team, projection, pos, season, week) VALUES ($1, $2, $3, $4, $5, $6)", player.Name, player.Team, player.Projection, player.Position, db.Config.Season, db.Config.Week)
	}
	fmt.Println("Finshed loading FFS to db")
}

// Draft Kings
func (db *ORM) LoadDK(players models.PlayerList) {
	fmt.Println("Loading DK to db")
	db.Conn.Exec("DELETE FROM dk WHERE season = $1 AND week = $2 AND dkid = $3", db.Config.Season, db.Config.Week, db.Config.DKID)
	for _, player := range players {
		db.Conn.Exec("INSERT INTO dk (name, team, wage, pos, season, week, dkid) VALUES ($1, $2, $3, $4, $5, $6, $7)", player.Name, player.Team, player.Wage, player.Position, db.Config.Season, db.Config.Week, db.Config.DKID)
	}
	fmt.Println("Finshed loading DK to db")
}

// FPL
func (db *ORM) LoadFPL(players models.PlayerList, page string) {
	res, _ := db.Conn.Query("SELECT * FROM fpl WHERE season = $1 AND week = $2", db.Config.Season, db.Config.Week)

	exists := res.Next()

	for _, player := range players {

		data := player.ValueForm

		if page == "transfers_out_event" {
			data = player.TransfersOutEvent
		} else if page == "transfers_in_event" {
			data = player.TransfersInEvent
		}

		if !exists {
			// Insert for first time
			_, err := db.Conn.Exec(`INSERT INTO fpl (name,
		team,
		pos,
		cost,
		selected,
		form,
		points,
		`+page+`,
		week,
		season) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, player.Name, player.Team, player.Position, player.Cost, player.Selected, player.Form, player.Points, data, db.Config.Week, db.Config.Season)
			if err != nil {
				fmt.Println(err)
			}

		} else {

			// Update
			_, err := db.Conn.Exec(`UPDATE fpl SET
		cost= $3,
		selected =$4,
		form =$5,
		points = $6,
		`+page+` = $7
		WHERE	name = $1 AND team = $2 AND week = $8 AND season = $9`, player.Name, player.Team, player.Cost, player.Selected, player.Form, player.Points, data, db.Config.Week, db.Config.Season)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}

// Roto
func (db *ORM) LoadRoto(players models.PlayerList) {
	db.Conn.Exec("DELETE FROM roto_players WHERE season = $1 AND week = $2", db.Config.Season, db.Config.Week)

	for _, player := range players {
		db.Conn.Exec("INSERT INTO roto_players (team, name, pos, status, returning_from_injury, week, season) VALUES ($1, $2, $3, $4, $5, $6, $7)", player.Team, player.Name, player.Position, player.Status, player.ReturningFromInjury, db.Config.Week, db.Config.Season)
	}

}

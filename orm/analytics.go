package orm

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/philmcp/Scientific_FF/models"
	"log"
)

func (db *ORM) GetMostTransferredIn() *models.PlayerList {

	rows, _ := db.Conn.Query(`SELECT name, team, transfers_in_event  FROM fpl
WHERE week = $1
ORDER BY transfers_in_event DESC
LIMIT 4
`, db.Config.Week)
	defer rows.Close()

	var (
		name      sql.NullString
		team      sql.NullString
		transfers sql.NullFloat64
	)

	players := models.PlayerList{}

	for rows.Next() {
		err := rows.Scan(&name, &team, &transfers)
		if err != nil {
			log.Println(err)
		}

		cur := models.Player{Name: name.String, Team: team.String, FPL: models.FPL{TransfersInEvent: transfers.Float64}}
		players = append(players, cur)

	}
	return &players
}

func (db *ORM) GetMostTransferredOut() *models.PlayerList {

	rows, _ := db.Conn.Query(`SELECT name, team, transfers_out_event  FROM fpl
WHERE week = $1
ORDER BY transfers_out_event DESC
LIMIT 4
`, db.Config.Week)
	defer rows.Close()

	var (
		name      sql.NullString
		team      sql.NullString
		transfers sql.NullFloat64
	)

	players := models.PlayerList{}

	for rows.Next() {
		err := rows.Scan(&name, &team, &transfers)
		if err != nil {
			log.Println(err)
		}

		cur := models.Player{Name: name.String, Team: team.String, FPL: models.FPL{TransfersOutEvent: transfers.Float64}}
		players = append(players, cur)

	}
	return &players
}

func (db *ORM) GetInForm() *models.PlayerList {

	rows, _ := db.Conn.Query(`SELECT name, team, form  FROM fpl
WHERE week = $1 AND form IS NOT NULL
ORDER BY form DESC
LIMIT 4

`, db.Config.Week)
	defer rows.Close()

	var (
		name sql.NullString
		team sql.NullString
		form sql.NullFloat64
	)

	players := models.PlayerList{}

	for rows.Next() {
		err := rows.Scan(&name, &team, &form)
		if err != nil {
			log.Println(err)
		}

		cur := models.Player{Name: name.String, Team: team.String, FPL: models.FPL{Form: form.Float64}}
		players = append(players, cur)

	}
	return &players
}

package main

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres Driver
	_ "github.com/lib/pq"
	"strings"
)

type DB struct {
	conn *sql.DB
}

// ConfigureDB sets credentials for db connect
func Connect(host string, port int, database string, user string, password string) *DB {
	fmt.Println("Starting up " + database + " db")
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &DB{db}
}

func (db *DB) getPlayers() PlayerPool {
	rows, _ := db.conn.Query(`SELECT dk.name, dk.team, dk.pos, status, returning_from_injury, wage, projection FROM roto_players
LEFT JOIN dk ON dk.week = roto_players.week AND dk.season = roto_players.season AND roto_players.name = dk.name AND roto_players.team = dk.team
LEFT JOIN ffs ON dk.week = ffs.week AND dk.season = roto_players.season AND ffs.name = dk.name AND ffs.team = dk.team
WHERE dk.season = $1 AND dk.week = $2 AND dkid = $3
ORDER BY dk.team ASC`, SEASON, WEEK, DKID)

	out := PlayerPool{}

	for pos, _ := range formation {
		out[pos] = []Player{}
	}
	for rows.Next() {
		var (
			name                  sql.NullString
			team                  sql.NullString
			pos                   sql.NullString
			status                sql.NullString
			returning_from_injury sql.NullBool
			wage                  sql.NullFloat64
			projection            sql.NullFloat64
		)

		rows.Scan(&name, &team, &pos, &status, &returning_from_injury, &wage, &projection)
		cur := Player{Name: name.String, Team: team.String, Status: status.String, ReturningFromInjury: returning_from_injury.Bool, Wage: wage.Float64, Projection: projection.Float64}

		// Can play multiple positions
		if strings.Contains(pos.String, "/") {
			spl := strings.Split(pos.String, "/")
			out[spl[0]] = append(out[spl[0]], cur)
			out[spl[1]] = append(out[spl[1]], cur)
		} else {
			out[pos.String] = append(out[pos.String], cur)
		}

		out["u"] = append(out["u"], cur)
	}

	for pos, _ := range formation {
		fmt.Printf("Player pool (%s): %d\n", pos, len(out[pos]))
	}

	return out

}

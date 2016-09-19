package main

import (
	"database/sql"
	"fmt"
	"log"

	// Postgres Driver
	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

// ConfigureDB sets credentials for db connect
func Connect(host string, port int, database string, user string, password string) *DB {
	log.Println("Starting up " + database + " db")
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, database)
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		fmt.Println(err.Error())
	}
	return &DB{db}
}

func (db *DB) GetPlayers() {
	rows, _ := db.conn.Query("SELECT name FROM players")

	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(name)
	}

}

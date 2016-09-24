package orm

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

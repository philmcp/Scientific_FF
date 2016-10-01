package orm

import (
	"database/sql"
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"log"
	// Postgres Driver
	_ "github.com/lib/pq"
)

type ORM struct {
	Conn   *sql.DB
	Config *models.Configuration
}

// ConfigureDB sets credentials for db connect
func NewORM(config *models.Configuration) *ORM {
	log.Println("Setting up " + config.Database.DB + " db")
	dbString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DB)

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &ORM{Conn: db, Config: config}
}

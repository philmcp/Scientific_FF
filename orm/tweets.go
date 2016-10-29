package orm

import (
	"github.com/philmcp/Scientific_FF/models"
	"log"
	// Postgres Driver
	_ "github.com/lib/pq"
)

func (db *ORM) InsertTweet(t *models.Tweet) {
	_, err := db.Conn.Exec("INSERT INTO tweets (id, text, screen_name, timestamp) VALUES ($1, $2, $3, $4)", t.ID, t.Text, t.ScreenName, t.Timestamp)
	if err != nil {
		log.Println(err)
	}
}

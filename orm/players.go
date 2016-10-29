package orm

import (
	"database/sql"
	"github.com/philmcp/Scientific_FF/models"
	"log"
	// Postgres Driver
	_ "github.com/lib/pq"
	"strings"
)

func (db *ORM) GetPlayers() (models.PlayerPool, float64) {
	rows, _ := db.Conn.Query(`SELECT dk.name, dk.team, fpl.opp_team, fpl.is_home, dk.pos, status, returning_from_injury, wage, projection/max_projection as projection, selected_by_percent/max_selected as selected,  form/max_form as form, total_points/max_points as points,

CASE WHEN transfers_out_event = 0 THEN 0 ELSE (transfers_in_event/transfers_out_event)/max_transfers END as transfer_ratio,
 value_form/max_value_form as value_form, points_per_game/max_points_per_game as points_per_game,
  avg_points_per_game/max_avg_points_per_game as avg_points_per_game
 FROM roto_players
LEFT JOIN dk ON dk.week = roto_players.week AND dk.season = roto_players.season AND roto_players.name = dk.name AND roto_players.team = dk.team
LEFT JOIN ffs ON dk.week = ffs.week AND dk.season = roto_players.season AND ffs.name = dk.name AND ffs.team = dk.team
LEFT JOIN fpl ON dk.week = fpl.week AND dk.season = fpl.season AND fpl.name = dk.name AND fpl.team = dk.team
JOIN (SELECT
	MAX(projection) max_projection,
	MAX(selected_by_percent) max_selected,
	MAX(form) max_form,
	MAX(total_points) max_points,
	MAX(value_form) max_value_form,
	MAX(transfers_in_event) max_transfers_in_event,
	MAX(transfers_out_event) max_transfers_out_event,
	MAX(points_per_game) max_points_per_game,
MAX(CASE WHEN transfers_out_event = 0 THEN 0 ELSE transfers_in_event/transfers_out_event END) as max_transfers,
	MAX(avg_points_per_game) max_avg_points_per_game
	FROM roto_players
	LEFT JOIN dk ON dk.week = roto_players.week AND dk.season = roto_players.season AND roto_players.name = dk.name AND roto_players.team = dk.team
	LEFT JOIN ffs ON dk.week = ffs.week AND dk.season = roto_players.season AND ffs.name = dk.name AND ffs.team = dk.team
	LEFT JOIN fpl ON dk.week = fpl.week AND dk.season = fpl.season AND fpl.name = dk.name AND fpl.team = dk.team
WHERE dk.season = $1 AND dk.week = $2 AND dkid = $3
	)
AS vals ON true
WHERE dk.season = $1 AND dk.week = $2 AND dkid = $3
ORDER BY dk.team ASC
`, db.Config.Season, db.Config.Week, db.Config.DKID)
	defer rows.Close()

	out := models.PlayerPool{}
	maxScore := 0.0

	for pos, _ := range db.Config.Formation {
		out[pos] = []models.Player{}
	}
	for rows.Next() {
		var (
			name                  sql.NullString
			team                  sql.NullString
			oppTeam               sql.NullString
			isHome                sql.NullBool
			pos                   sql.NullString
			status                sql.NullString
			returning_from_injury sql.NullBool
			wage                  sql.NullFloat64
			projection            sql.NullFloat64
			selected              sql.NullFloat64
			form                  sql.NullFloat64
			points                sql.NullFloat64
			value_form            sql.NullFloat64
			transfer_ratio        sql.NullFloat64
			pointsPerGame         sql.NullFloat64
			avgPointsPerGame      sql.NullFloat64
		)

		rows.Scan(&name, &team, &oppTeam, &isHome, &pos, &status, &returning_from_injury, &wage, &projection, &selected, &form, &points, &value_form, &transfer_ratio, &pointsPerGame, &avgPointsPerGame)
		cur := models.Player{
			Name:             name.String,
			Team:             team.String,
			OppTeam:          oppTeam.String,
			IsHome:           isHome.Bool,
			Wage:             wage.Float64,
			Projection:       projection.Float64,
			Position:         pos.String,
			PositionRaw:      pos.String,
			AvgPointsPerGame: avgPointsPerGame.Float64,
			Roto:             models.Roto{Status: status.String, ReturningFromInjury: returning_from_injury.Bool},
			FPL: models.FPL{
				SelectedByPercent: selected.Float64,
				Form:              form.Float64,
				TotalPoints:       points.Float64,
				ValueForm:         value_form.Float64,
				PointsPerGame:     pointsPerGame.Float64,
			},
		}

		if cur.GetScore() > maxScore {
			maxScore = cur.GetScore()
		}

		//log.Printf("%+v\n\n", cur)
		// Can play multiple positions
		if strings.Contains(pos.String, "/") {
			spl := strings.Split(pos.String, "/")
			temp := cur
			temp.Position = spl[0]
			out[spl[0]] = append(out[spl[0]], temp)

			if temp.IsGoodUtility() {
				out["u"] = append(out["u"], temp)
			}

			temp2 := cur
			temp2.Position = spl[1]
			out[spl[1]] = append(out[spl[1]], temp2)

			if temp2.IsGoodUtility() {
				out["u"] = append(out["u"], temp2)
			}

		} else {
			out[pos.String] = append(out[pos.String], cur)

			if cur.IsGoodUtility() {
				out["u"] = append(out["u"], cur)
			}

		}

		// Only consider m/f players for utility
	}

	for pos, _ := range db.Config.Formation {
		log.Printf("Player pool (%s): %d\n", pos, len(out[pos]))
	}

	return out, maxScore

}

package main

import (
	"fmt"

	//"log"
	"strconv"
	"strings"
)

// Some functions to parse the data from file

// Fantasy Football Scout
func (db *DB) parseFFSDB(csv *CSVData) {

	// We use DK (i.e. the value) as the golden standard
	posMapping := map[string]string{"gk": "gk", "def": "d", "mid": "m", "fwd": "f"}

	index := -1

	db.conn.Exec("DELETE FROM ffs WHERE season = $1 AND week = $2", season, gameWeek)

	for row, v := range csv.Data {

		// Get the index
		for k, val := range v {
			search := fmt.Sprintf("gw%d", gameWeek)
			if index == -1 && strings.Contains(val, search) {
				index = k
				fmt.Printf("Index is %d\n", index)
			}
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		if row == 0 {
			continue
		}

		i, _ := strconv.ParseFloat(v[index], 64)
		team := teamFFS2DK(v[1])

		name := getLastName(v[0])

		db.conn.Exec("INSERT INTO ffs (name, team, projection, pos, season, week) VALUES ($1, $2, $3, $4, $5, $6)", name, team, i, posMapping[v[2]], season, gameWeek)

	}

}

// Draft Kings
func (db *DB) parseDKDB(csv *CSVData) {
	db.conn.Exec("DELETE FROM dk WHERE season = $1 AND week = $2 AND dkid = $3", season, gameWeek, DKID)
	for i, v := range csv.Data {
		if i == 0 {
			continue
		}
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		i, _ := strconv.ParseFloat(strings.Replace(v[2], ".", "", -1), 64)

		name := getLastName(v[1])

		db.conn.Exec("INSERT INTO dk (name, team, wage, pos, season, week, dkid) VALUES ($1, $2, $3, $4, $5, $6, $7)", name, v[5], i, v[0], season, gameWeek, DKID)
	}

}

// Roto
func (db *DB) parseRotoPlayersDB(csv *CSVData) {
	db.conn.Exec("DELETE FROM roto_players WHERE season = $1 AND week = $2", season, gameWeek, DKID)
	for _, v := range csv.Data {
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		name := getLastName(v[1])

		db.conn.Exec("INSERT INTO roto_players (team, name, pos, status, returning_from_injury, week, season) VALUES ($1, $2, $3, $4, $5, $6, $7)", v[0], name, v[2], v[3], v[4], gameWeek, season)
	}

}

func getLastName(str string) string {
	source := parseEncoding(strings.ToLower(str))

	names := strings.Split(source, " ")
	name := names[len(names)-1]
	return name
}

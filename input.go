package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"log"
	"os"
	"strconv"
	"strings"
)

// Some functions to parse the data from file
var enc = charmap.Windows1252

func loadCSV(filename string) (csv CSVData) {
	fmt.Println("Loading CSV: " + filename)
	f, _ := os.Open(filename)

	r := transform.NewReader(f, enc.NewDecoder())
	scanner := bufio.NewScanner(r)

	i := 0

	for scanner.Scan() {
		curLine := ParseEncoding(scanner.Text())
		curRow := strings.Split(curLine, ",")

		if len(curRow) > 2 {
			csv.Data = append(csv.Data, curRow)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return csv
}

// Fantasy Football Scout
func (db *DB) loadFFS(csv *CSVData) {

	// We use DK (i.e. the value) as the golden standard
	posMapping := map[string]string{"gk": "gk", "def": "d", "mid": "m", "fwd": "f"}

	index := -1

	_, err := db.conn.Exec("DELETE FROM ffs WHERE season = $1 AND week = $2", SEASON, WEEK)
	if err != nil {
		log.Fatal(err)
	}

	for row, v := range csv.Data {

		// Get the index
		for k, val := range v {
			search := fmt.Sprintf("gw%d", WEEK)
			if index == -1 && strings.Contains(val, search) {
				index = k
				fmt.Printf("FFS Index is %d\n", index)
			}
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		if row == 0 {
			continue
		}

		i, _ := strconv.ParseFloat(v[index], 64)
		team := teamFFS2DK(v[1])

		name := getLastName(mapDuplicateNames(v[0]))

		db.conn.Exec("INSERT INTO ffs (name, team, projection, pos, season, week) VALUES ($1, $2, $3, $4, $5, $6)", name, team, i, posMapping[v[2]], SEASON, WEEK)

	}

}

// Draft Kings
func (db *DB) loadDK(csv *CSVData) {
	_, err := db.conn.Exec("DELETE FROM dk WHERE season = $1 AND week = $2 AND dkid = $3", SEASON, WEEK, DKID)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range csv.Data {
		if i == 0 {
			continue
		}
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		i, _ := strconv.ParseFloat(strings.Replace(v[2], ".", "", -1), 64)

		name := getLastName(mapDuplicateNames(v[1]))

		db.conn.Exec("INSERT INTO dk (name, team, wage, pos, season, week, dkid) VALUES ($1, $2, $3, $4, $5, $6, $7)", name, v[5], i, v[0], SEASON, WEEK, DKID)
	}

}

// Roto
func (db *DB) loadRoto(csv *CSVData) {
	_, err := db.conn.Exec("DELETE FROM roto_players WHERE season = $1 AND week = $2", SEASON, WEEK)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range csv.Data {
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		name := getLastName(v[1])
		team := teamRoto2DK(v[0])

		db.conn.Exec("INSERT INTO roto_players (team, name, pos, status, returning_from_injury, week, season) VALUES ($1, $2, $3, $4, $5, $6, $7)", team, name, v[2], v[3], v[4], WEEK, SEASON)
	}

}

func getLastName(str string) string {
	source := ParseEncoding(strings.ToLower(str))

	names := strings.Split(source, " ")
	name := names[len(names)-1]
	return name
}

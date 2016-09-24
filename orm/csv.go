package orm

// Fantasy Football Scout
func (db *DB) loadFFS(csv *CSVData) {

	// We use DK (i.e. the value) as the golden standard
	posMapping := map[string]string{"gk": "gk", "def": "d", "mid": "m", "fwd": "f"}

	index := -1

	_, err := db.conn.Exec("DELETE FROM ffs WHERE season = $1 AND week = $2", conf.Season, conf.Week)
	if err != nil {
		log.Fatal(err)
	}

	for row, v := range csv.Data {

		// Get the index
		for k, val := range v {
			search := fmt.Sprintf("gw%d", conf.Week)
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

		db.conn.Exec("INSERT INTO ffs (name, team, projection, pos, season, week) VALUES ($1, $2, $3, $4, $5, $6)", name, team, i, posMapping[v[2]], conf.Season, conf.Week)

	}

}

// Draft Kings
func (db *DB) loadDK(csv *CSVData) {
	_, err := db.conn.Exec("DELETE FROM dk WHERE season = $1 AND week = $2 AND dkid = $3", conf.Season, conf.Week, conf.DKID)
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

		db.conn.Exec("INSERT INTO dk (name, team, wage, pos, season, week, dkid) VALUES ($1, $2, $3, $4, $5, $6, $7)", name, v[5], i, v[0], conf.Season, conf.Week, conf.DKID)
	}

}

// FPL
func (db *DB) loadFPL(csv *CSVData, page string) {
	res, err := db.conn.Query("SELECT * FROM fpl WHERE season = $1 AND week = $2", conf.Season, conf.Week)

	if err != nil {
		log.Fatal(err)
	}

	exists := res.Next()

	for _, v := range csv.Data {
		//	if i == 0 {
		//		continue
		//	}
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		name := getLastName(mapDuplicateNames(ParseEncoding(v[0])))

		cost, _ := strconv.ParseFloat(v[3], 64)
		selected, _ := strconv.ParseFloat(v[4], 64)
		form, _ := strconv.ParseFloat(v[5], 64)
		points, _ := strconv.ParseFloat(v[6], 64)
		data, _ := strconv.ParseFloat(v[7], 64)

		if !exists {
			// Insert for first time
			_, err := db.conn.Exec(`INSERT INTO fpl (name,
		team,
		pos,
		cost,
		selected,
		form,
		points,
		`+page+`,
		week,
		season) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, name, v[1], v[2], cost, selected, form, points, data, conf.Week, conf.Season)
			if err != nil {
				fmt.Println(err)
			}

		} else {

			// Update
			_, err := db.conn.Exec(`UPDATE fpl SET
		pos= $3,
		cost= $4,
		selected =$5,
		form =$6,
		points = $7,
		`+page+` = $8 WHERE	name = $1 AND team = $2 AND week = $9 AND season = $10`, name, v[1], v[2], cost, selected, form, points, data, conf.Week, conf.Season)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

}

// Roto
func (db *DB) loadRoto(csv *CSVData) {
	_, err := db.conn.Exec("DELETE FROM roto_players WHERE season = $1 AND week = $2", conf.Season, conf.Week)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range csv.Data {
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		name := getLastName(v[1])
		team := teamRoto2DK(v[0])

		db.conn.Exec("INSERT INTO roto_players (team, name, pos, status, returning_from_injury, week, season) VALUES ($1, $2, $3, $4, $5, $6, $7)", team, name, v[2], v[3], v[4], conf.Week, conf.Season)
	}

}

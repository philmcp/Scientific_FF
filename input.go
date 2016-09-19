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

type CSVData struct {
	Data [][]string
}

func loadCSV(filename string) (csv CSVData) {
	fmt.Println("Loading CSV: " + filename)
	f, _ := os.Open(filename)

	r := transform.NewReader(f, enc.NewDecoder())
	scanner := bufio.NewScanner(r)

	i := 0

	for scanner.Scan() {
		curLine := parseEncoding(scanner.Text())
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

func parseEncoding(str string) string {
	str = strings.ToLower(str)

	str = strings.Replace(str, "è", "e", -1)
	str = strings.Replace(str, "á", "a", -1)
	str = strings.Replace(str, "é", "e", -1)
	str = strings.Replace(str, "ö", "o", -1)
	str = strings.Replace(str, "í", "i", -1)
	str = strings.Replace(str, "à", "a", -1)
	str = strings.Replace(str, "ó", "o", -1)
	str = strings.Replace(str, "ú", "u", -1)
	str = strings.Replace(str, "ü", "u", -1)

	return str
}

// Fantasy Football Scout
func (csv *CSVData) parseFFS() (l League) {
	l.Name = "Football Fantasy Scout"

	// We use DK (i.e. the value) as the golden standard
	posMapping := map[string]string{"gk": "gk", "def": "d", "mid": "m", "fwd": "f"}

	l.Teams = map[string]Team{}

	index := -1

	for _, v := range csv.Data {

		for k, val := range v {
			search := fmt.Sprintf("gw%d", gameWeek)
			if index == -1 && strings.Contains(val, search) {
				index = k
				fmt.Printf("Index is %d\n", index)
			}
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}

		if index < 0 {
			log.Fatal("Index not set")
		}

		i, _ := strconv.ParseFloat(v[index], 64)

		team := teamFFS2DK(v[1])

		name := parseName(v[0], team, l.Name)

		// Weird way of setting the number of players so that the MAX_VALUE thresh creates a big enough lineup to pick from
		if i > threads {
			threads = i
			fmt.Printf("New num threads is %f (%f)\n", threads, i)
		}

		curPlayer := Player{Name: name, Team: team, Position: posMapping[v[2]], Value: i}
		l.addPlayer(&curPlayer)

	}
	delete(l.Teams, "team")

	return l

}

// Draft Kings
func (csv *CSVData) parseDK() (l League) {
	l.Name = "Draft Kings"

	l.Teams = map[string]Team{}
	for i, v := range csv.Data {
		if i == 0 {
			continue
		}
		for k, _ := range v {
			v[k] = strings.TrimSpace(strings.Replace(v[k], "\"", "", -1))
		}
		i, _ := strconv.ParseFloat(strings.Replace(v[2], ".", "", -1), 64)

		name := parseName(v[1], v[5], l.Name)

		curPlayer := Player{Name: name, Team: v[5], Position: v[0], Wage: i}
		l.addPlayer(&curPlayer)

		// Get the games for this set - e.g. STK@LIV 11:00AM ET
		teams := strings.Split(v[3], " ")
		game := teams[0] + " vs " + teams[2]
		if _, exist := weekGames[game]; !exist {

			cur := Game{HomeTeam: teams[0], AwayTeam: teams[2], Time: teams[3] + " " + teams[4]}

			weekGames[game] = cur
		}
	}

	return l

}

// Load FFS lineups from http://www.fantasyfootballscout.co.uk/team-news/
func (csv *CSVData) parseFFSP() (l League) {
	l.Name = "FFS Lineup"

	l.Teams = map[string]Team{}
	for i, v := range csv.Data {
		if i == 0 {
			continue
		}
		for k, _ := range v {
			v[k] = strings.TrimSpace(v[k])
		}

		name := parseName(v[1], v[0], l.Name)

		curPlayer := Player{Name: name, Team: v[0], Position: "u"}
		l.addPlayer(&curPlayer)

	}

	return l

}

// Given a list of games, parses the PL files, then produces a league that can be compared against the DK league (to remove non starting players)
func parsePL() (l League) {
	fmt.Printf("Parsing PL - Games : %d\n", len(weekGames))
	l.Name = "Premier League"

	l.Teams = map[string]Team{}

	l.FinalPlayers = map[string][]Player{}
	l.FinalPlayers["playing"] = []Player{}

	for _, g := range weekGames {
		loc := inputFolder + "games/" + g.getPLURL() + ".csv"
		fmt.Println("Parsing " + loc)
		csv := loadCSV(loc)

		if len(csv.Data) < 5 {
			fmt.Println("The lineups haven't been announced yet!\n")
			continue
		} else {
			fmt.Println("The lineups have been announced!")
		}

		/* Team,Started,Name,EA PPI,FPL,BFR
		Norwich,Y,John Ruddy,240,2.0,3.56 */
		for i, v := range csv.Data {
			if i == 0 || v[1] == "n" {
				continue
			}

			team := teamPL2DK(v[0])
			name := v[2]

			name = parseName(name, team, l.Name)

			curPlayer := Player{Name: name, Team: team, Position: "playing", Wage: 0}

			l.addPlayer(&curPlayer)
			l.FinalPlayers["playing"] = append(l.FinalPlayers["playing"], curPlayer)
		}

	}

	return l
}

// Convert player name to a single, unified name
func parseName(str string, team string, source string) string {
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Replace(str, "\u00e1", "a", -1)
	str = strings.Replace(str, "ã¨", "e", -1)
	str = strings.Replace(str, "ã¡", "a", -1)
	str = strings.Replace(str, "ã©", "e", -1)

	source = strings.ToLower(source)

	if _, existSource := rename[source]; existSource {
		if _, existTeam := rename[source][team]; existTeam {
			if newPlayer, existsPlayer := rename[source][team][str]; existsPlayer {
				fmt.Println("Changing " + source + ": " + str + " to " + newPlayer)
				str = newPlayer
			}
		}
	}

	names := strings.Split(str, " ")
	name := names[len(names)-1]
	return name
}

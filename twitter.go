package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type TwitterAPI struct {
	AppKey    string
	AppSecret string
}

type Tweet struct {
	ID         int64  `json:"id"`
	Text       string `json:"text"`
	Timestamp  int64  `json:"timestamp_ms"`
	ScreenName string `json:"screen_name"`
}

type Injury struct {
	Name    string
	Injury  string
	Returns string
	Team    string
	Perc    string
}

func (api *TwitterAPI) getInjuryNews() {

	fmt.Println("Downloading injury new from twitter")

	query := "from:PremierInjuries FPL"
	t := &url.URL{Path: query}
	enc := t.String()

	output, err := exec.Command("php", "scrape/twitter.php", api.AppKey, api.AppSecret, enc, "3").Output()
	if err != nil {
		log.Fatal(err)
	}

	tweets := parseTweets(output)
	//fmt.Printf("%+v\n", tweets)
	for _, tweet := range tweets {

		if !strings.Contains(tweet.Text, "#FPL") {
			continue
		}

		if err != nil {
			fmt.Println(err)
		} else {
			if _, err := os.Stat(fmt.Sprintf("injury/output/%d.jpg", tweet.ID)); os.IsNotExist(err) {

				fmt.Printf("Tweet %d hasnt been seen before, posting injury\n", tweet.ID)
				inj := parseInjury(tweet.Text)

				// Has this player played this season?
				name := getLastName(inj.Name)
				isWorthyPlayer, _ := db.conn.Query("SELECT name FROM dk WHERE season = $1 AND name = $2", conf.Season, name)

				if isWorthyPlayer.Next() {
					conf.Buffer.postInjury(&inj, &tweet)
				} else {
					fmt.Println(name + " is NOT worthy...")
				}
			}
		}

	}

}

func parseInjury(text string) Injury {
	ret := Injury{}
	text = strings.Replace(text, "No Return Date", "Expected Return: Unknown", -1)

	spl := strings.Split(text, ": ")

	name_spl := strings.Split(spl[1], " - ")

	ret.Name = strings.TrimSpace(name_spl[0])

	inj_spl := strings.Split(strings.TrimSpace(name_spl[1]), "#")
	ret.Injury = strings.TrimSpace(strings.Replace(inj_spl[0], "Expected Return", "", -1))

	if strings.Count(text, "#") > 1 {
		team_spl := strings.Split(text, "#")
		team_spl2 := strings.Split(team_spl[2], " ")
		ret.Team = strings.ToLower(team_spl2[0])
	} else {
		ret.Team = ""
	}

	ret.Returns = strings.TrimSpace(strings.Replace(spl[2], "Status", "", -1))

	perc_spl := strings.Split(spl[3], " ")
	if strings.Contains(perc_spl[0], "%") {
		ret.Perc = strings.Replace(strings.TrimSpace(perc_spl[0]), "%", "", -1)
	} else {
		ret.Perc = "0"
	}

	return ret
}

func parseTweets(text []byte) []Tweet {
	ret := []Tweet{}

	var dat map[string]interface{}
	if err := json.Unmarshal(text, &dat); err != nil {
		fmt.Println(err)
	}
	tweets := dat["statuses"].([]interface{})

	for _, val := range tweets {
		cur := val.(map[string]interface{})
		user := cur["user"].(map[string]interface{})
		tweet := Tweet{}
		tweet.ID = int64(cur["id"].(float64))
		tweet.ScreenName = user["screen_name"].(string)
		tweet.Text = cur["text"].(string)
		tweet.Timestamp = (tweet.ID >> 22) + 1288834974657
		ret = append(ret, tweet)
	}

	return ret
}

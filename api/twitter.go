package api

import (
	"encoding/json"
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"log"
	"net/url"

	"os/exec"
	"strings"
)

type TwitterAPI struct {
	Config *models.Configuration
}

func NewTwitter(config *models.Configuration) *TwitterAPI {
	return &TwitterAPI{config}
}

func (api *TwitterAPI) getInjuryNews() []models.Tweet {

	fmt.Println("Downloading injury new from twitter")

	query := "from:PremierInjuries FPL"
	t := &url.URL{Path: query}
	enc := t.String()

	output, err := exec.Command("php", "scrape/twitter.php", api.Config.Twitter.AppKey, api.Config.Twitter.AppSecret, enc, "3").Output()
	if err != nil {
		log.Fatal(err)
	}

	return parseTweets(output)
	//fmt.Printf("%+v\n", tweets)

}

func parseInjury(text string) models.Injury {
	ret := models.Injury{}
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

func parseTweets(text []byte) []models.Tweet {
	ret := []models.Tweet{}

	var dat map[string]interface{}
	if err := json.Unmarshal(text, &dat); err != nil {
		fmt.Println(err)
	}
	tweets := dat["statuses"].([]interface{})

	for _, val := range tweets {
		cur := val.(map[string]interface{})
		user := cur["user"].(map[string]interface{})
		tweet := models.Tweet{}
		tweet.ID = int64(cur["id"].(float64))
		tweet.ScreenName = user["screen_name"].(string)
		tweet.Text = cur["text"].(string)
		tweet.Timestamp = (tweet.ID >> 22) + 1288834974657
		ret = append(ret, tweet)
	}

	return ret
}

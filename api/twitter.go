package api

import (
	"encoding/json"

	"github.com/philmcp/Scientific_FF/models"
	"log"
	"net/url"

	"os/exec"
)

type TwitterAPI struct {
	Config *models.Configuration
}

func NewTwitter(config *models.Configuration) *TwitterAPI {
	return &TwitterAPI{config}
}

func (api *TwitterAPI) GetInjuryNews() []models.Tweet {

	log.Println("Downloading injury new from twitter")

	query := "from:PremierInjuries FPL -100"
	t := &url.URL{Path: query}
	enc := t.String()

	output, err := exec.Command("php", "assets/scrape/twitter.php", api.Config.Twitter.AppKey, api.Config.Twitter.AppSecret, enc, "3").Output()

	if err != nil {
		log.Fatal(err)
	}

	return parseTweets(output)

}

func parseTweets(text []byte) []models.Tweet {
	ret := []models.Tweet{}

	var dat map[string]interface{}
	if err := json.Unmarshal(text, &dat); err != nil {
		log.Println(err)
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

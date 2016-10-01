package api

import (
	//"crypto/tls"
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"log"
	"net/url"
	"strings"
)

type BufferAPI struct {
	Config *models.Configuration
}

var (
	lineupText = []string{"Its #DFF time! Here is our #EPL #Draftkings lineup for #GW%d",
		"This week's #EPL #Draftkings lineup! #GW%d - Good luck all!",
		"Here's our lineup for today's #GW%d #Draftkings #DFF",
		"Algorithm driven #EPL #DFF lineups - here's today's #GW%d selection",
		"...and here it is, good luck all! #EPL #DFF #DraftKings #GW%d",
		"Here is our #AI selected #GW%d #DraftKings lineup - Good luck!",
		"Our algorithm driven #EPL #DraftKings lineups for #GW%d is here! Good luck!"}
)

func NewBuffer(config *models.Configuration) *BufferAPI {
	return &BufferAPI{config}
}

// Access Buffer.com REST API
func (b *BufferAPI) Post(text string, image string) {

	text = strings.Replace(text, "%", "", -1)

	encText, err := url.Parse(text)
	if err != nil {
		log.Println(err)
	}

	imgText, err := url.Parse(image)
	if err != nil {
		log.Println(err)
	}

	log.Printf("(BUFFER) text: %s image: %s\n", encText.String(), image)

	data := "text=" + encText.String() + "&now=true&profile_ids[]=" + b.Config.Buffer.TwitterID + "&profile_ids[]=" + b.Config.Buffer.FacebookID

	if image != "" {
		data += "&media[photo]=" + imgText.String()
	}
	updateUrl := fmt.Sprintf("https://api.bufferapp.com/1/updates/create.json?access_token=%s", b.Config.Buffer.AccessToken)
	utils.MakeRequest(updateUrl, "POST", data)
}

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

	// Twitter
	data := "text=" + encText.String() + "&now=true&profile_ids[]=" + b.Config.Buffer.TwitterID
	if image != "" {
		data += "&media[photo]=" + imgText.String()
	}
	updateUrl := fmt.Sprintf("https://api.bufferapp.com/1/updates/create.json?access_token=%s", b.Config.Buffer.AccessToken)
	utils.MakeRequest(updateUrl, "POST", data)

	// Facebook
	data = "text=" + encText.String() + "&now=true&profile_ids[]=" + b.Config.Buffer.FacebookID
	if image != "" {
		data += "&media[photo]=" + imgText.String()
	}
	updateUrl = fmt.Sprintf("https://api.bufferapp.com/1/updates/create.json?access_token=%s", b.Config.Buffer.AccessToken)
	utils.MakeRequest(updateUrl, "POST", data)
}

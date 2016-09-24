package main

import (
	//"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type BufferAPI struct {
	AccessToken string
	TwitterID   string
	FacebookID  string
}

// Access Buffer.com REST API
func (b *BufferAPI) post(text string, image string) {

	text = cleanURL(text)

	encText, err := url.Parse(text)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("(BUFFER) text: %s image: %s\n", encText.String(), image)

	data := "text=" + encText.String() + "&now=true&profile_ids[]=" + b.TwitterID

	if image != "" {
		data += "&media[photo]=" + image
	}

	fmt.Println("====" + data + "====")
	updateUrl := fmt.Sprintf("https://api.bufferapp.com/1/updates/create.json?access_token=%s", b.AccessToken)
	makeRequest(updateUrl, "POST", data)
}

func (b *BufferAPI) postLineup() {
	image := fmt.Sprintf("https://www.scoopanalytics.com/ff/output/%d/%d/%d.png", conf.Season, conf.Week, conf.DKID)

	text := []string{}
	text = append(text, "Its #DFF time! Here is our #EPL #Draftkings lineup for #GW%d")
	text = append(text, "This week's #EPL #Draftkings lineup! #GW%d - Good luck all!")
	text = append(text, "Here's our lineup for today's #GW%d #Draftkings #DFF")
	text = append(text, "Algorithm driven #EPL #DFF lineups - here's today's #GW%d selection")
	text = append(text, "...and here it is, good luck all! #EPL #DFF #DraftKings #GW%d")
	text = append(text, "Here is our #AI selected #GW%d #DraftKings lineup - Good luck!")
	text = append(text, "Our algorithm driven #EPL #DraftKings lineups for #GW%d is here! Good luck!")

	sel := Random(0, len(text)-1)

	cur := fmt.Sprintf(text[sel], conf.Week)

	b.post(cur, image)

}

func (b *BufferAPI) postInjury(inj *Injury, tweet *Tweet) {

	encName, err := url.Parse(inj.Name)
	if err != nil {
		fmt.Println(err)
	}

	encInjury, err := url.Parse(inj.Injury)
	if err != nil {
		fmt.Println(err)
	}

	encTeam, err := url.Parse(inj.Team)
	if err != nil {
		fmt.Println(err)
	}

	// Step 1, Generate and save the image
	url := fmt.Sprintf(conf.RemoteLoc+"/assets/scripts/injury/image.php?data=%s,%s,%s,%s,%s,%d", encName.String(), encInjury.String(), encTeam.String(), inj.Perc, inj.Returns, tweet.ID)
	fmt.Println("Getting " + url)

	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	// Step 2, post to buffer using saved image
	url2 := fmt.Sprintf(conf.RemoteLoc+"/assets/scripts/injury/output/%d.jpg", tweet.ID)
	fmt.Println("Posting " + url2)
	defer response.Body.Close()
	fmt.Println(url)
	if inj.Returns != "" {
		conf.Buffer.post("#FPL Injury news: "+inj.Name+" ("+strings.ToUpper(inj.Team)+") - "+inj.Injury+" - Returns: "+inj.Returns, url2)
	}
}

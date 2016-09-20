package main

import (
	"bytes"
	//"crypto/tls"
	"fmt"
	"io/ioutil"

	"net/http"
	"net/url"
)

var ACCESS_TOKEN = conf.Buffer.AccessToken
var UPDATE_URL = fmt.Sprintf("https://api.bufferapp.com/1/updates/create.json?access_token=%s", ACCESS_TOKEN)
var TWITTER_ID = conf.Buffer.TwitterID

// Access Buffer.com REST API
func postToBuffer() {
	image := fmt.Sprintf("https://www.scoopanalytics.com/ff/output/%d/%d/%d.png", SEASON, WEEK, DKID)
	encImage, _ := url.Parse(image)
	fmt.Println(encImage.String())

	text := []string{}
	text = append(text, "Its #DFF time! Here is our #EPL #Draftkings lineup for #GW%d")
	text = append(text, "This week's #EPL #Draftkings lineup! #GW%d - Good luck all!")
	text = append(text, "Here's our lineup for today's #GW%d #Draftkings #DFF")
	text = append(text, "Algorithm driven #EPL #DFF lineups - here's today's #GW%d selection")
	text = append(text, "...and here it is, good luck all! #EPL #DFF #DraftKings #GW%d")
	text = append(text, "Here is our #AI selected #GW%d #DraftKings lineup - Good luck!")
	text = append(text, "Our algorithm driven #EPL #DraftKings lineups for #GW%d is here! Good luck!")

	sel := Random(0, len(text)-1)

	cur := fmt.Sprintf(text[sel], WEEK)
	fmt.Println(cur)
	encText, _ := url.Parse(cur)
	fmt.Println(encText.String())

	data := "text=" + encText.String() + "&now=true&profile_ids[]=" + TWITTER_ID + "&media[photo]=" + encImage.String()
	fmt.Println(data)

	makeRequest(UPDATE_URL, "POST", data)
}

func makeRequest(url string, method string, vals string) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(vals)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

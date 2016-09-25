package main

/* func postInjury(){

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
				name := utils.GetLastName(inj.Name)
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
}*/

/*
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

*/

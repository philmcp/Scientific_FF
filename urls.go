package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func makeRequest(url string, method string, vals string) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(vals)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	ioutil.ReadAll(resp.Body)
	//	fmt.Println("response Body:", string(body))
}

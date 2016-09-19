package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"time"

	"os"
	"path/filepath"
)

// Generate a random number
func random(min, max int) int {
	//fmt.Printf("%d %d\n", min, max)
	rand.Seed(time.Now().UnixNano())
	ret := rand.Intn(max-min) + min
	return ret
}

func wget(url string, to string) {
	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}

	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	fmt.Println("Success!")
}

func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

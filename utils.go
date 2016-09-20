package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"

	"os"
	"path/filepath"
	"strings"
	"time"
)

// Generate a random number
func Random(min, max int) int {
	//fmt.Printf("%d %d\n", min, max)
	rand.Seed(time.Now().UnixNano())
	ret := rand.Intn(max-min) + min
	return ret
}

func Wget(url string, to string) {
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

func ParseEncoding(str string) string {
	str = strings.ToLower(str)

	str = strings.Replace(str, "è", "e", -1)
	str = strings.Replace(str, "á", "a", -1)
	str = strings.Replace(str, "é", "e", -1)
	str = strings.Replace(str, "ö", "o", -1)
	str = strings.Replace(str, "í", "i", -1)
	str = strings.Replace(str, "à", "a", -1)
	str = strings.Replace(str, "ó", "o", -1)
	str = strings.Replace(str, "ú", "u", -1)
	str = strings.Replace(str, "ü", "u", -1)
	return str
}

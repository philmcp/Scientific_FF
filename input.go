package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	//	"golang.org/x/text/transform"
	"log"
	"os"
	"strconv"
	"strings"
)

type CSVData struct {
	Data [][]string
}

// Some functions to parse the data from file
var enc = charmap.Windows1252

func loadCSV(filename string) (csv CSVData) {
	fmt.Println("Loading CSV: " + filename)
	f, _ := os.Open(filename)

	//	r := transform.NewReader(f, enc.NewDecoder())
	scanner := bufio.NewScanner(f)

	i := 0

	for scanner.Scan() {
		curLine := ParseEncoding(scanner.Text())
		curRow := strings.Split(curLine, ",")

		if len(curRow) > 2 {
			csv.Data = append(csv.Data, curRow)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	return csv
}

func getLastName(str string) string {
	source := ParseEncoding(strings.ToLower(str))

	names := strings.Split(source, " ")
	name := names[len(names)-1]
	return name
}

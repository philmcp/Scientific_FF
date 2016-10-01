package scrape

import (
	"bufio"

	"github.com/philmcp/Scientific_FF/models"
	"github.com/philmcp/Scientific_FF/utils"
	"golang.org/x/text/encoding/charmap"
	"log"
	"os"
	"strings"
)

/* Takes a file and produces a CSVData array */

var enc = charmap.Windows1252

func LoadCSV(filename string) (csv models.CSVData) {
	log.Println("Loading CSV: " + filename)
	f, _ := os.Open(filename)

	//	r := transform.NewReader(f, enc.NewDecoder())
	scanner := bufio.NewScanner(f)

	i := 0

	for scanner.Scan() {
		curLine := utils.ParseEncoding(scanner.Text())
		curRow := strings.Split(curLine, ",")

		if len(curRow) > 2 {
			csv.Data = append(csv.Data, curRow)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	defer f.Close()

	return csv
}

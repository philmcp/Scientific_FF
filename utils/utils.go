package utils

import (
	"math/rand"
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

func GetLastName(str string) string {
	source := ParseEncoding(strings.ToLower(str))

	names := strings.Split(source, " ")
	name := names[len(names)-1]
	return name
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

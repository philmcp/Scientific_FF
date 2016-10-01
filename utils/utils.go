package utils

import (
	"math/rand"
	"strings"
	"time"
)

// Generate a random number
func Random(min, max int) int {
	//log.Printf("%d %d\n", min, max)
	rand.Seed(time.Now().UnixNano())
	ret := rand.Intn(max-min) + min
	return ret
}

func GetDisplayName(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

func DBToStructName(str string) string {
	str = strings.Title(strings.Replace(str, "_", " ", -1))
	str = strings.Replace(str, " ", "", -1)
	return str
}

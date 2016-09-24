package models

import (
	"math/rand"
	"time"
)

// Generate a random number
func random(min, max int) int {
	//fmt.Printf("%d %d\n", min, max)
	rand.Seed(time.Now().UnixNano())
	ret := rand.Intn(max-min) + min
	return ret
}

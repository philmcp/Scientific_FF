package main

import (
	"fmt"
	"time"
)

/* Generate a new optimal linup */
func generate() {
	fmt.Println("Generating a new optimal line up")

	fmt.Printf("Running for Week: %d Season: %d Game: %d\n", conf.Week, conf.Season, conf.DKID)
	time.Sleep(time.Second * 1)

	inputFolder = fmt.Sprintf("input/%d/%d/", conf.Season, conf.Week)
	outputFolder = fmt.Sprintf("output/%d/%d/", conf.Season, conf.Week)

	pool := scrapeData()
	pool.getBestLineup()
	//postToBuffer()
}

func (pool *PlayerPool) getBestLineup() {

	//	Spawn some threads to select random team combinations
	for k := 0.0; k < (conf.Threads); k++ {
		go thread(pool, conf.MinValue+(k*conf.ValueJump))
	}

	// Run for 10 mins
	time.Sleep(90 * time.Second)
}

// Thread to select random team combinations
func thread(pool *PlayerPool, minValue float64) {
	printGap := 10000
	fmt.Printf("Searching for lineups with a min value of %f\n", minValue)
	i := 0
	for {
		pool.randomLineup(minValue)
		i++
		if i%printGap == 0 {
			iter += i
			i = 0
			if iter%(printGap*int(conf.Threads)) == 0 {
				now := time.Now().UnixNano()
				elapsed := float64(now-start) / 1000000000.0
				speed := float64(iter) / elapsed
				fmt.Printf("Thread: %d after %fs - %f/s\n", iter, elapsed, speed)
			}

		}
	}
}

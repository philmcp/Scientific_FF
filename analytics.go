package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"strings"
	"time"
)

func analyticsTransfersIn(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	loadFPLData()
	players := db.GetMostTransferredIn()

	text := fmt.Sprintf("This week's most SWAPPED IN players: ")

	for _, cur := range *players {
		temp := humanize.SI(cur.FPL.TransfersInEvent, "")
		spl := strings.Split(temp, ".")
		text += fmt.Sprintf("%s %s, ", cur.GetDisplayName(), spl[0]+"k")
	}

	text = strings.Trim(text, ", ") + " #FPL"
	fmt.Println(text)
	buffer.Post(text, "")
}

func analyticsTransfersOut(day string) {
	if time.Now().Weekday().String() != day {
		return
	}
	loadFPLData()
	players := db.GetMostTransferredOut()

	text := fmt.Sprintf("This week's most SWAPPED OUT players: ")

	for _, cur := range *players {
		temp := humanize.SI(cur.FPL.TransfersOutEvent, "")
		spl := strings.Split(temp, ".")
		text += fmt.Sprintf("%s %s, ", cur.GetDisplayName(), spl[0]+"k")
	}

	text = strings.Trim(text, ", ") + " #FPL"
	fmt.Println(text)
	buffer.Post(text, "")
}

func analyticsInForm(day string) {
	if time.Now().Weekday().String() != day {
		return
	}

	loadFPLData()
	players := db.GetInForm()

	text := fmt.Sprintf("Most IN FORM players: ")

	for _, cur := range *players {
		text += fmt.Sprintf("%s %.1f, ", cur.GetDisplayName(), cur.FPL.Form)
	}

	text = strings.Trim(text, ", ") + " #FPL"

	buffer.Post(text, "")
}

package models

import (
	"github.com/dustin/go-humanize"
	"strings"
	"time"
)

type Injury struct {
	Name    string
	Injury  string
	Returns string
	Team    string
	Perc    string
}

func NewInjury(text string) Injury {
	ret := Injury{}
	text = strings.Replace(text, "No Return Date", "Expected Return: Unknown", -1)

	spl := strings.Split(text, ": ")

	name_spl := strings.Split(spl[1], " - ")

	ret.Name = strings.TrimSpace(name_spl[0])

	inj_spl := strings.Split(strings.TrimSpace(name_spl[1]), "#")
	ret.Injury = strings.TrimSpace(strings.Replace(inj_spl[0], "Expected Return", "", -1))

	if strings.Count(text, "#") > 1 {
		team_spl := strings.Split(text, "#")
		team_spl2 := strings.Split(team_spl[2], " ")
		ret.Team = strings.ToLower(team_spl2[0])
	} else {
		ret.Team = ""
	}

	date := strings.TrimSpace(strings.Replace(spl[2], "Status", "", -1))

	if date == "" || date == "Unknown" {
		ret.Returns = "Return date unknown"
	} else {
		layout := "02-01-2006"
		t, _ := time.Parse(layout, date)
		returnDate := humanize.Time(t)
		if strings.Contains(returnDate, "hours") {
			ret.Returns = "Returns tomorrow (" + t.Format("02 Jan") + ")"
		} else {
			ret.Returns = "Returns " + returnDate + " (" + t.Format("02 Jan") + ")"
		}

	}

	perc_spl := strings.Split(spl[3], " ")
	if strings.Contains(perc_spl[0], "%") {
		ret.Perc = strings.Replace(strings.TrimSpace(perc_spl[0]), "%", "", -1)
	} else {
		ret.Perc = "0"
	}

	return ret
}

package models

type CSVData struct {
	Data [][]string
}

type Lineup struct {
	Team       PlayerPool
	Projection float64
	Wage       float64
}

package models

type Lineup struct {
	Team       PlayerPool
	Projection float64
	Wage       float64
	NumTeams   int
}

type CSVData struct {
	Data [][]string
}

type Data struct {
	FPL  PlayerList
	FFS  PlayerList
	Roto PlayerList
	DK   PlayerList
}

type Configuration struct {
	RemoteLoc    string
	OutputFolder string
	Database     struct {
		Host     string
		Port     int
		DB       string
		User     string
		Password string
	}
	FFScout struct {
		Username string
		Password string
	}
	Buffer struct {
		AccessToken string
		TwitterID   string
		FacebookID  string
	}
	Twitter struct {
		AppKey    string
		AppSecret string
	}
	MinNumTeams int
	MaxWage     float64
	Formation   map[string]int
	Threads     float64
	ValueJump   float64
	MinValue    float64
	Season      int
	Week        int
	DKID        int
	DKName      string
}

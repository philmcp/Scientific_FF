package draw

import (
	"fmt"
	"github.com/philmcp/Scientific_FF/models"
)

type Draw struct {
	Config *models.Configuration
	Folder string
}

func NewDraw(config *models.Configuration) *Draw {
	return &Draw{
		Config: config,
		Folder: fmt.Sprintf("assets/data/generate/output/%d/%d/", config.Season, config.Week),
	}
}

package draw

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	_ "github.com/golang/freetype/truetype"
	"github.com/leekchan/accounting"
	"github.com/philmcp/Scientific_FF/models"
	_ "golang.org/x/image/font"
	"image"
	"image/png"
	//	"image/color"
	"log"
	"os"
)

func (d *Draw) DrawTeam(l *models.Generated) {
	// Decode the JPEG data. If reading from file, create a reader with

	pixels := map[string][]PixelPos{
		"gk": []PixelPos{PixelPos{616, 876}},
		"d":  []PixelPos{PixelPos{347, 694}, PixelPos{855, 694}},
		"m":  []PixelPos{PixelPos{404, 466}, PixelPos{813, 466}},
		"f":  []PixelPos{PixelPos{480, 257}, PixelPos{785, 257}},
		"u":  []PixelPos{PixelPos{1210, 144}},
	}

	pixelsSalary := PixelPos{1267, 400}
	pixelsWeek := PixelPos{636, 83}
	pixelsGames := PixelPos{640, 130}

	//	d.resetOutput()
	flag.Parse()
	rgba := loadImage("assets/images/templates/pitch.png")

	blackFont := loadFont(image.Black, 45, *fontDIN)
	blackFont.SetClip(rgba.Bounds())
	blackFont.SetDst(rgba)

	whiteFont := loadFont(image.White, 20, *fontDIN)
	whiteFont.SetClip(rgba.Bounds())
	whiteFont.SetDst(rgba)

	ppgFont := loadFont(image.Black, 16, *fontDIN)
	ppgFont.SetClip(rgba.Bounds())
	ppgFont.SetDst(rgba)

	weekFont := loadFont(image.White, 90, *fontDIN)
	weekFont.SetClip(rgba.Bounds())
	weekFont.SetDst(rgba)

	gamesFont := loadFont(image.White, 45, *fontDIN)
	gamesFont.SetClip(rgba.Bounds())
	gamesFont.SetDst(rgba)

	redFont := loadFont(image.White, 80, *fontDIN)
	redFont.SetClip(rgba.Bounds())
	redFont.SetDst(rgba)

	salary := 0.0

	// Draw the team
	for pos, players := range l.Team {
		i := 0
		for _, player := range players {

			playerCoor := pixels[pos][i]
			// Player name
			pt := freetype.Pt(playerCoor.X+7, playerCoor.Y+47)
			_, err := blackFont.DrawString(player.GetDisplayName(), pt)
			if err != nil {
				log.Println(err)
				return
			}

			salary += player.Wage
			teamName := PixelPos{}
			ppg := PixelPos{}

			if pos == "gk" {
				teamName.X = 100
				teamName.Y = 85
				ppg.X = 228
				ppg.Y = 55
			} else if pos == "d" {
				playerCoor.X += 100
				playerCoor.Y += 60
				teamName.X = 13
				teamName.Y = 25
				ppg.X = 149
				ppg.Y = 1
			} else if pos == "m" {
				playerCoor.X += 90
				playerCoor.Y += 57
				teamName.X = 16
				teamName.Y = 28
				ppg.X = 145
				ppg.Y = 1

			} else if pos == "f" {
				playerCoor.X += 80
				playerCoor.Y += 57
				teamName.X = 3
				teamName.Y = 20
				ppg.X = 100
				ppg.Y = -4
			} else if pos == "u" {
				playerCoor.X += 62
				playerCoor.Y += 57
				teamName.X = 18
				teamName.Y = 27
				ppg.X = 118
				ppg.Y = 1

			}

			// Team name
			pt = freetype.Pt(playerCoor.X+teamName.X, playerCoor.Y+teamName.Y)
			_, err = whiteFont.DrawString(player.GetGame(), pt)
			if err != nil {
				log.Println(err)
				return
			}

			// PPG - HARD CODED
			pt = freetype.Pt(playerCoor.X+ppg.X, playerCoor.Y+ppg.Y)
			_, err = ppgFont.DrawString(fmt.Sprintf("PPG: %.1f", player.AvgPointsPerGame), pt)
			if err != nil {
				log.Println(err)
				return
			}

			i++
		}
	}

	// Draw salary
	pt := freetype.Pt(pixelsSalary.X, pixelsSalary.Y)
	form := accounting.Accounting{Symbol: "£", Precision: 0}
	_, err := redFont.DrawString(form.FormatMoney(salary), pt)

	if err != nil {
		log.Println(err)
		return
	}

	// Draw Week
	pt = freetype.Pt(pixelsWeek.X, pixelsWeek.Y)
	_, err = weekFont.DrawString(fmt.Sprintf("Week %d", d.Config.Week), pt)

	if err != nil {
		log.Println(err)
		return
	}

	// Draw Games
	pt = freetype.Pt(pixelsGames.X, pixelsGames.Y)
	_, err = gamesFont.DrawString(d.Config.DKName, pt)

	if err != nil {
		log.Println(err)
		return
	}

	// Save that RGBA image to disk.
	outputLoc := d.Config.OutputFolder + "lineup.png"
	//	log.Println("Drawing to " + outputLoc)
	outFile, err := os.Create(outputLoc)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Fatal(err)
	}
	err = b.Flush()
	if err != nil {
		log.Fatal(err)
	}

	//log.Println("Wrote out.png OK.")
}

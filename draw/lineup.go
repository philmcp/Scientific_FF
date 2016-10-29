package draw

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	_ "github.com/golang/freetype/truetype"
	"github.com/philmcp/Scientific_FF/models"

	"github.com/philmcp/Scientific_FF/utils"
	_ "golang.org/x/image/font"

	"image"
	"image/draw"
	"image/png"
	"time"
	//	"image/color"
	"log"
	"os"
)

func (d *Draw) DrawLineup(l *models.Lineup, id int64) {
	log.Println("Drawing lineup")

	pixelsHome := PixelPos{123, 206}
	pixelsAway := PixelPos{125, 175}
	pixelsTeam := PixelPos{550, 109}
	playerDif := 31
	//	pixelsReturns := PixelPos{640, 130}

	flag.Parse()
	rgba := loadImage("assets/images/templates/lineup.png")

	awayFont := loadFont(image.White, 35, *fontSourceSansLight)
	awayFont.SetClip(rgba.Bounds())
	awayFont.SetDst(rgba)

	homeFont := loadFont(image.White, 85, *fontDIN)
	homeFont.SetClip(rgba.Bounds())
	homeFont.SetDst(rgba)

	teamFont := loadFont(image.White, 24, *fontSourceSansRegular)
	teamFont.SetClip(rgba.Bounds())
	teamFont.SetDst(rgba)

	// Draw home
	home := freetype.Pt(pixelsHome.X, pixelsHome.Y)
	_, err := homeFont.DrawString(l.Team, home)

	if err != nil {
		log.Println(err)
		return
	}

	// Draw away
	subtext := "vs " + l.OppTeam
	if subtext == "vs " {
		subtext = "Today's lineup"
	}

	subtext += " (" + time.Now().Format("02 Jan") + ")"

	away := freetype.Pt(pixelsAway.X, pixelsAway.Y+80)
	_, err = awayFont.DrawString(subtext, away)

	if err != nil {
		log.Println(err)
		return
	}

	// Team

	for i, player := range l.Players {
		cur := freetype.Pt(pixelsTeam.X, pixelsTeam.Y+(playerDif*i))
		_, err := teamFont.DrawString(player, cur)

		if err != nil {
			log.Println(err)
			return
		}
	}

	// Draw home logo
	homeSlug := utils.GenerateSlug(l.Team)
	fmt.Println("assets/images/logos/" + homeSlug + ".png")
	fImg2, err := os.Open("assets/images/logos/" + homeSlug + ".png")
	if err == nil {
		if l.Team != "" {
			defer fImg2.Close()
			img2, _, _ := image.Decode(fImg2)

			draw.Draw(rgba, rgba.Bounds(), img2, image.Point{-120, -276}, draw.Over)
		} else {
			log.Println(err)
		}

		// Draw away logo

		awaySlug := utils.GenerateSlug(l.OppTeam)
		fmt.Println("assets/images/logos/" + awaySlug + ".png")
		fImg2, err = os.Open("assets/images/logos/" + awaySlug + ".png")
		if l.Team != "" && err == nil {
			defer fImg2.Close()
			img2, _, _ := image.Decode(fImg2)

			draw.Draw(rgba, rgba.Bounds(), img2, image.Point{-230, -276}, draw.Over)
		} else {
			log.Println(err)
		}
	}

	// Save that RGBA image to disk.
	outputLoc := d.Config.OutputFolder + fmt.Sprintf("/lineups/%d.png", id)
	log.Println("Drawing to " + outputLoc)

	os.Chmod(outputLoc, 0775)
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

}

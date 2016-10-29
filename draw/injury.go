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

	//	"image/color"
	"log"
	"os"
)

func (d *Draw) DrawInjury(l *models.Injury, id int64) {
	log.Println("Drawing injury")

	pixelsName := PixelPos{124, 312}
	pixelsInjury := PixelPos{126, 360}
	//	pixelsReturns := PixelPos{640, 130}

	flag.Parse()
	rgba := loadImage("assets/images/templates/injury.png")

	injuryFont := loadFont(image.White, 30, *fontSourceSansRegular)
	injuryFont.SetClip(rgba.Bounds())
	injuryFont.SetDst(rgba)

	nameFont := loadFont(image.White, 105, *fontDIN)
	nameFont.SetClip(rgba.Bounds())
	nameFont.SetDst(rgba)

	// Draw name
	name := freetype.Pt(pixelsName.X, pixelsName.Y)
	_, err := nameFont.DrawString(utils.GetDisplayName(l.Name), name)

	if err != nil {
		log.Println(err)
		return
	}

	// Draw Injury

	retText := l.Injury + " · " + l.Returns

	injury := freetype.Pt(pixelsInjury.X, pixelsInjury.Y)
	_, err = injuryFont.DrawString(retText, injury)

	if err != nil {
		log.Println(err)
		return
	}

	// Draw logo
	if l.Team != "" {
		fmt.Println("assets/images/logos/" + l.Team + ".png")
		fImg2, _ := os.Open("assets/images/logos/" + l.Team + ".png")
		defer fImg2.Close()
		img2, _, _ := image.Decode(fImg2)

		draw.Draw(rgba, rgba.Bounds(), img2, image.Point{-790, -105}, draw.Over)
	}

	// Save that RGBA image to disk.
	outputLoc := d.Config.OutputFolder + fmt.Sprintf("/injuries/%d.png", id)
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

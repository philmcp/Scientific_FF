package main

/*
import (
	"bufio"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	_ "github.com/golang/freetype/truetype"
	"github.com/leekchan/accounting"
	_ "golang.org/x/image/font"
	"image"
	//	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type PixelPos struct {
	X int
	Y int
}

var (
	dpi        = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile   = flag.String("fontfile", "fonts/DIN Condensed Bold.ttf", "filename of the ttf font")
	hinting    = flag.String("hinting", "none", "none | full")
	playerSize = flag.Float64("playerSize", 48, "font size in points")
	salarySize = flag.Float64("salarySize", 80, "font size in points")
	weekSize   = flag.Float64("weekSize", 90, "font size in points")
	gamesSize  = flag.Float64("gamesSize", 45, "font size in points")
	teamSize   = flag.Float64("teamSize", 27, "font size in points")
	spacing    = flag.Float64("spacing", 1.5, "line spacing (e.g. 2 means double spaced)")
	wonb       = flag.Bool("whiteonblack", false, "white text on a black background")
	pixels     = map[string][]PixelPos{
		"gk": []PixelPos{PixelPos{616, 876}},
		"d":  []PixelPos{PixelPos{347, 694}, PixelPos{855, 694}},
		"m":  []PixelPos{PixelPos{404, 466}, PixelPos{813, 466}},
		"f":  []PixelPos{PixelPos{480, 257}, PixelPos{785, 257}},
		"u":  []PixelPos{PixelPos{1210, 144}},
	}

	pixelsSalary = PixelPos{1267, 400}
	pixelsWeek   = PixelPos{636, 83}
	pixelsGames  = PixelPos{640, 130}
)

func resetOutput() {
	fmt.Println("Reseting " + outputFolder)
	//	RemoveContents(outputFolder)
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		os.Mkdir(outputFolder, 0777)
	}
}

func loadImage() *image.RGBA {
	reader, err := os.Open("output/pitch_template.png")

	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	src, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := src.Bounds()
	rgba := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(rgba, rgba.Bounds(), src, rgba.Rect.Min, draw.Src)

	return rgba
}

func loadFont(color *image.Uniform, size *float64) *freetype.Context {
	fontBytes, err := ioutil.ReadFile(*fontfile)
	if err != nil {
		log.Println(err)
		return nil
	}
	f, err := freetype.ParseFont(fontBytes)

	if err != nil {
		log.Println(err)
		return nil
	}

	c := freetype.NewContext()
	c.SetDPI(*dpi)
	c.SetFont(f)
	c.SetFontSize(*size)
	c.SetSrc(color)

	return c
}

func (t *Team) drawTeam() {
	// Decode the JPEG data. If reading from file, create a reader with
	resetOutput()
	flag.Parse()
	rgba := loadImage()

	blackFont := loadFont(image.Black, playerSize)
	blackFont.SetClip(rgba.Bounds())
	blackFont.SetDst(rgba)

	whiteFont := loadFont(image.White, teamSize)
	whiteFont.SetClip(rgba.Bounds())
	whiteFont.SetDst(rgba)

	weekFont := loadFont(image.White, weekSize)
	weekFont.SetClip(rgba.Bounds())
	weekFont.SetDst(rgba)

	gamesFont := loadFont(image.White, gamesSize)
	gamesFont.SetClip(rgba.Bounds())
	gamesFont.SetDst(rgba)

	redFont := loadFont( /*image.Uniform.RGBA(50, 50, 50, 1)  image.White, salarySize)
	redFont.SetClip(rgba.Bounds())
	redFont.SetDst(rgba)

	salary := 0.0

	// Draw the team
	for pos, players := range t.Players {
		i := 0
		for _, player := range players {

			playerCoor := pixels[pos][i]
			// Player name
			pt := freetype.Pt(playerCoor.X+7, playerCoor.Y+int(blackFont.PointToFixed(*playerSize)>>6))
			_, err := blackFont.DrawString(player.getDisplayName(), pt)
			if err != nil {
				log.Println(err)
				return
			}

			salary += player.Wage
			teamName := PixelPos{}

			if pos == "gk" {
				teamName.X = 118
				teamName.Y = 89
			} else if pos == "d" {
				playerCoor.X += 112
				playerCoor.Y += 60
				teamName.X = 18
				teamName.Y = 32
			} else if pos == "m" {
				playerCoor.X += 101
				playerCoor.Y += 57
				teamName.X = 19
				teamName.Y = 32

			} else if pos == "f" {
				playerCoor.X += 99
				playerCoor.Y += 57
				teamName.X = 2
				teamName.Y = 24
			} else if pos == "u" {
				playerCoor.X += 82
				playerCoor.Y += 57
				teamName.X = 15
				teamName.Y = 30

			}

			// Team name
			pt = freetype.Pt(playerCoor.X+teamName.X, playerCoor.Y+teamName.Y)
			_, err = whiteFont.DrawString(strings.ToUpper(player.Team), pt)
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
	_, err = weekFont.DrawString(fmt.Sprintf("Week %d", gameWeek), pt)

	if err != nil {
		log.Println(err)
		return
	}

	// Draw Games
	pt = freetype.Pt(pixelsGames.X, pixelsGames.Y)
	_, err = gamesFont.DrawString(DKName, pt)

	if err != nil {
		log.Println(err)
		return
	}

	// Save that RGBA image to disk.
	outputLoc := fmt.Sprintf("output/%d/%d/%d.png", season, gameWeek, DKID)
	fmt.Println("Drawing to " + outputLoc)
	outFile, err := os.Create(outputLoc)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println("Wrote out.png OK.")
}

*/

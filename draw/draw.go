package draw

import (
	"flag"
	"github.com/golang/freetype"
	_ "github.com/golang/freetype/truetype"
	"github.com/philmcp/Scientific_FF/models"

	"image"
	"image/draw"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
)

var (
	dpi                 = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontDIN             = flag.String("din", "assets/fonts/DIN Condensed Bold.ttf", "filename of the ttf font")
	fontSourceSansLight = flag.String("source sans", "assets/fonts/source-sans-pro.light.ttf", "filename of the ttf font")
)

type PixelPos struct {
	X int
	Y int
}

type Draw struct {
	Config *models.Configuration
}

func NewDraw(config *models.Configuration) *Draw {
	return &Draw{
		Config: config,
	}
}

func loadImage(path string) *image.RGBA {
	reader, err := os.Open(path)

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

func loadFont(color *image.Uniform, size float64, font string) *freetype.Context {
	fontBytes, err := ioutil.ReadFile(font)

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
	c.SetFontSize(size)
	c.SetSrc(color)

	return c
}

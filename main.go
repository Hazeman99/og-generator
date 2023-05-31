package main

import (
	"errors"
	"image/color"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
)

func main() {
	width := 1200
	height := 630

	dc := gg.NewContext(width, height)

	if err := loadBackgroundImage(dc); err != nil {
		panic(err)
	}

	addOverlay(dc)

	fontPath := filepath.Join("fonts", "OpenSans-Regular.ttf")
	if err := dc.LoadFontFace(fontPath, 50); err != nil {
		panic(errors.New("unable to load font"))
	}
	dc.SetColor(color.Black)

	s := os.Args[1]
	addTitle(dc, s)

	if err := dc.LoadFontFace(fontPath, 45); err != nil {
		panic(errors.New("unable to load font"))
	}

	website := "https://hazemhadi.com"
	addWebsite(dc, website)

	dc.SavePNG("./output.png")
}

func loadBackgroundImage(dc *gg.Context) (err error) {
	backgroundImage, err := gg.LoadImage("./OG-Portfolio-BaseImage.png")
	if err != nil {
		return err
	}

	dc.DrawImage(backgroundImage, 0, 0)

	return nil
}

func addOverlay(dc *gg.Context) {
	margin := 20.0
	x := margin
	y := margin
	w := float64(dc.Width()) - (2.0 * margin)
	h := float64(dc.Height()) - (2.0 * margin)
	dc.SetColor(color.RGBA{255, 255, 255, 255})
	dc.DrawRectangle(x, y, w, h)
	dc.Fill()
}

func addTitle(dc *gg.Context, title string) {
	textRightMargin := 60.0
	textTopMargin := 90.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
}

func addWebsite(dc *gg.Context, website string) {
	textColor := color.Black
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)

	_, textHeight := dc.MeasureString(website)
	marginY := 30.0
	x := 70.0
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(website, x, y)
}

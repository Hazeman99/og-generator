package main

import (
	"errors"
	"flag"
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"github.com/fogleman/gg"
)

func main() {
	// Setup the canvas
	width := 1200
	height := 630

	dc := gg.NewContext(width, height)

	if err := loadBackgroundImage(dc); err != nil {
		panic(err)
	}

	addOverlay(dc)

	fontPath := filepath.Join("fonts", "OpenSans-Bold.ttf")
	if err := dc.LoadFontFace(fontPath, 50); err != nil {
		panic(errors.New("unable to load font"))
	}
	dc.SetColor(color.Black)

	// Parse args from the command line
	title := flag.String("title", "", "title of the article")
	description := flag.String("description", "", "description of the article")
	author := flag.String("author", "", "author of the article")
	minutes := flag.Int("minutes", 0, "minutes to read the article")
	website := flag.String("website", "", "website of the article")

	flag.Parse()

	addTitle(dc, *title)

	fontPath = filepath.Join("fonts", "OpenSans-Regular.ttf")
	if err := dc.LoadFontFace(fontPath, 34); err != nil {
		panic(errors.New("unable to load font"))
	}

	// set the color to gray
	dc.SetHexColor("#4a4a4a")
	addDescription(dc, *description, len(*title))

	if err := dc.LoadFontFace(fontPath, 38); err != nil {
		panic(errors.New("unable to load font"))
	}

	addMinutesAndAuthor(dc, *minutes, *author)

	addWebsite(dc, *website)

	// if icon does not exist, it will not be added
	_, err := os.Stat("./icon.png")
	if err != nil {
		fmt.Println("Error getting file. Skipping...")
	}

	if err == nil {
		err = addIcon(dc, "./icon.png")
		if err != nil {
			panic(err)
		}
	}

	err = dc.SavePNG("./output.png")
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully generated image. Check output.png")
}

func loadBackgroundImage(dc *gg.Context) (err error) {
	backgroundImage, err := gg.LoadImage("./background.png")
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
	textTopMargin := 100.0
	x := textRightMargin
	y := textTopMargin
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin
	dc.DrawStringWrapped(title, x, y, 0, 0, maxWidth, 1.5, gg.AlignLeft)
}

func addDescription(dc *gg.Context, description string, titleLength int) {
	// take into account that it depends on Title for where it starts
	textRightMargin := 60.0
	textTopMargin := 100.0
	x := textRightMargin
	// split string into letters and count the number of lines
	if titleLength > 50 {
		textTopMargin = 140.0
	}

	y := textTopMargin + 80
	maxWidth := float64(dc.Width()) - textRightMargin - textRightMargin

	if len(description) > 197 {
		description = description[:197] + "..."
	}

	dc.DrawStringWrapped(description, x, y, 0, 0, maxWidth, 1.8, gg.AlignLeft)
}

func addMinutesAndAuthor(dc *gg.Context, minutes int, author string) {
	textColor := color.Black
	r, g, b, _ := textColor.RGBA()
	mutedColor := color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(200),
	}
	dc.SetColor(mutedColor)

	text := fmt.Sprintf("%d min read • %s", minutes, author)

	_, textHeight := dc.MeasureString(text)
	marginY := 110.0
	x := 70.0
	y := float64(dc.Height()) - textHeight - marginY
	dc.DrawString(text, x, y)
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

func addIcon(dc *gg.Context, iconPath string) (err error) {
	icon, err := gg.LoadImage(iconPath)
	if err != nil {
		return err
	}

	dc.DrawImage(icon, dc.Width()-100, dc.Height()-100)

	return nil
}

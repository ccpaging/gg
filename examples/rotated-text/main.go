package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/ccpaging/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func savePNG(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("could not encode PNG to %q: %w", path, err)
	}

	return file.Close()
}

func main() {
	const S = 400
	dc := gg.NewDeviceContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	text := "Hello, world!"
	w, h := dc.MeasureString(text)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 180, w, h)
	dc.Stroke()
	dc.DrawStringAnchored(text, 100, 180, 0.0, 0.0)
	savePNG("out.png", dc.Image())
}

package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/ccpaging/gg"
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
	img, err := gg.LoadImage("../baboon.png")
	if err != nil {
		panic(err)
	}
	pattern := gg.NewSurfacePattern(img, gg.RepeatBoth)
	dc := gg.NewDeviceContext(600, 600)
	dc.MoveTo(20, 20)
	dc.LineTo(590, 20)
	dc.LineTo(590, 590)
	dc.LineTo(20, 590)
	dc.ClosePath()
	dc.SetFillStyle(pattern)
	dc.Fill()
	savePNG("out.png", dc.Image())
}

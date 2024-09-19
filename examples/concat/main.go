package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
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
	im1, err := gg.LoadImage("../baboon.png")
	if err != nil {
		panic(err)
	}

	im2, err := gg.LoadImage("../gopher.png")
	if err != nil {
		panic(err)
	}

	s1 := im1.Bounds().Size()
	s2 := im2.Bounds().Size()

	width := int(math.Max(float64(s1.X), float64(s2.X)))
	height := s1.Y + s2.Y

	dc := gg.NewDeviceContext(width, height)
	dc.DrawImage(im1, 0, 0)
	dc.DrawImage(im2, 0, s1.Y)
	savePNG("out.png", dc.Image())
}

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
	const NX = 4
	const NY = 3
	img, err := gg.LoadImage("../gopher.png")
	if err != nil {
		panic(err)
	}
	w := img.Bounds().Size().X
	h := img.Bounds().Size().Y
	dc := gg.NewDeviceContext(w*NX, h*NY)
	for y := 0; y < NY; y++ {
		for x := 0; x < NX; x++ {
			dc.DrawImage(img, x*w, y*h)
		}
	}
	savePNG("out.png", dc.Image())
}

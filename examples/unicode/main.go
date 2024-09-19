package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/ccpaging/gg"
	"github.com/flopp/go-findfont"
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
	const S = 4096 * 2
	const T = 16 * 2
	const F = 28
	dc := gg.NewDeviceContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fontPath, err := findfont.Find("Xolonium.ttf")
	if err != nil {
		panic(err)
	}
	if err := dc.LoadFontFace(fontPath, F); err != nil {
		panic(err)
	}
	for r := 0; r < 256; r++ {
		for c := 0; c < 256; c++ {
			i := r*256 + c
			x := float64(c*T) + T/2
			y := float64(r*T) + T/2
			dc.DrawStringAnchored(string(rune(i)), x, y, 0.5, 0.5)
		}
	}
	savePNG("out.png", dc.Image())
}

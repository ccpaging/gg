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
	const S = 1024
	dc := gg.NewDeviceContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	fontPath, err := findfont.Find("Impact.ttf")
	if err != nil {
		panic(err)
	}
	if err := dc.LoadFontFace(fontPath, 96); err != nil {
		panic(err)
	}
	dc.SetRGB(0, 0, 0)
	s := "ONE DOES NOT SIMPLY"
	n := 6 // "stroke" size
	for dy := -n; dy <= n; dy++ {
		for dx := -n; dx <= n; dx++ {
			if dx*dx+dy*dy >= n*n {
				// give it rounded corners
				continue
			}
			x := S/2 + float64(dx)
			y := S/2 + float64(dy)
			dc.DrawStringAnchored(s, x, y, 0.5, 0.5)
		}
	}
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(s, S/2, S/2, 0.5, 0.5)
	savePNG("out.png", dc.Image())
}

package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
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
	dc := gg.NewDeviceContext(1000, 1000)
	for j := 0; j < 10; j++ {
		for i := 0; i < 10; i++ {
			x := float64(i)*100 + 50
			y := float64(j)*100 + 50
			a1 := rand.Float64() * 2 * math.Pi
			a2 := a1 + rand.Float64()*math.Pi + math.Pi/2
			dc.DrawArc(x, y, 40, a1, a2)
			// dc.ClosePath()
		}
	}
	dc.SetRGB(0, 0, 0)
	dc.FillPreserve()
	dc.SetRGB(1, 1, 1)
	dc.SetLineWidth(8)
	dc.StrokePreserve()
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(4)
	dc.StrokePreserve()
	savePNG("out.png", dc.Image())
}

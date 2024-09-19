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
	const S = 1024
	const N = 2048
	dc := gg.NewDeviceContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	for i := 0; i <= N; i++ {
		t := float64(i) / N
		d := t*S*0.4 + 10
		a := t * math.Pi * 2 * 20
		x := S/2 + math.Cos(a)*d
		y := S/2 + math.Sin(a)*d
		r := t * 8
		dc.DrawCircle(x, y, r)
	}
	dc.Fill()
	savePNG("out.png", dc.Image())
}

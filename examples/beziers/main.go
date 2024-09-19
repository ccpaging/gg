package main

import (
	"fmt"
	"image"
	"image/png"
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

func random() float64 {
	return rand.Float64()*2 - 1
}

func point() (x, y float64) {
	return random(), random()
}

func drawCurve(dc *gg.DeviceContext) {
	dc.SetRGBA(0, 0, 0, 0.1)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(12)
	dc.Stroke()
}

func drawPoints(dc *gg.DeviceContext) {
	dc.SetRGBA(1, 0, 0, 0.5)
	dc.SetLineWidth(2)
	dc.Stroke()
}

func randomQuadratic(dc *gg.DeviceContext) {
	x0, y0 := point()
	x1, y1 := point()
	x2, y2 := point()
	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	drawCurve(dc)
	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	drawPoints(dc)
}

func randomCubic(dc *gg.DeviceContext) {
	x0, y0 := point()
	x1, y1 := point()
	x2, y2 := point()
	x3, y3 := point()
	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	drawCurve(dc)
	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	drawPoints(dc)
}

func main() {
	const (
		S = 256
		W = 8
		H = 8
	)
	dc := gg.NewDeviceContext(S*W, S*H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for j := 0; j < H; j++ {
		for i := 0; i < W; i++ {
			x := float64(i)*S + S/2
			y := float64(j)*S + S/2
			dc.Push()
			dc.Translate(x, y)
			dc.Scale(S/2, S/2)
			if j%2 == 0 {
				randomCubic(dc)
			} else {
				randomQuadratic(dc)
			}
			dc.Pop()
		}
	}
	savePNG("out.png", dc.Image())
}

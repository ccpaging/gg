package main

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"

	"github.com/ccpaging/gg"
)

type Point struct {
	X, Y float64
}

func Polygon(n int, x, y, r float64) []Point {
	result := make([]Point, n)
	for i := 0; i < n; i++ {
		a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
		result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
	}
	return result
}

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
	n := 5
	points := Polygon(n, 512, 512, 400)
	dc := gg.NewDeviceContext(1024, 1024)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()
	savePNG("out.png", dc.Image())
}

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
	const W = 400
	const H = 500
	img, err := gg.LoadImage("../gopher.png")
	if err != nil {
		panic(err)
	}
	iw, ih := img.Bounds().Dx(), img.Bounds().Dy()
	dc := gg.NewDeviceContext(W, H)
	// draw outline
	dc.SetHexColor("#ff0000")
	dc.SetLineWidth(1)
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Stroke()
	// draw full image
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.DrawRectangle(100, 210, float64(iw), float64(ih))
	dc.Stroke()
	dc.DrawImage(img, 100, 210)
	// draw image with current matrix applied
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 0, float64(iw), float64(ih)/2+20.0)
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImageAnchored(img, 100, 0, 0.0, 0.0)
	savePNG("out.png", dc.Image())
}

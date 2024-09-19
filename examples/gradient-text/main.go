package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/ccpaging/gg"
	"github.com/flopp/go-findfont"
)

const (
	W = 1024
	H = 512
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
	dc := gg.NewDeviceContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)
	fontPath, err := findfont.Find("Impact.ttf")
	if err != nil {
		panic(err)
	}
	dc.LoadFontFace(fontPath, 128)
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the DeviceContext as an alpha mask
	mask := dc.AsMask()

	// clear the DeviceContext
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	// using the mask, fill the DeviceContext with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	savePNG("out.png", dc.Image())
}

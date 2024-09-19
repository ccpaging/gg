package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
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
	img, err := gg.LoadImage("../baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewDeviceContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(img, 0, 0)
	savePNG("out.png", dc.Image())
}

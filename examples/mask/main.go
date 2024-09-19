package main

import (
	"log"

	"github.com/ccpaging/gg"
)

func main() {
	img, err := gg.LoadImage("../baboon.png")
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewDeviceContext(512, 512)
	dc.DrawRoundedRectangle(0, 0, 512, 512, 64)
	dc.Clip()
	dc.DrawImage(img, 0, 0)
	dc.SavePNG("out.png")
}

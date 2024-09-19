package main

import (
	"github.com/ccpaging/gg"
	"github.com/flopp/go-findfont"
)

func main() {
	const S = 1024
	dc := gg.NewDeviceContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	fontPath, err := findfont.Find("Arial.ttf")
	if err != nil {
		panic(err)
	}
	if err := dc.LoadFontFace(fontPath, 96); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Hello, world!", S/2, S/2, 0.5, 0.5)
	dc.SavePNG("out.png")
}

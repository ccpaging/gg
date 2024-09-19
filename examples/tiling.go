package main

import "github.com/ccpaging/gg"

func main() {
	const NX = 4
	const NY = 3
	img, err := gg.LoadPNG("examples/gopher.png")
	if err != nil {
		panic(err)
	}
	w := img.Bounds().Size().X
	h := img.Bounds().Size().Y
	dc := gg.NewDeviceContext(w*NX, h*NY)
	for y := 0; y < NY; y++ {
		for x := 0; x < NX; x++ {
			dc.DrawImage(img, x*w, y*h)
		}
	}
	dc.SavePNG("out.png")
}

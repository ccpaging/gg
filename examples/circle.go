package main

import "github.com/ccpaging/gg"

func main() {
	dc := gg.NewDeviceContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG("out.png")
}

package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/ccpaging/gg"
	"github.com/flopp/go-findfont"
)

const TEXT = "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."

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
	const W = 1024
	const H = 1024
	const P = 16
	dc := gg.NewDeviceContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawLine(W/2, 0, W/2, H)
	dc.DrawLine(0, H/2, W, H/2)
	dc.DrawRectangle(P, P, W-P-P, H-P-P)
	dc.SetRGBA(0, 0, 1, 0.25)
	dc.SetLineWidth(3)
	dc.Stroke()
	dc.SetRGB(0, 0, 0)
	fontPath, err := findfont.Find("Arial Bold.ttf")
	if err != nil {
		fontPath, err = findfont.Find("arialbd.ttf")
		if err != nil {
			panic(err)
		}
	}
	if err := dc.LoadFontFace(fontPath, 18); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped("UPPER LEFT", P, P, 0, 0, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("UPPER RIGHT", W-P, P, 1, 0, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("BOTTOM LEFT", P, H-P, 0, 1, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("BOTTOM RIGHT", W-P, H-P, 1, 1, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("UPPER MIDDLE", W/2, P, 0.5, 0, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("LOWER MIDDLE", W/2, H-P, 0.5, 1, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("LEFT MIDDLE", P, H/2, 0, 0.5, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("RIGHT MIDDLE", W-P, H/2, 1, 0.5, 0, 1.5, gg.AlignRight)
	fontPath, err = findfont.Find("Arial.ttf")
	if err != nil {
		panic(err)
	}
	if err := dc.LoadFontFace(fontPath, 12); err != nil {
		panic(err)
	}
	dc.DrawStringWrapped(TEXT, W/2-P, H/2-P, 1, 1, W/3, 1.75, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2-P, 0, 1, W/3, 2, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2-P, H/2+P, 1, 0, W/3, 2.25, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2+P, 0, 0, W/3, 2.5, gg.AlignLeft)
	savePNG("out.png", dc.Image())
}

package gg

import (
	"image"
	"image/color"

	"github.com/golang/freetype/raster"
)

type RepeatOp int

const (
	RepeatBoth RepeatOp = iota
	RepeatX
	RepeatY
	RepeatNone
)

type Pattern interface {
	ColorAt(x, y int) color.Color
}

// Solid Pattern
type solidPattern struct {
	color color.Color
}

func (p *solidPattern) ColorAt(x, y int) color.Color {
	return p.color
}

func NewSolidPattern(color color.Color) Pattern {
	return &solidPattern{color: color}
}

// Surface Pattern
type surfacePattern struct {
	img image.Image
	op  RepeatOp
}

func (p *surfacePattern) ColorAt(x, y int) color.Color {
	b := p.img.Bounds()
	switch p.op {
	case RepeatX:
		if y >= b.Dy() {
			return color.Transparent
		}
	case RepeatY:
		if x >= b.Dx() {
			return color.Transparent
		}
	case RepeatNone:
		if x >= b.Dx() || y >= b.Dy() {
			return color.Transparent
		}
	}
	x = x%b.Dx() + b.Min.X
	y = y%b.Dy() + b.Min.Y
	return p.img.At(x, y)
}

func NewSurfacePattern(img image.Image, op RepeatOp) Pattern {
	return &surfacePattern{img: img, op: op}
}

type patternPainter struct {
	img  *image.RGBA
	mask *image.Alpha
	p    Pattern
}

// Paint satisfies the Painter interface.
func (r *patternPainter) Paint(ss []raster.Span, done bool) {
	b := r.img.Bounds()
	for _, s := range ss {
		if s.Y < b.Min.Y {
			continue
		}
		if s.Y >= b.Max.Y {
			return
		}
		if s.X0 < b.Min.X {
			s.X0 = b.Min.X
		}
		if s.X1 > b.Max.X {
			s.X1 = b.Max.X
		}
		if s.X0 >= s.X1 {
			continue
		}
		const m = 1<<16 - 1
		y := s.Y - r.img.Rect.Min.Y
		x0 := s.X0 - r.img.Rect.Min.X
		// RGBAPainter.Paint() in $GOPATH/src/github.com/golang/freetype/raster/paint.go
		i0 := (s.Y-r.img.Rect.Min.Y)*r.img.Stride + (s.X0-r.img.Rect.Min.X)*4
		i1 := i0 + (s.X1-s.X0)*4
		for i, x := i0, x0; i < i1; i, x = i+4, x+1 {
			ma := s.Alpha
			if r.mask != nil {
				ma = ma * uint32(r.mask.AlphaAt(x, y).A) / 255
				if ma == 0 {
					continue
				}
			}
			c := r.p.ColorAt(x, y)
			cr, cg, cb, ca := c.RGBA()
			dr := uint32(r.img.Pix[i+0])
			dg := uint32(r.img.Pix[i+1])
			db := uint32(r.img.Pix[i+2])
			da := uint32(r.img.Pix[i+3])
			a := (m - (ca * ma / m)) * 0x101
			r.img.Pix[i+0] = uint8((dr*a + cr*ma) / m >> 8)
			r.img.Pix[i+1] = uint8((dg*a + cg*ma) / m >> 8)
			r.img.Pix[i+2] = uint8((db*a + cb*ma) / m >> 8)
			r.img.Pix[i+3] = uint8((da*a + ca*ma) / m >> 8)
		}
	}
}

func newPatternPainter(img *image.RGBA, mask *image.Alpha, p Pattern) *patternPainter {
	return &patternPainter{img, mask, p}
}

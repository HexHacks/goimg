package imgutil

import (
	"image"
	"image/color"
)

type PerPixelFuncI func(x, y int) color.RGBA
type PerPixelFuncF func(x, y float32) color.RGBA

func TraverseSequential(img *image.RGBA, fnc PerPixelFuncI) {
	rct := img.Rect
	for y := 0; y < rct.Max.Y; y++ {
		for x := 0; x < rct.Max.X; x++ {
			start := (y-rct.Min.Y)*img.Stride + (x-rct.Min.X)*4
			rgba := img.Pix[start : start+4]

			clr := fnc(x, y)

			rgba[0] = clr.R
			rgba[1] = clr.G
			rgba[2] = clr.B
			rgba[3] = clr.A
		}
	}
}

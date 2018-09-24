package imgutil

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
)

type PerPixelFuncI func(x, y int) color.RGBA
type PerPixelFuncF func(x, y float32) color.RGBA

func traverseSequentialRgba(img *image.RGBA, fnc PerPixelFuncI) {
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

func Traverse(img image.Image, fc PerPixelFuncI) {
	rgba, ok := img.(*image.RGBA)
	if ok {
		log.Fatal("Can't convert to RGBA")
	}

	traverseSequentialRgba(rgba, fc)
}

func CreateImageFile(folder string, time int) (*os.File, error) {
	fullFolder := fmt.Sprintf("output/%v", folder)
	err := os.MkdirAll(fullFolder, 0700)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("frame_%v.png", time)
	output := fmt.Sprintf("%v/%v", fullFolder, name)
	return os.Create(output)
}

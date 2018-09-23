package scenes

import (
	"fmt"
	"image"
)

type Demo struct {
	test int
	runa rune
}

func NewDemo() *Demo {
	d := &Demo{
		test: 1,
		runa: 'a',
	}

	fmt.Printf("CreateDemo()")
	return d
}

func (d *Demo) Name() string {
	return "Demo"
}

func (d *Demo) Load() {

}

func (d *Demo) Unload() {

}

func (d Demo) Render(img *image.RGBA, t float64) {

}

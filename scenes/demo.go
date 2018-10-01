package scenes

import (
	"fmt"
	"image"
)

type Demo struct {
}

func NewDemo() *Demo {
	d := &Demo{}

	fmt.Printf("CreateDemo()")
	return d
}

func (d *Demo) String() string {
	return "Demo"
}

func (d *Demo) Load() {

}

func (d *Demo) Unload() {

}

func (d Demo) Render(t float64) image.Image {
	return nil
}

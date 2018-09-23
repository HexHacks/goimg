package scenes

import (
	"fmt"
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

func (d *Demo) LoadResources() {

}
func (d *Demo) Render(t float64) {

}
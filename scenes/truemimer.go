package scenes

import "image"

type TrueMimer struct {
}

func NewTrueMimer() *TrueMimer {
	d := &TrueMimer{}

	return d
}

func (d *TrueMimer) String() string {
	return "TrueMimer"
}

func (d *TrueMimer) Load() {

}

func (d *TrueMimer) Unload() {

}

func (d *TrueMimer) Render(img image.Image, t float64) {

}

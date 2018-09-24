package scenes

import (
	"image"
)

type Scene interface {
	Load()
	Unload()
	Render(img image.Image, t float64)
}

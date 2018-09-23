package scenes

import (
	"image"
)

type Scene interface {
	Load()
	Unload()
	Render(img *image.RGBA, t float64)
}

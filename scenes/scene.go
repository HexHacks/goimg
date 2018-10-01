package scenes

import "image"

type Scene interface {
	Load()
	Unload()
	Render(t float64) image.Image
}

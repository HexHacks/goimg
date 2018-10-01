package main

import (
	"image/png"
	"log"

	"./imgutil"
	"github.com/fogleman/fauxgl"
	"github.com/hexhacks/goimg/scenes"
	"github.com/nfnt/resize"
)

const (
	scale  = 1    // Supersampling
	width  = 1920 // Output width in pixels
	height = 1080 // Output height in pixels
)

func main() {
	// Create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)
	//context.ClearColorBufferWith(fauxgl.HexColor("#FFF8E3"))

	scene := scenes.NewTrueMimer(context)
	scene.Load()
	defer scene.Unload()

	image := scene.Render(0)

	// Downsample image for antialiasing
	image = resize.Resize(width, height, image, resize.Bilinear)

	// save image
	// fauxgl.SavePNG("output/truemimer/truemimer.png", image)
	file, err := imgutil.CreateImageFile(scene.String(), 0)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(file, image)
	if err != nil {
		log.Fatal(err)
	}
}

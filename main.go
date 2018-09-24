package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"

	"./imgutil"
	"./scenes"
)

var (
	sceneRing scenes.SceneRing
)

func main() {
	printTodo()

	fmt.Println("Creating new RGBA image")
	img := image.NewRGBA(image.Rect(0, 0, 1024, 800))

	fmt.Println("Clearing image")
	clearImage(img)

	fmt.Println("Opening test.png")
	file, err := imgutil.CreateImageFile("test", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encoding png")
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func clearImage(img *image.RGBA) {
	imgutil.Traverse(img, func(x, y int) color.RGBA {
		return color.RGBA{255, 0, 0, 255}
	})
}

func printTodo() {
	fmt.Println("TODO:::")
	fmt.Println("-> ")
	fmt.Println("-> Basic image traverse")
	fmt.Println("-> Output folder facility (CreateNewFile)")
	fmt.Println("-> Concurrent, per pixel, pipeline")
	fmt.Println("-> Math-function abstractions Function, RealRange, IntRange")
	fmt.Println("-> FFMPEG from image files")
	fmt.Println(":::::::")
	fmt.Println()
}

package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	printTodo()

	fmt.Println("Creating new RGBA image")
	img := image.NewRGBA(image.Rect(0, 0, 1024, 800))

	fmt.Println("Clearing image")
	clearImage(img)

	fmt.Println("Opening test.png")
	file, err := createImageFile("test", 1)
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

func createImageFile(folder string, time int) (*os.File, error) {
	fullFolder := fmt.Sprintf("output/%v", folder)
	err := os.MkdirAll(fullFolder, os.ModeDir)
	if err != nil {
		return nil, err
	}

	name := fmt.Sprintf("frame_%v.png", time)
	output := fmt.Sprintf("%v/%v", fullFolder, name)
	return os.Create(output)
}

func clearImage(img *image.RGBA) {
	rct := img.Rect
	for y := 0; y < rct.Max.Y; y++ {
		for x := 0; x < rct.Max.X; x++ {
			start := (y-rct.Min.Y)*img.Stride + (x-rct.Min.X)*4
			rgba := img.Pix[start : start+4]

			rgba[0] = 255
			rgba[1] = 0
			rgba[2] = 0
			rgba[3] = 255
		}
	}
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

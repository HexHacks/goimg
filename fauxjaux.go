package main

import (
	"fmt"
	"image/png"
	"log"
	"time"

	"./imgutil"

	"github.com/fogleman/fauxgl"
	"github.com/hexhacks/goimg/scenes"
	"github.com/nfnt/resize"
)

const (
	scale  = 1    // optional supersampling
	width  = 1920 // output width in pixels
	height = 1080 // output height in pixels
	fovy   = 30   // vertical field of view in degrees
	near   = 1    // near clipping plane
	far    = 10   // far clipping plane
)

var (
	eye    = fauxgl.V(0, 0, -5)                   // camera position
	center = fauxgl.V(0, -0.07, 0)                // view center position
	up     = fauxgl.V(0, 1, 0)                    // up vector
	light  = fauxgl.V(-0.75, 1, 0.25).Normalize() // light direction
	color  = fauxgl.HexColor("#468966")           // object color
)

func main() {

	tn := time.Now()
	mesh, err := LoadMesh()
	fmt.Println("LoadMesh() ", time.Now().Sub(tn))

	// create a rendering context
	context := fauxgl.NewContext(width*scale, height*scale)
	//context.ClearColorBufferWith(fauxgl.HexColor("#FFF8E3"))

	trueMimer := scenes.NewTrueMimer()
	trueMimer.Load()

	tn = time.Now()
	trueMimer.Render(context.ColorBuffer, 0)
	fmt.Println("Render() ", time.Now().Sub(tn))

	trueMimer.Unload()

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	// use builtin phong shader
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = color
	context.Shader = shader

	// render
	tn = time.Now()
	context.DrawMesh(mesh)
	fmt.Println("DrawMesh() ", time.Now().Sub(tn))

	// downsample image for antialiasing
	image := context.Image()
	image = resize.Resize(width, height, image, resize.Bilinear)

	// save image
	// fauxgl.SavePNG("output/truemimer/truemimer.png", image)
	file, err := imgutil.CreateImageFile("truemimer", 0)
	if err != nil {
		log.Fatal(err)
	}

	err = png.Encode(file, image)
	if err != nil {
		log.Fatal(err)
	}
}

func LoadMesh() (*fauxgl.Mesh, error) {
	// load a mesh
	mesh, err := fauxgl.LoadOBJ("input/truemimer.obj")
	if err != nil {
		return mesh, err
	}

	// fit mesh in a bi-unit cube centered at the origin
	mesh.BiUnitCube()

	// smooth the normals
	mesh.SmoothNormalsThreshold(fauxgl.Radians(30))

	return mesh, nil
}

package scenes

import (
	"fmt"
	"image"
	"time"

	"github.com/fogleman/fauxgl"
	"github.com/hexhacks/goimg/jaux"
	"github.com/hexhacks/goimg/shaders"
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
	eye    = fauxgl.V(0, 0, 5)                     // camera position
	center = fauxgl.V(0, -0.07, 0)                 // view center position
	up     = fauxgl.V(0, 1, 0)                     // up vector
	light  = fauxgl.V(-0.75, -1, 1.25).Normalize() // light direction
	color  = fauxgl.HexColor("#468966")            // object color
)

type TrueMimer struct {
	*fauxgl.Context
	mesh *fauxgl.Mesh
}

func NewTrueMimer(context *fauxgl.Context) *TrueMimer {
	this := &TrueMimer{context, nil}

	// create transformation matrix and light direction
	aspect := float64(width) / float64(height)
	matrix := fauxgl.LookAt(eye, center, up).Perspective(fovy, aspect, near, far)

	shader := shaders.NewScalesShader(matrix, light, eye)
	//shader.Texture = texture
	shader.ObjectColor = color
	context.Shader = shader

	return this
}

func (d *TrueMimer) String() string {
	return "TrueMimer"
}

func (this *TrueMimer) Load() error {
	tn := time.Now()

	mesh, err := jaux.LoadMesh("truemimer")
	if err != nil {
		return err
	}

	this.mesh = mesh

	fmt.Println("TrueMimer.LoadMesh() ", time.Now().Sub(tn))
	return nil
}

func (this *TrueMimer) Unload() {

}

func (this *TrueMimer) Render(t float64) image.Image {
	tn := time.Now()
	//this.Context.DrawMesh(this.mesh)
	fmt.Println("DrawMesh() ", time.Now().Sub(tn))
	return this.Context.Image()
}

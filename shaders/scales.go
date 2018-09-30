package shaders

import (
	"github.com/fogleman/fauxgl"
)

type ScalesShader struct {
	*fauxgl.PhongShader
}

func NewScalesShader(
	matrix fauxgl.Matrix,
	lightDirection, cameraPosition fauxgl.Vector) *ScalesShader {

	p := fauxgl.NewPhongShader(matrix, lightDirection, cameraPosition)
	s := &ScalesShader{p}

	return s
}

func (s *ScalesShader) Vertex(v fauxgl.Vertex) fauxgl.Vertex {
	v.Output = s.PhongShader.Matrix.MulPositionW(v.Position)

	return v
}

func (s *ScalesShader) Fragment(v fauxgl.Vertex) fauxgl.Color {
	color := s.PhongShader.Fragment(v)

	ndotl := v.Normal.Dot(s.LightDirection)

	MulC(&color, ndotl)

	return color
}

func MulC(c *fauxgl.Color, b float64) {
	c.R *= b
	c.G *= b
	c.B *= b
}

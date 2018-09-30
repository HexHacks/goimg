package jaux

import (
	"fmt"

	"github.com/fogleman/fauxgl"
)

func LoadMesh(name string) (*fauxgl.Mesh, error) {

	baseName := fmt.Sprintf("input/%v", name)
	objName := baseName + ".obj"

	mesh, err := fauxgl.LoadOBJ(objName)
	if err != nil {
		return nil, err
	}

	// fit mesh in a bi-unit cube centered at the origin
	mesh.BiUnitCube()

	// smooth the normals
	mesh.SmoothNormalsThreshold(fauxgl.Radians(30))

	return mesh, nil
}

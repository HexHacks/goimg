package scenes

import (
	
)

type Scene interface {
	Name() string
	LoadResources()
	Render(t float64)
}
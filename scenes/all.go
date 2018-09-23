package scenes

import (
)

type SceneRing struct {
	scenes []Scene
	curr int
}

func NewSceneRing() *SceneRing {
	s := &SceneRing{
		[]Scene{		// scenes
			NewDemo(),
		},
		0,				// curr
	}

	s.curr = len(s.scenes) - 1

	return s
}

func (s SceneRing) Render(t float64) {
	
}


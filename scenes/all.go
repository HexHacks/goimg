package scenes

import "image"

type SceneRing struct {
	scenes []Scene
	curr   int
}

func NewSceneRing() *SceneRing {
	s := &SceneRing{
		[]Scene{ // scenes
			NewDemo(),
		},
		0, // curr
	}

	s.curr = len(s.scenes) - 1

	return s
}

func (s *SceneRing) current() Scene {
	return s.scenes[s.curr]
}

func (s *SceneRing) Load() {
	s.current().Load()
}

func (s *SceneRing) Unload() {
	s.current().Unload()
}

func (s SceneRing) Render(img *image.RGBA, t float64) {
	s.current().Render(img, t)
}

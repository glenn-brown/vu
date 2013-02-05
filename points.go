package vu

import "github.com/banthar/gl"

type Points []Point

func (pp Points) Render() {
	gl.Begin(gl.POINTS)
	for _, p := range pp {
		gl.Vertex2d(p.X, p.Y)
	}
	gl.End()
}

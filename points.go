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

// pp.Bounds() returns the minumum and maximum X,Y,Z values occuring
// in Points pp.
func (pp Points) Bounds() (min, max Point) {
	min = pp[0]
	max = pp[0]
	for _, p := range pp {
		min = min.Max(p)
		max = max.Max(p)
	}
	return
}

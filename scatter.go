package vu

import "github.com/banthar/gl"

type Scatter struct {
	points Points3d
}

func NewScatter(pp Points3d) *Scatter {
	return &Scatter{pp}
}

func (s *Scatter) Render(w, h, d float64) {
	gl.PushMatrix()
	min, max := s.points.Bounds()
	dim := max.Sub(min)
	gl.Scaled(w/dim.X, h/dim.Y, d/dim.Z)
	gl.Translated(-min.X, -min.Y, -min.Z)

	// Draw axes: red X, green Y, blue Z.

	gl.Begin(gl.LINES)
	gl.LineWidth(1.5)
	gl.Color3ub(255, 0, 0)
	gl.Vertex3d(min.X, min.Y, min.Z)
	gl.Vertex3d(max.X, min.Y, min.Z)
	gl.Color3ub(0, 255, 0)
	gl.Vertex3d(min.X, min.Y, min.Z)
	gl.Vertex3d(min.X, max.Y, min.Z)
	gl.Color3ub(0, 0, 255)
	gl.Vertex3d(min.X, min.Y, min.Z)
	gl.Vertex3d(min.X, min.Y, max.Z)
	gl.End()

	// Draw 2d plots on the XY, YZ, and XZ planes.

	gl.PointSize(10.0)
	gl.Begin(gl.POINTS)

	// X plot
	gl.Color4ub(255, 0, 0, 31)
	for _, p := range s.points {
		gl.Vertex3d(p.X, min.Y, min.Z)
	}
	// Y plot
	gl.Color4ub(0, 255, 0, 31)
	for _, p := range s.points {
		gl.Vertex3d(min.X, p.Y, min.Z)
	}
	// Z plot
	gl.Color4ub(0, 0, 255, 31)
	for _, p := range s.points {
		gl.Vertex3d(min.X, min.Y, p.Z)
	}

	// XY plot
	gl.Color4ub(255, 255, 0, 63)
	for _, p := range s.points {
		gl.Vertex3d(p.X, p.Y, min.Z)
	}
	// YZ plot
	gl.Color4ub(0, 255, 255, 63)
	for _, p := range s.points {
		gl.Vertex3d(min.X, p.Y, p.Z)
	}
	// XZ plot
	gl.Color4ub(255, 0, 255, 63)
	for _, p := range s.points {
		gl.Vertex3d(p.X, min.Y, p.Z)
	}

	// XYZ plot
	gl.Color4ub(255, 255, 255, 128)
	for _, p := range s.points {
		gl.Vertex3d(p.X, p.Y, p.Z)
	}
	gl.End()
	gl.PopMatrix()
}

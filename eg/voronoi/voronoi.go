package voronoi

import (
	"github.com/banthar/gl"
	vor "github.com/glenn-brown/voronoi"
	"math/rand"
)

//////////////// Voronoi

type Voronoi struct {
	edges     vor.Edges
	points    []*vor.Point
	cellWalls []vor.CellWall
}

func New() *Voronoi {
	D := 1.0

	rng := rand.New(rand.NewSource(42))
	pts := make([]*vor.Point, 600)
	for i := 0; i < len(pts); i++ {
		pts[i] = &vor.Point{D * rng.Float64(), D * rng.Float64()}
	}
	d := &Voronoi{}
	d.edges = vor.GetEdges(pts, 1.0, 1.0)
	cells := d.edges.Cells()
	d.points = make([]*vor.Point, 0, len(cells))
	d.cellWalls = make([]vor.CellWall, 0, len(cells))
	for p, w := range d.edges.Cells() {
		d.points = append(d.points, p)
		d.cellWalls = append(d.cellWalls, w)
	}
	return d
}

func boundedPoint(p *vor.Point) bool {
	return 0 <= p.X && p.X <= 1.0 && 0 <= p.Y && p.Y <= 1.0
}

func boundedEdge(e *vor.Edge) bool {
	return boundedPoint(e.Start) && boundedPoint(e.End)
}

func (v *Voronoi) Render(w, h, d float64) {
	gl.PushMatrix()
	gl.Scaled(w, h, 1)

	rng := rand.New(rand.NewSource(42))
	rng.Seed(42)

	// Draw fill colors

	for _, wall := range v.cellWalls {
		gl.Color3ub(
			uint8(100+rng.Int31n(128)),
			uint8(100+rng.Int31n(128)),
			uint8(100+rng.Int31n(128)))
		gl.Begin(gl.TRIANGLE_FAN)
		for _, p := range wall {
			if p != nil && boundedPoint(p) {
				gl.Vertex2d(p.X, p.Y)
			}
		}
		gl.End()
	}

	// Draw lines

	gl.LineWidth(1.5)
	gl.Color3ub(0, 128, 0)
	gl.Begin(gl.LINES)
	for _, edge := range v.edges {
		if boundedEdge(edge) {
			gl.Vertex2d(edge.Start.X, edge.Start.Y)
			gl.Vertex2d(edge.End.X, edge.End.Y)
		}
	}
	gl.End()

	// Draw points.

	gl.PointSize(2)
	gl.Color3ub(255, 255, 255)
	gl.Begin(gl.POINTS)
	for _, point := range v.points {
		gl.Vertex2d(point.X, point.Y)
	}
	gl.End()
	gl.PopMatrix()
}

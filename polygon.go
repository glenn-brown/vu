package vu

import "github.com/banthar/gl"

// A polygon is a list of vertices in clockwise order.
type Polygon []Point

// Collision returns the first collision along seg.A->seg.B with
// Polygon p, or nil if there is no collision.
func (p Polygon) Intersection(seg Segment) *Point {
	l := len(p)
	at := (Segment{p[l-1], p[0]}).Intersection(seg)
	for i := 0; i < l-1; i++ {
		x := (Segment{p[i], p[i+1]}).Intersection(seg)
		if x == nil {
			continue
		}
		if at == nil ||
			seg.A.Sub(*x).LenSquared() <
				seg.A.Sub(*at).LenSquared() {
			at = x
		}
	}
	return at
}

// Collision returns the first collision along seg.A->seg.B with Polygon p, 
// or nil if there is no collision.
func (p Polygon) Intersects(seg Segment) bool {
	l := len(p)
	if (Segment{p[l-1], p[0]}).Intersects(seg) {
		return true
	}
	for i := 0; i < l-1; i++ {
		if (Segment{p[i], p[i+1]}).Intersects(seg) {
			return true
		}
	}
	return false
}

// IsConvex returns true iff the polygon is convex.
func (p Polygon) IsConvex() bool {
	l := len(p)
	if l < 3 {
		return true
	}
	if CCW(p[l-2], p[l-1], p[0]) || CCW(p[l-1], p[0], p[1]) {
		return false
	}
	for i := 1; i < l-2; i++ {
		if CCW(p[i-1], p[i], p[i+1]) {
			return false
		}
	}
	return true
}

// Return true iff the point is in the convex polygon.
func (poly Polygon) Contains(a Point) bool {
	l := len(poly)
	if CCW(poly[l-1], poly[0], a) {
		return false
	}
	for i := 1; i < l; i++ {
		if CCW(poly[i-1], poly[i], a) {
			return false
		}
	}
	return true
}

func (poly Polygon) Render() {
	modes := [3]gl.GLenum{gl.LINE_LOOP, gl.POLYGON, gl.POINTS}
	for i := range modes {
		gl.Begin(modes[i])
		last := len(poly) - 1
		for i := range poly {
			pt := poly[last-i]
			gl.Vertex2d(pt.X, pt.Y)
		}
		gl.End()
	}
}

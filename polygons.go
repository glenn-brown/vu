package vu

type Polygons []Polygon

// Collison returns the location of the first collision along
// seg.A->seg.B with an edge of the a polygon, or nil if there
// is no collision.
func (pp Polygons) Intersection(seg Segment) (at *Point) {
	for _, p := range pp {
		x := p.Intersection(seg)
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

// HasCollision returns true iff the segment intersects an edge of any
// of the polygons.
func (pp Polygons) Intersect(s Segment) bool {
	for _, p := range pp {
		if p.Intersects(s) {
			return true
		}
	}
	return false
}

func (pp Polygons) Render() {
	for _, p := range pp {
		p.Render()
	}
}

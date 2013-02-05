package vu

type Segment struct{ A, B Point }

// Return true iff segment a intersects segment b.
func (a Segment) Intersects(b Segment) bool {
	return CCW(a.A, b.A, b.B) != CCW(a.B, b.A, b.B) &&
		CCW(a.A, a.B, b.A) != CCW(a.A, a.B, b.B)
}

// Return the pointer where segments a and b intersect, or nil if none.
func (a Segment) Intersection(b Segment) *Point {
	if !a.Intersects(b) {
		return nil
	}
	x1 := a.A.X
	x2 := a.B.X
	x3 := b.A.X
	x4 := b.B.X
	y1 := a.A.Y
	y2 := a.B.Y
	y3 := b.A.Y
	y4 := b.B.Y
	alpha := x1*y2 - y1*x2
	beta := x3*y4 - y3*x4
	x1_x2 := x1 - x2
	x3_x4 := x3 - x4
	y1_y2 := y1 - y2
	y3_y4 := y3 - y4
	d := x1_x2*y3_y4 - y1_y2*x3_x4
	return &Point{
		(alpha*x3_x4 - x1_x2*beta) / d,
		(alpha*y3_y4 - y1_y2*beta) / d}
}

func (s Segment) Len() float64 {
	return s.B.Sub(s.A).Len()
}

func (s Segment) LenSquared() float64 {
	return s.B.Sub(s.A).LenSquared()
}

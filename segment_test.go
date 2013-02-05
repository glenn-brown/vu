package vu

import "testing"

var p00 = Point{0, 0}
var p01 = Point{0, 1}
var p10 = Point{1, 0}
var p11 = Point{1, 1}

func TestIntersects(t *testing.T) {
	// Parallel segments don't intersect.
	if (Segment{p00, p01}).Intersects(Segment{p10, p11}) {
		t.Fail()
	}
	// Common-endpoint segments don't intersect.
	if (Segment{p00, p01}).Intersects(Segment{p01, p11}) {
		t.Fail()
	}
	// Crossing segments do intersect.
	if !(Segment{p00, p11}).Intersects(Segment{p10, p01}) {
		t.Fail()
	}
	// A segment does not intersect itself.
	if (Segment{p00, p01}).Intersects(Segment{p00, p01}) {
		t.Fail()
	}
}

func TestIntersection(t *testing.T) {
	// Crossing segments have the right intersection
	i := (Segment{p00, p11}).Intersection(Segment{p10, p01})
	if !i.Equals(Point{0.5, 0.5}) {
		t.Fail()
	}
	// Parallel segments have no intersection.
	i = (Segment{p00, p01}).Intersection(Segment{p10, p11})
	if i != nil {
		t.Fail()
	}
}

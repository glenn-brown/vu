package vu

import "math"

type Point struct{ X, Y float64 }

func (a Point) Add(b Point) Point   { return Point{a.X + b.X, a.Y + b.Y} }
func (a Point) Dot(b Point) float64 { return a.X*b.X + a.Y*b.Y }
func (a Point) Equals(b Point) bool { return a.X == a.Y && b.X == b.Y }
func (a Point) Len() float64        { return math.Sqrt(a.Dot(a)) }
func (a Point) LenSquared() float64 { return a.Dot(a) }
func (a Point) Sub(b Point) Point   { return Point{a.X - b.X, a.Y - b.Y} }

// Return a point where each of X,Y,Z is the minimum value in input
// point a or b.
//
func (a Point) Min(b Point) Point {
	return Point{math.Min(a.X, b.X), math.Min(a.Y, b.Y)}
}

// Return a point where each of X,Y,Z is the maximum value in input
// point a or b.
//
func (a Point) Max(b Point) Point {
	return Point{math.Max(a.X, b.X), math.Max(a.Y, b.Y)}
}

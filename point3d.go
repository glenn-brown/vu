package vu

import (
	"math"
	"math/rand"
)

type Point3d struct{ X, Y, Z float64 }

func (a Point3d) Add(b Point3d) Point3d {
	return Point3d{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}
func (a Point3d) Cross(b Point3d) Point3d {
	return Point3d{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
	}
}
func (a Point3d) Dot(b Point3d) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}
func (a Point3d) Equals(b Point3d) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}
func (a Point3d) Len() float64 {
	return math.Sqrt(a.Dot(a))
}
func (a Point3d) LenSquared() float64 {
	return a.Dot(a)
}
func (a Point3d) Sub(b Point3d) Point3d {
	return Point3d{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Return a point where each of X,Y,Z is the minimum value in input
// point a or b.
//
func (a Point3d) Min(b Point3d) Point3d {
	return Point3d{
		math.Min(a.X, b.X),
		math.Min(a.Y, b.Y),
		math.Min(a.Z, b.Z)}
}

// Return a point where each of X,Y,Z is the maximum value in input
// point a or b.
//
func (a Point3d) Max(b Point3d) Point3d {
	return Point3d{
		math.Max(a.X, b.X),
		math.Max(a.Y, b.Y),
		math.Max(a.Z, b.Z)}
}

// Return a random point from the Normal distribution.
//
func Point3dNormal() Point3d {
	return Point3d{
		rand.NormFloat64(),
		rand.NormFloat64(),
		rand.NormFloat64(),
	}
}

package vu

import "math"

type Point struct{ X, Y float64 }

func (a Point) Add(b Point) Point   { return Point{a.X + b.X, a.Y + b.Y} }
func (a Point) Dot(b Point) float64 { return a.X*b.X + a.Y*b.Y }
func (a Point) Equals(b Point) bool { return a.X == a.Y && b.X == b.Y }
func (a Point) Len() float64        { return a.X*a.X + a.Y*a.Y }
func (a Point) LenSquared() float64 { return math.Sqrt(a.Len()) }
func (a Point) Sub(b Point) Point   { return Point{a.X - b.X, a.Y - b.Y} }

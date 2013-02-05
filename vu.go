// Package "vu" creates dynamic 3D and 2D visualization with OpenGL.
// 
// For example, the following produces multiple 3D views of a dynamic
// dataset:
// 
//     package main
//     import (
//         "github.com/glenn-brown/vu"
//         "math"
//         "time"
//     )
//     func main() {
//          points := []vu.Point{{-1,-1,-1},{0,0,0},{1,1,1},{2,2,2}}
//          c := vu.Cube(vu.ScatterPlot(&point))
//          v := vu.Wiggle(vu.Hbox(c, vu.Vbox(vu.X(c), vu.Y(c))))
//          w := vu.NewWindow()
// 	 start := time.Now()
//          for {
// 	     t := time.Now().Sub(start).Seconds()
// 	     ports[1].X, ports[1].Y = math.Sin(t), math.Cos(t)
// 	     w.render(v)
//          }
//     }
// 
// To create a custom renderer, implement the Render interface.
//
package vu

import "math"

// A Renderer draws itself in the region bounded by x=0, y=0, z=0,
// x=w, y=h, and z=d, using the gl (OpenGL) package.
//
type Renderer interface {
	Render(w, h, d float64)
}

// Return true iff moving from a to b to c turns counterclockwise.
//
func CCW(a, b, c Point) bool {
	// a->b->c is ccw iff sin(Θ) < 0
	//      c
	//     /
	//    /Θ
	// b +---- a
	v1 := a.Sub(b)
	v2 := c.Sub(b)
	x := v1.X*v2.Y - v1.Y*v2.X // x = |v1 x v2| = sin(Θ)|v1||v2|
	return x < 0
}

// Surge smoothly stops on each integer value.
// Mathematically, surge(x)==x for integer values,
// and surge'(x)==0 for integer values.
//
func surge(x float64) float64 {
	return x/2.0 - math.Sin(math.Pi*x)/(2*math.Pi)
}

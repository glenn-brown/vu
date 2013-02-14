package main

import (
	"github.com/glenn-brown/vu"
	"github.com/glenn-brown/vu/examples/voronoi"
	"math"
	"math/rand"
)

func point() vu.Point3d {
	x := 1.0 * rand.NormFloat64()
	return vu.Point3d{x, math.Sin(x), math.Cos(x)}
}

func points() (pp vu.Points3d) {
	for i := 0; i < 100; i++ {
		pp = append(pp, point())
	}
	return
}

func main() {

	// Create nested views.

	dflt := vu.Frame(voronoi.New())
	frame := vu.Frame(vu.Landscape(dflt))
	vbox := vu.Vbox(dflt, vu.Stack(frame, vu.Cube(vu.Spin(vu.NewScatter(points())))))
	wiggle := vu.Wiggle(vu.Frame(vu.Hbox(vu.Portrait(dflt), vbox)))
	w, err := vu.NewWindow(wiggle)
	if err != nil {
		panic("NewWindow failed")
	}

	// Render them until the window is closed.

	for {
		w.Render()
	}
}

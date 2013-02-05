package main

import (
	"github.com/glenn-brown/vu"
	"github.com/glenn-brown/vu/examples/voronoi"
)

func main() {

	// Create nested views.

	dflt := vu.Frame(voronoi.New())
	frame := vu.Frame(vu.Landscape(dflt))
	vbox := vu.Vbox(dflt, vu.Stack(frame, vu.Cube(vu.Spin(dflt))))
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

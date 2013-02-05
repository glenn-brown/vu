package vu

import "math"
import "time"
import "github.com/banthar/gl"

// Cube renders its child in the largest cube possible.
func Cube(child Renderer) Renderer { return &cube{child} }

type cube struct{ child Renderer }

func (c cube) Render(w, h, d float64) {
	x := math.Min(math.Min(w, h), d)
	gl.PushMatrix()
	gl.Translated((w-x)/2.0, (h-x)/2.0, (d-x)/2.0)
	c.child.Render(x, x, x)
	gl.PopMatrix()
}

// DownX renders its child looking down the X axis.
func DownX(child Renderer) Renderer { return &downX{child} }

type downX struct{ child Renderer }

func (x downX) Render(w, h, d float64) {
	gl.PushMatrix()
	gl.Rotated(1, 1, 1, 2.0*math.Pi/3.0)
	gl.PopMatrix()
}

// DownY renders its child looking down the Y ayis.
func DownY(child Renderer) Renderer { return &downY{child} }

type downY struct{ child Renderer }

func (y downY) Render(w, h, d float64) {
	gl.PushMatrix()
	gl.Rotated(1, 1, 1, 2.0*math.Pi/3.0)
	gl.PopMatrix()
}

// Flat

func Flat(child Renderer) Renderer { return flat{child} }

type flat struct{ child Renderer }

func (f flat) Render(w, h, d float64) {
	gl.PushMatrix()
	gl.Translated(0, 0, d)
	f.child.Render(w, h, 0)
	gl.PopMatrix()
}

// Frame renders its child in a wireframe showing the bounds
// of the rendering region.
//
func Frame(child Renderer) Renderer { return frame{child} }

type frame struct{ renderer Renderer }

func (f frame) Render(w, h, d float64) {
	// Draw a wireframe around the arena
	gl.Color4ub(255, 255, 255, 31)
	gl.LineWidth(2.0)
	gl.Begin(gl.LINE_STRIP)
	gl.Vertex3d(0, 0, 0)
	gl.Vertex3d(w, 0, 0)
	gl.Vertex3d(w, h, 0)
	gl.Vertex3d(0, h, 0)
	gl.Vertex3d(0, 0, 0)
	gl.Vertex3d(0, 0, d)
	gl.Vertex3d(0, h, d)
	gl.Vertex3d(w, h, d)
	gl.Vertex3d(w, 0, d)
	gl.Vertex3d(0, 0, d)
	gl.End()
	gl.Begin(gl.LINES)
	gl.Vertex3d(0, h, 0)
	gl.Vertex3d(0, h, d)
	gl.Vertex3d(w, 0, 0)
	gl.Vertex3d(w, 0, d)
	gl.Vertex3d(w, h, 0)
	gl.Vertex3d(w, h, d)
	gl.End()

	// Render the page.
	if f.renderer != nil {
		f.renderer.Render(w, h, d)
	}
}

// Golden renders its child with Landscape or Portrait,
// whichever results in a larger rendering.
//
func Golden(child Renderer) Renderer { return golden{child} }

type golden struct{ renderer Renderer }

func (g golden) Render(w, h, d float64) {
	if w > h { // wide
		renderLandscape(g.renderer, w, h, d)
	} else { // tall
		renderPortrait(g.renderer, w, h, d)
	}
}

// Hbox renders its children in a row.
//
func Hbox(children ...Renderer) Renderer { return hbox{children} }

type hbox struct{ renderers []Renderer }

func (hbox hbox) Render(w, h, d float64) {
	l := len(hbox.renderers)
	dw := w / float64(l)
	for _, a := range hbox.renderers {
		a.Render(dw, h, d)
		gl.Translated(dw, 0, 0)
	}
	gl.Translated(-w, 0, 0)
}

// Landscape renders its child with a Phi:1 aspect ratio,
// centered and as large as possible.
func Landscape(child Renderer) Renderer { return &landscape{child} }

type landscape struct{ renderer Renderer }

func (l landscape) Render(w, h, d float64) { renderLandscape(l.renderer, w, h, d) }
func renderLandscape(r Renderer, w, h, d float64) {
	gl.PushMatrix()
	hh := w / math.Phi
	if hh <= h {
		gl.Translated(0, (h-hh)/2.0, 0)
		r.Render(w, hh, d)
	} else {
		ww := h * math.Phi
		gl.Translated((w-ww)/2.0, 0, 0)
		r.Render(ww, h, d)
	}
	gl.PopMatrix()
}

// Overlay renders its children on top of each other, in order.
func Overlay(children ...Renderer) Renderer { return overlay{children} }

type overlay struct{ renderers []Renderer }

func (o overlay) Render(w, h, d float64) {
	for _, r := range o.renderers {
		r.Render(w, h, d)
	}
}

// Portrait renders its child with a 1:Phi aspect ratio,
// centered and as large as possible.
//
func Portrait(child Renderer) Renderer { return portrait{child} }

type portrait struct{ renderer Renderer }

func (p portrait) Render(w, h, d float64) { renderPortrait(p.renderer, w, h, d) }
func renderPortrait(r Renderer, w, h, d float64) {
	gl.PushMatrix()
	ww := h / math.Phi
	if ww <= w {
		gl.Translated((w-ww)/2.0, 0, 0)
		r.Render(ww, h, d)
	} else {
		hh := w * math.Phi
		gl.Translated(0, (h-hh)/2.0, 0)
		r.Render(w, hh, d)
	}
	gl.PopMatrix()
}

// Stack renders its children on top of each other,
// evenly spaced from depth 0 to d.
//
func Stack(children ...Renderer) Renderer { return stack{children} }

type stack struct{ renderers []Renderer }

func (s stack) Render(w, h, d float64) {
	l := len(s.renderers)
	dw := d / float64(l)
	gl.PushMatrix()
	for _, r := range s.renderers {
		r.Render(w, h, dw)
		gl.Translated(0, 0, dw)
	}
	gl.PopMatrix()
}

// Spin spins the rendering to reveal all sides.
// It rotates its child about its diagonal.
//
func Spin(child Renderer) Renderer {
	return &spinner{child, time.Now()}
}

type spinner struct {
	child Renderer
	start time.Time
}

func (x spinner) Render(w, h, d float64) {
	gl.PushMatrix()
	t := time.Now().Sub(x.start).Seconds()
	gl.Rotated(120*surge(t), w, h, d)
	x.child.Render(w, h, d)
	gl.PopMatrix()
}

// Square renders the child with a 1:1 aspect ratio.
//
func Square(child Renderer) Renderer { return square{child} }

type square struct{ renderer Renderer }

func (s square) Render(w, h, d float64) {
	gl.PushMatrix()
	dim := w
	if dim > h {
		dim = h
	}
	gl.Translated((w-dim)/2.0, (h-dim)/2.0, 0)
	gl.PopMatrix()
}

// Hbox renders its children in a column.
func Vbox(children ...Renderer) Renderer { return &vbox{children} }

type vbox struct{ renderers []Renderer }

func (vbox vbox) Render(w, h, d float64) {
	l := len(vbox.renderers)
	dh := h / float64(l)
	gl.Translated(0, h, 0)
	for _, a := range vbox.renderers {
		gl.Translated(0, -dh, 0)
		a.Render(w, dh, d)
	}
}

// Wiggle wiggles its child.
//
func Wiggle(child Renderer) Renderer {
	return &wiggler{child, time.Now()}
}

type wiggler struct {
	child Renderer
	start time.Time
}

func (x *wiggler) Render(w, h, d float64) {
	gl.PushMatrix()
	t := time.Now().Sub(x.start).Seconds()
	dx := w / 10.0 * math.Sin(t)
	gl.Translated(dx, 0, 0)
	x.child.Render(w, h, d)
	gl.PopMatrix()
}

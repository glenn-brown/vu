PACKAGE

package vu
    import "github.com/glenn-brown/vu"

    Package "vu" creates dynamic 3D and 2D visualization with OpenGL.

    For example, the following produces multiple 3D views of a dynamic
    dataset:

	    package main
	    import (
	        "github.com/glenn-brown/vu"
	        "math"
	        "time"
	    )
	    func main() {
	         points := []vu.Point{{-1,-1,-1},{0,0,0},{1,1,1},{2,2,2}}
	         c := vu.Cube(vu.ScatterPlot(&point))
	         v := vu.Wiggle(vu.Hbox(c, vu.Vbox(vu.X(c), vu.Y(c))))
	         w := vu.NewWindow()
		 start := time.Now()
	         for {
		     t := time.Now().Sub(start).Seconds()
		     ports[1].X, ports[1].Y = math.Sin(t), math.Cos(t)
		     w.render(v)
	         }
	    }

    To create a custom renderer, implement the Render interface.

FUNCTIONS

func CCW(a, b, c Point) bool
    Return true iff moving from a to b to c turns counterclockwise.


TYPES

type Point struct{ X, Y float64 }

func (a Point) Add(b Point) Point

func (a Point) Dot(b Point) float64

func (a Point) Equals(b Point) bool

func (a Point) Len() float64

func (a Point) LenSquared() float64

func (a Point) Sub(b Point) Point

type Points []Point

func (pp Points) Render()

type Polygon []Point
    A polygon is a list of vertices in clockwise order.

func (poly Polygon) Contains(a Point) bool
    Return true iff the point is in the convex polygon.

func (p Polygon) Intersection(seg Segment) *Point
    Collision returns the first collision along seg.A->seg.B with Polygon p,
    or nil if there is no collision.

func (p Polygon) Intersects(seg Segment) bool
    Collision returns the first collision along seg.A->seg.B with Polygon p,
    or nil if there is no collision.

func (p Polygon) IsConvex() bool
    IsConvex returns true iff the polygon is convex.

func (poly Polygon) Render()

type Polygons []Polygon

func (pp Polygons) Intersect(s Segment) bool
    HasCollision returns true iff the segment intersects an edge of any of
    the polygons.

func (pp Polygons) Intersection(seg Segment) (at *Point)
    Collison returns the location of the first collision along seg.A->seg.B
    with an edge of the a polygon, or nil if there is no collision.

func (pp Polygons) Render()

type Renderer interface {
    Render(w, h, d float64)
}
    A Renderer draws itself in the region bounded by x=0, y=0, z=0, x=w,
    y=h, and z=d, using the gl (OpenGL) package.

func Cube(child Renderer) Renderer
    Cube renders its child in the largest cube possible.

func DownX(child Renderer) Renderer
    DownX renders its child looking down the X axis.

func DownY(child Renderer) Renderer
    DownY renders its child looking down the Y ayis.

func Flat(child Renderer) Renderer

func Frame(child Renderer) Renderer
    Frame renders its child in a wireframe showing the bounds of the
    rendering region.

func Golden(child Renderer) Renderer
    Golden renders its child with Landscape or Portrait, whichever results
    in a larger rendering.

func Hbox(children ...Renderer) Renderer
    Hbox renders its children in a row.

func Landscape(child Renderer) Renderer
    Landscape renders its child with a Phi:1 aspect ratio, centered and as
    large as possible.

func Overlay(children ...Renderer) Renderer
    Overlay renders its children on top of each other, in order.

func Portrait(child Renderer) Renderer
    Portrait renders its child with a 1:Phi aspect ratio, centered and as
    large as possible.

func Spin(child Renderer) Renderer
    Spin spins the rendering to reveal all sides. It rotates its child about
    its diagonal.

func Square(child Renderer) Renderer
    Square renders the child with a 1:1 aspect ratio.

func Stack(children ...Renderer) Renderer
    Stack renders its children on top of each other, evenly spaced from
    depth 0 to d.

func Vbox(children ...Renderer) Renderer
    Hbox renders its children in a column.

func Wiggle(child Renderer) Renderer
    Wiggle wiggles its child.

type Segment struct{ A, B Point }

func (a Segment) Intersection(b Segment) *Point
    Return the pointer where segments a and b intersect, or nil if none.

func (a Segment) Intersects(b Segment) bool
    Return true iff segment a intersects segment b.

func (s Segment) Len() float64

func (s Segment) LenSquared() float64

type Window struct {
    // contains filtered or unexported fields
}

func NewWindow(r Renderer) (*Window, error)

func (w *Window) Render()
    Render draws the contents of the window, but no more than 30 times per
    second.


SUBDIRECTORIES

	examples


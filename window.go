package vu

import (
	"github.com/banthar/gl"
	"github.com/jteeuwen/glfw"
	"os"
	"runtime"
	"time"
)

type Window struct {
	renderer   Renderer
	lastRender time.Time
}

var busy = false

func NewWindow(r Renderer) (*Window, error) {
	w := &Window{r, time.Now()}
	var err error
	// Fail if the window is in use, even after its
	// finalizer has had a change to run.
	if busy {
		runtime.GC()
		runtime.GC()
		if busy {
			return nil, err
		}
	}
	if err = glfw.Init(); err != nil {
		goto abort_with_nothing
	}
	if err = glfw.OpenWindow(1024, 1024, 0, 0, 0, 0, 0, 0, glfw.Windowed); err != nil {
		goto abort_with_glfw_init
	}
	if gl.Init() != 0 {
		goto abort_with_open_window
	}
	glfw.SetWindowTitle(os.Args[0])
	glfw.SetWindowSizeCallback(reshape)
	glInit()
	reshape(1024, 1024)
	runtime.SetFinalizer(w, finalizeWindow)
	busy = true
	return w, nil

abort_with_open_window:
	glfw.CloseWindow()
abort_with_glfw_init:
	glfw.Terminate()
abort_with_nothing:
	return nil, error(err)
}

// This function is called when the window is garbage collected
func finalizeWindow(*Window) {
	glfw.CloseWindow()
	glfw.Terminate()
	busy = false
}

// Render draws the contents of the window, but no more than 30
// times per second.
func (w *Window) Render() {
	// Don't render more than 30 times per second.
	t := time.Now()
	if time.Since(w.lastRender).Seconds() < 1.0/30.0 {
		return
	}
	w.lastRender = t

	// Retrieve the view width and height, in pixels.
	params := make([]int32, 4)
	gl.GetIntegerv(gl.VIEWPORT, params)

	// Infer the rendering volume from the viewport size, so the
	// the rendering code can assume that the width of a pixel is
	// ~1.0.  This is important when rendering text.  This must
	// match the gl.Frustum configuration.
	width := float64(params[2])
	height := float64(params[3])
	depth := (width + height) / 2

	// Redraw the window, but not too often.
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	w.renderer.Render(width, height, depth)
	glfw.SwapBuffers()

	// Exit if the user presses escape or the window was closed.
	if glfw.Key(glfw.KeyEsc) != 0 || glfw.WindowParam(glfw.Opened) == 0 {
		os.Exit(0)
	}
}

// Handle the window being resized.
func reshape(width int, height int) {
	// Render into the entire window.
	gl.Viewport(0, 0, width, height)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()

	// Configure the Frustrum so the front face of the rendering
	// volume fills the screen.

	w := float64(width)
	h := float64(height)
	depth := (w + h) / 2
	near := depth / 2.0
	right := w / 4.0
	top := h / 4.0
	far := 4 * depth
	gl.Frustum(-right, right, -top, top, near, far)

	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	// Center the rendering volume in the viewport.  Its origin
	// is at the far lower left corner.

	gl.Translated(-w/2, -h/2, -2*depth)
}

// Initialize OpenGL settings.
func glInit() {
	pos := []float32{5.0, 5.0, 10.0, 0.0}

	gl.Lightfv(gl.LIGHT0, gl.POSITION, pos)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.NORMALIZE)
	gl.Enable(gl.COLOR_MATERIAL)
	gl.ColorMaterial(gl.FRONT_AND_BACK, gl.AMBIENT_AND_DIFFUSE)
	gl.Enable(gl.LINE_SMOOTH)
	gl.Enable(gl.POINT_SMOOTH)
}

// Draw unit vectors at the origin.
// The X vector is red, Y is green, and Z is blue.
func drawOrigin() {
	gl.Color3ub(0xff, 0, 0)
	gl.Begin(gl.LINES)
	gl.Vertex3d(0, 0, 0)
	gl.Vertex3d(1, 0, 0)
	gl.End()

	gl.Color3ub(0, 0xff, 0)
	gl.Begin(gl.LINES)
	gl.Vertex3d(0, 0, 0)
	gl.Vertex3d(0, 1, 0)
	gl.End()

	gl.Color3ub(0, 0, 0xff)
	gl.Begin(gl.LINES)
	gl.Vertex3d(0, 0, 0)
	gl.Vertex3d(0, 0, 1)
	gl.End()
}

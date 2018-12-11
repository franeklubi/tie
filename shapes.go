package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
)

var (
    // slice holding all chars
    letters     []Image

    // switch indicating if text texture is loaded
    text_loaded bool    = false
)

// BeginShape (together with EndShape) provides a toolset for drawing
// more complex shapes
//
// It should be called with one of the OpenGL consts (see Constants section)
//
// To draw a shape - one of Vertex functions has to be called between the
// BeginShape and EndShape functions
func BeginShape(mode uint32) {
    refreshDepth()
    gl.Begin(mode)
}

// EndShape should be called after BeginShape
//
// It draws a described shape to the buffer
func EndShape() {
    gl.End()
}

// Vertex should be called between BeginShape and EndShape
//
// Vertex lets you specify a vertex of the shape to be drawn using
// standard coordinates ((0,0) at upper left)
func Vertex(x, y float64) {
    gl.Vertex2d(CaToGl(x, y))
}

// Vertex3D should be called between BeginShape and EndShape
//
// Vertex3D lets you specify a 3D vertex of the shape to be drawn using
// standard coordinates ((0,0) at upper left)
func Vertex3D(x, y, z float64) {
    x_gl, y_gl := CaToGl(x, y)
    z_gl := ReMap(x, 0, float64(Width/Height*Width), -1, 1)
    gl.Vertex3d(x_gl, y_gl, z_gl)
}

// VertexGl should be called between BeginShape and EndShape
//
// VertexGl lets you specify a 3D vertex of the shape to be drawn using
// OpenGL standard coordinates ((0,0) at the center)
func VertexGl(x, y, z float64) {
    gl.Vertex3d(x, y, z)
}

// Ellipse draws an ellipse
func Ellipse(x, y, w, h float64) {

    // fill section
    if ( fill_switch ) {

        // setting fill colour
        updateColor(fill)

        BeginShape(TRIANGLE_STRIP)
            var i float64 = 0
            for ; i < 361; i++ {
                pos_x := Cos(DegToRad(i))*w/2 + x
                pos_y := Sin(DegToRad(i))*h/2 + y

                Vertex(x, y)

                Vertex(pos_x, pos_y)
            }
        EndShape()
    }

    // stroke section
    if ( stroke_switch ) {

        // setting opengl colour to stroke
        updateColor(stroke)

        // turning off depth refreshing only for stroke
        change := refresh_depth_mode
        DepthRefreshOff()

        var pos_x, pos_y float64

        BeginShape(LINE_LOOP)
            for i := 0.0; i < 360; i++ {
                pos_x = Cos(DegToRad(i))*(w+1)/2 + x
                pos_y = Sin(DegToRad(i))*(h+1)/2 + y

                Vertex(pos_x, pos_y)
            }
        EndShape()

        // turn depth refresh on only if it was previously turned on
        if ( change ) {
            DepthRefreshOn()
        }
    }
}

// Rect draws a rectangle
func Rect(x, y, w, h float64) {

    if ( fill_switch ) {
        // calculating coords from cartesian
        x_m, y_m := CaToGl(x, y)
        w_m, h_m := CaToGl(x+w, y+h)

        // setting fill colour
        updateColor(fill)

        refreshDepth()

        // and drawing it
        gl.Rectd(x_m, y_m, w_m, h_m)
    }

    if ( stroke_switch ) {

        // setting stroke colour
        updateColor(stroke)

        // stoky stroke gonna choke
        change := refresh_depth_mode
        DepthRefreshOff()

        // drawing my boi stroke
        BeginShape(LINE_LOOP)
            Vertex(x-1, y-1)
            Vertex(x+w+1, y-1)
            Vertex(x+w+1, y+h+1)
            Vertex(x-1, y+h+1)
        EndShape()

        // yatta yatta
        if ( change ) {
            DepthRefreshOn()
        }
    }
}

// Line draws a line
func Line(x1, y1, x2, y2 float64) {
    updateColor(stroke)

    BeginShape(LINES)
        Vertex(x1, y1)
        Vertex(x2, y2)
    EndShape()
}

// Point draws a point
func Point(x, y float64) {
    updateColor(stroke)

    BeginShape(POINTS)
        Vertex(x, y)
    EndShape()
}

// Cube draws a cuboid
func Cube(size float64) {

    updateColor(fill)

    BeginShape(QUADS)
        VertexGl( size,  size, -size)
        VertexGl(-size,  size, -size)
        VertexGl(-size,  size,  size)
        VertexGl( size,  size,  size)

        VertexGl( size, -size,  size)
        VertexGl(-size, -size,  size)
        VertexGl(-size, -size, -size)
        VertexGl( size, -size, -size)

        VertexGl( size,  size,  size)
        VertexGl(-size,  size,  size)
        VertexGl(-size, -size,  size)
        VertexGl( size, -size,  size)

        VertexGl( size, -size, -size)
        VertexGl(-size, -size, -size)
        VertexGl(-size,  size, -size)
        VertexGl( size,  size, -size)

        VertexGl(-size,  size,  size)
        VertexGl(-size,  size, -size)
        VertexGl(-size, -size, -size)
        VertexGl(-size, -size,  size)

        VertexGl( size,  size, -size)
        VertexGl( size,  size,  size)
        VertexGl( size, -size,  size)
        VertexGl( size, -size, -size)
    EndShape()
}

// Sphere draws a sphere
func Sphere(size float64, fidelity uint32) {

    updateColor(fill)

    fide64 := float64(fidelity)

    for theta := 0.0; theta <= fide64; theta++ {

        BeginShape(TRIANGLE_STRIP)

            lon1 := ReMap(theta, 0, fide64, -PI, PI)
            lon2 := ReMap(theta+1, 0, fide64, -PI, PI)

            for fi := 0.0; fi <= fide64; fi++ {

                lat1 := ReMap(fi, 0, fide64, -PI/2, PI/2)

                x1 := size * Sin(lon1) * Cos(lat1)
                y1 := size * Sin(lon1) * Sin(lat1)
                z1 := size * Cos(lon1)

                x2 := size * Sin(lon2) * Cos(lat1)
                y2 := size * Sin(lon2) * Sin(lat1)
                z2 := size * Cos(lon2)

                VertexGl(x1, y1, z1)
                VertexGl(x2, y2, z2)
            }

        EndShape()
    }
}

// Text draws a string to the screen
func Text(txt string, size int, center_align bool) {
    // checking if texture containing letters (font.go) is loaded
    if ( !text_loaded ) {
        Println("Text texture not loaded, loading...")

        // loading letters into an array
        for x := 0; x < 96; x++ {
            letter := letters_tex.GetPixels(64*x, 0, 64, 135)
            letters = append(letters, letter)
        }

        text_loaded = true

        Println("Loaded!")
    }

    for x := 0; x < len(txt); x++ {

        if ( center_align ) {
            Translate(-float64((size*len(txt)))/2, 0, 0)
        }

        // checking if char is supported
        if ( txt[x] < 32 || txt[x] > 126 ) {
            Println("String contains unsupported characters!")

            // printing err char
            PastePixels(letters[95], float64(x*size), 0, float64(size), float64(size)*2.11)
        } else {

            // printing char
            PastePixels(letters[txt[x]-32], float64(x*size), 0, float64(size), float64(size)*2)
        }

        if ( center_align ) {
            Translate(float64((size*len(txt)))/2, 0, 0)
        }
    }
}

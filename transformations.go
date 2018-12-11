package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
)

var (
    // slice of markers for Push() and Pop() functions
    pushpop_tra []int = []int{0}

    // transformation queue
    tra_queue   []transformation = []transformation{}
)

// transformation counts as any type implementing this interface
type transformation interface {
    undo()
}

type rotation struct {
    rotation_name   string
    rotation_value  float64
}

// (rotation) undo undoes rotation
func (r rotation) undo() {
    switch r.rotation_name {
        case "x":
            RotateX(-r.rotation_value)
        case "y":
            RotateY(-r.rotation_value)
        case "z":
            RotateZ(-r.rotation_value)
    }
}

type scalation struct {
    x, y, z float64
}

// (scalation) undo undoes scalation
func (s scalation) undo() {
    Scale(1/s.x, 1/s.y, 1/s.z)
}

type translationGl struct {
    x, y, z float64
}

// (translationGl) undo undoes translationGl
func (t translationGl) undo() {
    TranslateGl(-t.x, -t.y, -t.z)
}

type translation struct {
    x, y, z float64
}

// (translation) undo undoes translation
func (t translation) undo() {
    Translate(-t.x, -t.y, -t.z)
}

// Push stores current state of transformations
//
// To restore it, see: Pop
func Push() {
    pushpop_tra = append(pushpop_tra, len(tra_queue))
}

// Pop should be called after Push
//
// It restores Your matrix to state saved by Push earlier
func Pop() {
    if ( len(tra_queue) != 0 ) {
        last_push := pushpop_tra[len(pushpop_tra)-1]
        neutralize(tra_queue[last_push:len(tra_queue)])
        tra_queue = tra_queue[:last_push]

        pushpop_tra = pushpop_tra[:len(pushpop_tra)-1]
    }
}

// RotateX rotates the matrix along X axis
func RotateX(degrees float64) {
    tra_queue = append(tra_queue, rotation{"x", degrees})

    Translate(-Width/2, -Height/2, 0)
        gl.Rotated(-degrees, 1, 0, 0)
    Translate(Width/2, Height/2, 0)
}

// RotateY rotates the matrix along Y axis
func RotateY(degrees float64) {
    tra_queue = append(tra_queue, rotation{"y", degrees})

    Translate(-Width/2, -Height/2, 0)
        gl.Rotated(-degrees, 0, 1, 0)
    Translate(Width/2, Height/2, 0)
}

// RotateZ rotates the matrix along Z axis
func RotateZ(degrees float64) {
    tra_queue = append(tra_queue, rotation{"z", degrees})

    Translate(-Width/2, -Height/2, 0)
        gl.Rotated(-degrees, 0, 0, 1)
    Translate(Width/2, Height/2, 0)
}

// Translate translates the matrix
func Translate(x, y, z float64) {
    tra_queue = append(tra_queue, translation{x, y, z})
    x_gl := ReMap(x, 0, Width, 0, 2)
    y_gl := ReMap(y, 0, Height, 0, 2)
    z_gl := ReMap(z, 0, Width/Height*Width, 0, 2)

    gl.Translated(x_gl, -y_gl, z_gl)
}

// TranslateGl translates the matrix using
// the OpenGL standard coordinates ((0,0) at the center)
func TranslateGl(x, y, z float64) {
    tra_queue = append(tra_queue, translationGl{x, y, z})

    gl.Translated(x, -y, z)
}

// Scale scales the matrix
func Scale(x, y, z float64) {
    tra_queue = append(tra_queue, scalation{x, y, z})

    if ( x == 0 || y == 0 || z == 0 ) {
        Println("Scale:", x, y, z)
        panic("Attempted division by zero")
    }

    gl.Scaled(x, y, z)
}

// MoveBackFor2D translates the matrix back from the point of origin, so that the entire 2d canvas can be seen
func MoveBackFor2D() {
    z_gl := ReMap(Width/Height*Width/2, 0, Width/Height*Width, 0, 2)

    gl.Translated(0, 0, -z_gl)
}

// MoveForwardFor3D translates the matrix forward, so that the camera is at the point of origin
func MoveForwardFor3D() {
    z_gl := ReMap(Width/Height*Width/2, 0, Width/Height*Width, 0, 2)

    gl.Translated(0, 0, z_gl)
}

// neutralize neutralizes given queue of transformations
func neutralize(queue []transformation) {

    // iterating through the queue in reverse order
    for x := len(queue)-1; x >= 0; x-- {
        t := queue[x]
        t.undo()
    }
}

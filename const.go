package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
    "math"
)

// Consts helpful when interfacing with tie
const (
    // keyboard special keys
    ENTER       string  = "ā"
    ESC         string  = "Ā"
    BACKSPACE   string  = "ă"
    CTRL        string  = "ŕ"
    L_CTRL      string  = "ŕ"
    R_CTRL      string  = "ř"
    SHIFT       string  = "Ŕ"
    L_SHIFT     string  = "Ŕ"
    R_SHIFT     string  = "Ř"
    ALT         string  = "Ŗ"
    ALT_GR      string  = "Ś"
    UP          string  = "ĉ"
    DOWN        string  = "Ĉ"
    LEFT        string  = "ć"
    RIGHT       string  = "Ć"

    // function keys
    F1          string  = "Ģ"
    F2          string  = "ģ"
    F3          string  = "Ĥ"
    F4          string  = "ĥ"
    F5          string  = "Ħ"
    F6          string  = "ħ"
    F7          string  = "Ĩ"
    F8          string  = "ĩ"
    F9          string  = "Ī"
    F10         string  = "ī"
    F11         string  = "Ĭ"
    F12         string  = "ĭ"

    // mouse keys
    MOUSE_LEFT      glfw.MouseButton    = glfw.MouseButtonLeft
    MOUSE_RIGHT     glfw.MouseButton    = glfw.MouseButtonRight
    MOUSE_MIDDLE    glfw.MouseButton    = glfw.MouseButtonMiddle

    // fill and stroke identifiers for updateColor()
    fill    bool    = true
    stroke  bool    = false

    // math consts
    PI  float64 = math.Pi
    E   float64 = math.E

    // OpenGL consts
    POINTS          uint32  = gl.POINTS
    LINES           uint32  = gl.LINES
    LINE_STRIP      uint32  = gl.LINE_STRIP
    LINE_LOOP       uint32  = gl.LINE_LOOP
    POLYGON         uint32  = gl.POLYGON
    TRIANGLES       uint32  = gl.TRIANGLES
    TRIANGLE_STRIP  uint32  = gl.TRIANGLE_STRIP
    TRIANGLE_FAN    uint32  = gl.TRIANGLE_FAN
    QUADS           uint32  = gl.QUADS
    QUAD_STRIP      uint32  = gl.QUAD_STRIP
)

package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
)

var (
    // when the respective events get triggered, these variables will change correspondingly
    Key_pressed bool        = false
    Key         string      = ""
    MouseX      float64     = 0
    MouseY      float64     = 0
    Mouse_key   glfw.MouseButton
    ScrollValue float64     = 0
)

// sets window callbacks
func setCallbacks(w *glfw.Window) {

    w.SetCursorPosCallback(mouseMovedHandler)

    w.SetMouseButtonCallback(mousePressedHandler)

    w.SetKeyCallback(keyHandler)

    w.SetScrollCallback(scrollHandler)

    w.SetSizeCallback(resizeHandler)
}

func mouseMovedHandler(window *glfw.Window, xpos, ypos float64) {
    MouseX = xpos
    MouseY = ypos
    mouseMoved()
}

func mousePressedHandler(w *glfw.Window, button glfw.MouseButton,
    action glfw.Action, mod glfw.ModifierKey) {

    if ( action == glfw.Press ) {
        Mouse_key = button
        mousePressed()
    }
    if ( action == glfw.Release ) {
        mouseReleased()
    }
}

func keyHandler(w *glfw.Window, keycode glfw.Key, scancode int,
    action glfw.Action, mods glfw.ModifierKey) {

    if ( action == glfw.Press ) {
        Key_pressed = true
        Key = string(keycode)
        keyPressed()
    }
    if ( action == glfw.Release ) {
        Key_pressed = false
        Key = string(keycode)
        keyReleased()
    }
}

func scrollHandler(w *glfw.Window, x float64, y float64) {
    if ( x != 0 ) {
        ScrollValue = x
        scrolledX()
    }
    if ( y != 0 ) {
        ScrollValue = y
        scrolledY()
    }
}

func resizeHandler(w *glfw.Window, width, height int) {
    gl.Viewport(0, 0, int32(width), int32(height))
    Width = float64(width)
    Height = float64(height)
}

package tie

import (
    "github.com/go-gl/glfw/v3.2/glfw"
)

var (
    // Width and Height of a created window
    Width       float64
    Height      float64

    // private window vars
    window      *glfw.Window    = nil
    title       string          = "placeholder_title"
    resizable   bool            = false

    // fullscreen vars
    fullscreen              bool    = false
    switch_to_fullscreen    bool    = false
)

// ShowMouse shows mouse
func ShowMouse() {
    window.SetInputMode(glfw.CursorMode, glfw.CursorNormal);
}

// HideMouse hides mouse
func HideMouse() {
    window.SetInputMode(glfw.CursorMode, glfw.CursorHidden);
}

// SetCursorPos sets cursor position
func SetCursorPos(x, y float64) {
    window.SetCursorPos(x, y)
}

// SetSize sets window size
func SetSize(x, y int) {
    window.SetSize(x, y);
}

// SetPos sets window position
func SetPos(x, y int) {
    window.SetPos(x, y);
}

// ToggleFullscreen toggles between fullscreen and windowed modes
func ToggleFullscreen() {
    switch_to_fullscreen = true
}

// fullscreenSwitcher handles switching to fullscreen and back
func fullscreenSwitcher() {
    fullscreen = !fullscreen

    var new_window *glfw.Window

    if ( fullscreen ) {
        Println("Switching to fullscreen")
        new_window = fullscreenWindow(int(Width), int(Height), title)

    } else {
        Println("Switching to windowed")
        new_window = stdWindow(int(Width), int(Height), title)
    }

    // destroying old window
    window.Destroy()

    // replacing old window with a new one
    window = new_window

    setCallbacks(window)

    activateAlpha()
}

// stdWindow returns a standard *glfw.Window window
func stdWindow(width, height int, title string) (*glfw.Window) {

    // initializing glfw
    if err := glfw.Init(); err != nil {
        panic(err)
    }

    // checking if window is to be resizable
    if ( !resizable ) {
        glfw.WindowHint(glfw.Resizable, glfw.False)
    }

    // creating a window
    window, err := glfw.CreateWindow(width, height, title, nil, nil)

    if ( err != nil ) {
        panic(err)
    }

    // binding the window to the current context
    window.MakeContextCurrent()

    return window
}

// fullscreenWindow returns a fullscreen *glfw.Window window
func fullscreenWindow(w, h int, t string) (*glfw.Window) {

    // getting primary monitor
    monitor := glfw.GetPrimaryMonitor()

    // creating a window
    window, err := glfw.CreateWindow(w, h, t, monitor, window)

    if ( err != nil ) {
        panic(err)
    }

    // binding the window to the current context
    window.MakeContextCurrent()

    return window
}

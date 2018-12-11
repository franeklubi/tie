package tie

import (
    "runtime"

    "github.com/go-gl/gl/v2.1/gl"
    "github.com/go-gl/glfw/v3.2/glfw"
)

var (
    // total frame count
    Frames      float64 = 0

    // indicates if the engine should Redraw on its own
    looping     bool    = true

    // indicates if the engine is initialized
    initialized bool    = false
)

// Init initializes engine and sets window's properties
func Init(width, height float64, title string, resizable bool) {

    // important for glfw so we always call it from the same thread
    runtime.LockOSThread()

    // and configuring
    conf(width, height, title, resizable)
}

// conf sets global vars and marks initialized as true
func conf(w, h float64, t string, r bool) {
    Width = w
    Height = h
    title = t
    resizable = r

    initialized = true
}

// Launch launches main loop
func Launch() {
    if ( !initialized ) {
        Println("Failed to launch, engine not initialized")
    }

    if ( !funcs_passed ) {
        Println("Failed to launch, functions not loaded")
    }

    if ( initialized && funcs_passed ) {
        Println("Launching...")
    } else {
        return
    }

    // calling the setup function
    preload()

    // initializing glfw
    window = stdWindow(int(Width), int(Height), title)

    // setting callbacks to events in the window
    setCallbacks(window)

    // defer glfw termination
    defer glfw.Terminate()

    var program = initProgram()

    // enabling smooth lines
    gl.Enable(gl.LINE_SMOOTH);
    gl.Hint(gl.LINE_SMOOTH_HINT, gl.NICEST)

    // using the program and clearing the window
    drawPre(program)

    seedRandom()

    MoveBackFor2D()

    LOG("Launch successful!")

    // calling the setup function
    setup()

    finalizeFrame()

    // starting loop drawing after setup
    drawingLoop()
}

// MAIN DRAWING LOOP
func drawingLoop() {

    for ( !window.ShouldClose() ) {

        // pooling events
        glfw.PollEvents()

        if ( looping ) {
            // clearing depth buffer every frame
            gl.Clear(gl.DEPTH_BUFFER_BIT);

            Redraw()

            // if user toggles fullscreen event we must assure it is happening
            // between frames, not during one
            if ( switch_to_fullscreen ) {
                fullscreenSwitcher()

                // this false is crucial, otherwise we'd just be creating
                // windows like some mad lad
                switch_to_fullscreen = false
            }
        }
    }
}

// Redraw executes the draw loop once
func Redraw() {

    draw()

    finalizeFrame()

    Frames++
}

// finalizeFrame neutralizes transformations and swaps buffers
func finalizeFrame() {
    // neutralizing any transformations
    neutralize(tra_queue)

    // emptying transformation queue
    tra_queue = tra_queue[:0]

    // clear the pushpop transformation queue
    pushpop_tra = pushpop_tra[:0]

    window.SwapBuffers()
}

// NoLoop turns off automatic redrawing
func NoLoop() {
    looping = false
}

// Loop turns on automatic redrawing
func Loop() {
    looping = true
}

// initProgram initializes OpenGL and returns an initialized program
func initProgram() (uint32) {
    if err := gl.Init(); err != nil {
        panic(err)
    }

    prog := gl.CreateProgram()
    gl.LinkProgram(prog)
    return prog
}

func activateAlpha() {
    // enabling alpha channel
    gl.Enable(gl.BLEND)
    gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
}

// setGL sets crucial OpenGL properties
func setGL() {
    activateAlpha()

    gl.Viewport(0, 0, int32(Width), int32(Height))
    gl.MatrixMode(gl.PROJECTION)

    gl.LoadIdentity()
    gl.Frustum(-1, 1, -1, 1, 1, 1000)
    gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)

    gl.Enable(gl.DEPTH_TEST)
    gl.Clear(gl.DEPTH_BUFFER_BIT)

    gl.MatrixMode(gl.MODELVIEW)
    gl.LoadIdentity()
}

// drawPre uses the program, clears the window, and calls setGL
func drawPre(program uint32) {
    gl.Clear(gl.COLOR_BUFFER_BIT)
    gl.Clear(gl.DEPTH_BUFFER_BIT)
    gl.UseProgram(program)

    setGL()
}

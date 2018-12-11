package tie

var (
    // functions to-be-passed
    preload         func()  = func(){}
    setup           func()  = func(){}
    draw            func()  = func(){}
    keyPressed      func()  = func(){}
    keyReleased     func()  = func(){}
    mouseMoved      func()  = func(){}
    mousePressed    func()  = func(){}
    mouseReleased   func()  = func(){}
    scrolledY       func()  = func(){}
    scrolledX       func()  = func(){}

    // codenames of functions to-be-passed
    funcsNames     []string = []string{
        "preload",      "setup",
        "draw",         "keyPressed",
        "keyReleased",  "mouseMoved",
        "mousePressed", "mouseReleased",
        "scrolledY",    "scrolledX",
    }

    // indicates if the minimum amount of functions has been passed
    funcs_passed    bool    = false
)

// PassFunctions serves as a middleman between a user and the engine.
// It receives a number of funcs that are going to be called at certain
// engine's events
func PassFunctions(p ...func()) {

    if ( len(p) < 3 ) {
        Println("You need at least preload, setup, and draw.")
        panic("Not enough functions passed.")
        return
    }

    for x := 0; x < len(funcsNames); x++ {
        if ( x < len(p) ) {
            Println(x, "LOADING", funcsNames[x])
            loadFunc(x, p[x])
        } else {
            Println(x, "SKIPPING", funcsNames[x])
        }
    }

    funcs_passed = true
}

// loadFunc serves as a helper func to PassFunctions, it is used to replace
// local variables containing an empty func with passed func
func loadFunc(x int, to_load func()) {
    switch x {
        case 0:
            preload         = to_load
        case 1:
            setup           = to_load
        case 2:
            draw            = to_load
        case 3:
            keyPressed      = to_load
        case 4:
            keyReleased     = to_load
        case 5:
            mouseMoved      = to_load
        case 6:
            mousePressed    = to_load
        case 7:
            mouseReleased   = to_load
        case 8:
            scrolledY       = to_load
        case 9:
            scrolledX       = to_load
    }
}

// Copyright 2018 franeklubi.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

/*
Package tie provides a Processing-like API for simple and fun drawing,
game making, data and algorithm visualization, generally - art :).

To start writing a new sketch, You need to initialize the engine first:
    tie.Init(width, height, title, is_window_resizable)

Then You need to pass the functions You want to act as the ones listed below,
in the right order, but only preload, setup, and draw are necessary:
    preload         // called only once, before setup, nothing can be drawn here
    setup           // called only once, before draw, you can draw here
    draw            // called once every frame
    keyPressed      // called every time a key gets pressed
    keyReleased     // called every time a key gets released
    mouseMoved      // called every time the mouse is moved
    mousePressed    // called every time a mouse key gets pressed
    mouseReleased   // called every time a mouse key gets released
    scrolledY       // called every time the mouse scroll Y-axis gets scrolled
    scrolledX       // called every time the mouse scroll X-axis gets scrolled

To do that, call PassFunctions with the functions You want to use as arguments:
    tie.PassFunctions(
        preload,
        setup,
        draw,
    )

The only thing that's left is launching the engine with:
    tie.Launch()


The whole sketch should look something like this:
    package main

    // import the package
    import (
        "github.com/franeklubi/tie"
    )

    func main() {
        // initialize engine in main
        tie.Init(500, 500, "window_name", false)
        //   width, height, window_name, is_resizable

        // pass all the functions you want used by the engine
        tie.PassFunctions(
            preload,
            setup,
            draw,
        )

        // launch the engine
        tie.Launch()
    }

    var (
        img tie.Image
    )

    func preload() {
        img = tie.LoadImage("/path/to/image.png")
    }

    func setup() {
        tie.Background(255, 255, 255, 255)
    }

    func draw() {
        tie.Fill(0, 255, 255, 255)
        tie.Rect(tie.Width/2-50, tie.Height/2-50, 100, 100)
    }

For more examples visit https://github.com/franeklubi/tie-examples/

*/
package tie

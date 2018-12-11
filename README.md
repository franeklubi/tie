# franeklubi/tie

[![GoDoc](https://godoc.org/github.com/franeklubi/tie?status.svg)](https://godoc.org/github.com/franeklubi/tie)

![logo](assets/tie_logo_64.png)

Package `franeklubi/tie` provides a Processing-like API for simple and fun drawing,
game making, data and algorithm visualization, and generally - art :)

---

* [Installation](#installation)
* [Dependencies](#dependencies)
* [Features](#features)
* [Example](#example)

---

## Installation

To install this package:

```sh
go get github.com/franeklubi/tie
```

You'll also need to install dependencies. (see [Dependencies](#dependencies))

## Dependencies

This package depends on two other packages:
* gl v2.1
* glfw v3.2

To install these, use:
```
go get github.com/go-gl/gl/v2.1/gl
```
then
```
go get github.com/go-gl/glfw/v3.2/glfw
```
and You should be all set! :)

## Features

Main features include:
* it is beginner friendly,
* has Processing-like API - no need to learn a new framework,
* includes easy image handling
* ...

## Example

( For more examples visit [franeklubi/tie-examples](https://github.com/franeklubi/tie-examples/)! )

```go
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

// called only once, before setup, nothing can be drawn here
func preload() {
    img = tie.LoadImage("/path/to/image.png")
}

// called only once, before draw, you can draw here
func setup() {
    tie.Background(255, 255, 255, 255)
}

// called once every frame
func draw() {
    // drawing loaded image
    tie.Fill(255, 255, 255, 255)
    tie.PastePixels(img, 0, 0, tie.Width, tie.Height)

    // drawing ellipse
    tie.Fill(0, 255, 0, 255)
    tie.Ellipse(tie.Width/2, tie.Height/2, 200, 200)

    // drawing rectangle
    tie.Fill(0, 255, 255, 255)
    tie.Rect(tie.Width/2-50, tie.Height/2-50, 100, 100)
}
```

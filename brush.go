package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
)

var (
    fill_color      ColorGl   = ColorGl{1, 1, 1, 1}
    stroke_color    ColorGl   = ColorGl{0, 0, 0, 1}
    fill_switch     bool    = true
    stroke_switch   bool    = true

    refresh_depth_mode = true
)

// Background fills the whole frame with one colour
func Background( r, g, b, a byte ) {
    rf, gf, bf, af := RgbToGl( r, g, b, a )
    gl.ClearColor(float32(rf), float32(gf), float32(bf), float32(af))
    gl.Clear(gl.COLOR_BUFFER_BIT)
}

// Fill sets colour used to fill shapes based on RGBA values
func Fill( r, g, b, a byte ) {
    col1, col2, col3, alpha := RgbToGl( r, g, b, a )
    fill_color = ColorGl{col1, col2, col3, alpha}
    fill_switch = true

    updateColor(fill)
}

// NoFill turns off fill
func NoFill() {
    fill_switch = false
}

// Stroke sets colour used to draw shapes' outline based on RGBA values
// (applies to Line and Point too)
func Stroke( r, g, b, a byte ) {
    col1, col2, col3, alpha := RgbToGl( r, g, b, a )
    stroke_color = ColorGl{col1, col2, col3, alpha}
    stroke_switch = true

    updateColor(stroke)
}

// StrokeWidth sets width of the stroke
func StrokeWidth(width float64) {
    gl.LineWidth(float32(width))
    gl.PointSize(float32(width))
}

// NoStroke turns off stroke
func NoStroke() {
    stroke_switch = false
}

// updateColor updates used colour based on mode
func updateColor(mode bool) {
    var col ColorGl
    if ( mode == fill ) {
        col = fill_color
    } else if ( mode == stroke ) {
        col = stroke_color
    }

    gl.Color4d(col.R, col.G, col.B, col.A)
}

// DepthRefreshOn turns automatic depth refreshing on
func DepthRefreshOn() {
    refresh_depth_mode = true
}

// DepthRefreshOff turns automatic depth refreshing off
//
// Every time a shape is drawn the depth buffer is cleared to prevent Z-fighting
func DepthRefreshOff() {
    refresh_depth_mode = false
}

// refreshDepth clears depth buffer if depth refresh mode is on
func refreshDepth() {
    if ( refresh_depth_mode ) {
        gl.Clear(gl.DEPTH_BUFFER_BIT);
    }
}

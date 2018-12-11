package tie

import (
    "math/rand"
    "time"
)

// Random returns a random float64 between 0 and n
func Random(n float64) (float64) {
    return rand.Float64() * n
}

// seedRandom seeds rand with current time
func seedRandom() {
    rand.Seed(time.Now().UTC().UnixNano())
}

// ReMap converts a number that exists in one range to another
func ReMap( val, start1, stop1, start2, stop2 float64 ) (float64) {
    return (val-start1)*((stop2-start2)/(stop1-start1))+start2
}

// LinInt provides linear interpolation (lerp)
func LinInt(val, target, rate float64) (float64) {
    return val+(target-val)*rate
}

// Limit constricts a value between a given range
func Limit(val, min, max float64) (float64) {
    if ( val < min ) {
        return min
    }
    if ( val > max ) {
        return max
    }
    return val
}

// DegToRad converts degrees to radians
func DegToRad(degrees float64) (float64) {
    return degrees * (PI/180)
}

// RadToDeg converts radians to degrees
func RadToDeg(radians float64) (float64) {
    return radians * 180 / PI
}

// RgbToGl maps 0-255 range bytes to 0.0-1.0 range floats
func RgbToGl(r, g, b, a byte) (float64, float64, float64, float64) {

	var r_ret float64 = ReMap( float64(r), 0.0, 255.0, 0.0, 1.0 )
	var g_ret float64 = ReMap( float64(g), 0.0, 255.0, 0.0, 1.0 )
	var b_ret float64 = ReMap( float64(b), 0.0, 255.0, 0.0, 1.0 )
	var a_ret float64 = ReMap( float64(a), 0.0, 255.0, 0.0, 1.0 )

	return r_ret, g_ret, b_ret, a_ret
}

// Desired usage:
//     Fill(HsvToRgb(h, s, v, a))
//     Stroke(HsvToRgb(h, s, v, a))
//     Background(HsvToRgb(h, s, v, a))
//
// RgbToHsv returns RGBA values based on HSV and Alpha values.
//
// (hue=0-360, saturation=0-1, value=0-1, alpha=0-255)
func HsvToRgb(hue, saturation, value float64, alpha byte) (byte, byte, byte, byte) {
    var chroma float64 = value*saturation

    var hue_m float64 = hue/60

    var x float64 = chroma * ( 1 - Abs( Mod(hue_m, 2) - 1 ) )
    var r, g, b float64 = 0, 0, 0

    switch(true) {
        case ( 0 <= hue_m && hue_m <= 1 ):
            r, g, b = chroma, x, 0

        case ( 1 <  hue_m && hue_m <= 2 ):
            r, g, b = x, chroma, 0

        case ( 2 <  hue_m && hue_m <= 3 ):
            r, g, b = 0, chroma, x

        case ( 3 <  hue_m && hue_m <= 4 ):
            r, g, b = 0, x, chroma

        case ( 4 <  hue_m && hue_m <= 5 ):
            r, g, b = x, 0, chroma

        case ( 5 <  hue_m && hue_m <= 6 ):
            r, g, b = chroma, 0, x
    }

    m := value-chroma
    r, g, b = (r+m)*255, (g+m)*255, (b+m)*255

    return byte(r), byte(g), byte(b), alpha
}

// CaToGl maps the tie coordinates ((0,0) at upper left) to
// OpenGL standard coordinates ((0,0) at the center)
func CaToGl(x, y float64) (float64, float64) {

    var x_r float64 = ReMap(x, 0, float64(Width), -1, 1)
    var y_r float64 = ReMap(y, 0, float64(Height), 1, -1)

    return x_r, y_r
}

// GlToCa maps OpenGL standard coordinates ((0,0) at the center)
// to the tie coordinates ((0,0) at upper left)
func GlToCa(x, y float64) (float64, float64) {

    var x_r float64 = ReMap(x, -1, 1, 0, float64(Width))
    var y_r float64 = ReMap(y, 1, -1, 0, float64(Height))

    return x_r, y_r
}

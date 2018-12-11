package tie

import (
    "math"
    "log"
    "fmt"
)

// Println links to fmt.Println
func Println(itm ...interface{}) {
    fmt.Println(itm...)
}

// Print links to fmt.Print
func Print(itm ...interface{}) {
    fmt.Print(itm...)
}

// LOG links to log.Println
func LOG(itm ...interface{}) {
    log.Println(itm...)
}

// Ceil links to math.Ceil
func Ceil(n float64) (float64) {
    return math.Ceil(n)
}

// Floor links to math.Floor
func Floor(n float64) (float64) {
    return math.Floor(n)
}

// Sin links to math.Sin
func Sin(n float64) (float64) {
    return math.Sin(n)
}

// Asin links to math.Asin
func Asin(n float64) (float64) {
    return math.Asin(n)
}

// Cos links to math.Cos
func Cos(n float64) (float64) {
    return math.Cos(n)
}

// Acos links to math.Acos
func Acos(n float64) (float64) {
    return math.Acos(n)
}

// Tan links to math.Tan
func Tan(n float64) (float64) {
    return math.Tan(n)
}

// Atan links to math.Atan
func Atan(n float64) (float64) {
    return math.Atan(n)
}

// Atan2 links to math.Atan2
func Atan2(n1, n2 float64) (float64) {
    return math.Atan2(n1, n2)
}

// Sqrt links to math.Sqrt
func Sqrt(n float64) (float64) {
    return math.Sqrt(n)
}

// Abs links to math.Abs
func Abs(n float64) (float64) {
    return math.Abs(n)
}

// Mod links to math.Mod
func Mod(n1, n2 float64) (float64) {
    return math.Mod(n1, n2)
}

// Rotate links to RotateZ
func Rotate(degrees float64) {
    RotateZ(degrees)
}

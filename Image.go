package tie

// Image is used to hold the loaded image's pixel and size data
//
// Pixels are stored in a down to up && left to right arrangement
//
// To manipulate them see (*Image) Replace
type Image struct {
    Pixels []byte   // Pixels of a loaded image
    W, H int        // Width and Height of an Image
}

// Color is used to hold the 0-255 range RGBA values
type Color struct {
    R, G, B, A byte
}

// ColorGl is used to hold the 0-1 range RGBA values
type ColorGl struct {
    R, G, B, A float64
}

// GenImage generates an Image of specified dimensions filled with a given Color
func GenImage(c Color, w, h int) (Image) {
    var img Image = Image{[]byte{}, w, h}

    for x := 0; x < w*h; x++ {
        img.PushPixel(c)
    }

    return img
}

// PixelAt returns Color of a pixel at given coordinates
func (img *Image) PixelAt(x, y int) (Color) {

    x = int( Limit(float64(x), 0, float64(img.W)-1) )
    y = int( Limit(float64(img.H-1-y), 0, float64(img.H)-1) )

    r := img.Pixels[img.W*y*4+x*4]
    g := img.Pixels[img.W*y*4+x*4+1]
    b := img.Pixels[img.W*y*4+x*4+2]
    a := img.Pixels[img.W*y*4+x*4+3]

    return Color{r, g, b, a}
}

// PushPixel appends a Color pixel to an Image
func (img *Image) PushPixel(c Color) {
    pixel := []byte{c.R, c.G, c.B, c.A}
    img.Pixels = append(pixel, img.Pixels...)
}

// Replace replaces a pixel in an Image with a Color
func (img *Image) Replace(x, y int, pixel Color) {
    x = int( Limit(float64(x), 0, float64(img.W)-1) )
    y = int( Limit(float64(img.H-1-y), 0, float64(img.H)-1) )

    img.Pixels[img.W*y*4+x*4]   = pixel.R
    img.Pixels[img.W*y*4+x*4+1] = pixel.G
    img.Pixels[img.W*y*4+x*4+2] = pixel.B
    img.Pixels[img.W*y*4+x*4+3] = pixel.A
}

// GetPixels returns a section of received Image
func (img *Image) GetPixels(x, y, w, h int) (Image) {

    var pixels []byte = []byte{}
    for j := h-1; j >= 0; j-- {
        for i := 0; i < w; i++ {
            c := img.PixelAt(i+x, j+y)
            pixels = append(pixels, c.R)
            pixels = append(pixels, c.G)
            pixels = append(pixels, c.B)
            pixels = append(pixels, c.A)
        }
    }

    return Image{pixels, w, h}
}

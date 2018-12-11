package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
    "image"
    "image/png"
    "os"
)

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

// ToArray returns a byte array containing Color's values
func (c Color) ToArray() ([]byte) {
    return []byte{c.R, c.G, c.B, c.A}
}

// ColorGl is used to hold the 0-1 range RGBA values
type ColorGl struct {
    R, G, B, A float64
}

// ToArray returns a float64 array containing ColorGl's values
func (c ColorGl) ToArray() ([]float64) {
    return []float64{c.R, c.G, c.B, c.A}
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

    // checking for image corruption
    if ( img.isCorrupted() ) {
        Println("Image is corrupted or missing pixels")
        return Color{0, 0, 0, 0}
    }

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

    // checking for image corruption
    if ( img.isCorrupted() ) {
        Println("Image is corrupted or missing pixels")
        return
    }

    x = int( Limit(float64(x), 0, float64(img.W)-1) )
    y = int( Limit(float64(img.H-1-y), 0, float64(img.H)-1) )

    img.Pixels[img.W*y*4+x*4]   = pixel.R
    img.Pixels[img.W*y*4+x*4+1] = pixel.G
    img.Pixels[img.W*y*4+x*4+2] = pixel.B
    img.Pixels[img.W*y*4+x*4+3] = pixel.A
}

// GetPixels returns a section of received Image
func (img *Image) GetPixels(x, y, w, h int) (Image) {

    // checking for image corruption
    if ( img.isCorrupted() ) {
        Println("Image is corrupted or missing pixels")
        return Image{[]byte{}, 0, 0}
    }

    var pixels []byte = []byte{}
    for j := h-1; j >= 0; j-- {
        for i := 0; i < w; i++ {
            c := img.PixelAt(i+x, j+y).ToArray()
            pixels = append(pixels, c...)
        }
    }

    return Image{pixels, w, h}
}

// Save saves an image to a file in a png format
func (img *Image) Save(path string) {

    // checking for image corruption
    if ( img.isCorrupted() ) {
        Println("Image is corrupted or missing pixels")
        return
    }

    // creating an empty image
    created := image.NewRGBA(image.Rect(0, 0, img.W, img.H))

    // inverting the image before saving
    proper_y_axis := []byte{}
    for y := 0; y < img.H; y++ {
        for x := 0; x < img.W; x++ {
            c := img.PixelAt(x, y).ToArray()
            proper_y_axis = append(proper_y_axis, c...)
        }
    }

    created.Pix = proper_y_axis

    out, err := os.Create(path)

    if ( err != nil ) {
        Println(err)
        Println("Can't create file at", path)
        return
    }

    defer out.Close()

    err = png.Encode(out, created)

    if ( err != nil ) {
        Println(err)
        Println("Can't encode file from given img")
        return
    }

    Println("Saved to", path)
}

// CopyPixels returns an Image containing current frame buffer
func CopyPixels() (Image) {
    // making an empty []byte array
    pixels := make([]byte, 4*int32(Width*Height))

    // reading pixels from current buffer
    gl.ReadPixels(0, 0, int32(Width), int32(Height),
        gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

    return Image{pixels, int(Width), int(Height)}
}

// PastePixels draws an Image to the screen
func (img *Image) PastePixels(x, y, w, h float64) {

    // checking for image corruption
    if ( img.isCorrupted() ) {
        Println("Image is corrupted or missing pixels")
        return
    }

    // updating colour so that the texture will be able to have
    // a tint based on current fill
    updateColor(fill)

	var texture uint32
	gl.GenTextures(1, &texture)
    gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.TexImage2D(
		gl.TEXTURE_2D, 0,
		gl.RGBA,
		int32(img.W), int32(img.H), 0,
		gl.RGBA,
		gl.UNSIGNED_BYTE, gl.Ptr(img.Pixels),
	)

	gl.Enable(gl.TEXTURE_2D)
    	BeginShape(QUADS)
    		gl.TexCoord2i(0, 1)
            Vertex(x, y)

    		gl.TexCoord2i(1, 1)
            Vertex(x+w, y)

    		gl.TexCoord2i(1, 0)
            Vertex(x+w, y+h)

    		gl.TexCoord2i(0, 0)
            Vertex(x, y+h)
    	EndShape()
    gl.Disable(gl.TEXTURE_2D)

    gl.DeleteTextures(1, &texture)
}

// isCorrupted checks if there is the exact amount of pixels in an image
func (img *Image) isCorrupted() (bool) {
    area := img.W*img.H
    if ( area != len(img.Pixels)/4 && area != 0 ) {
        return true
    }
    return false
}

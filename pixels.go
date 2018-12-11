package tie

import (
    "github.com/go-gl/gl/v2.1/gl"
)

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
func PastePixels(img Image, x, y, w, h float64) {

    // checking for image corruption
    if ( img.W*img.H != len(img.Pixels)/4 ) {
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

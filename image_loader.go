package tie

import (
    "image"
    "image/png"
    "image/jpeg"
    "os"
    "io"
)

// LoadImage loads an image (.png or .jpg) from a given path
func LoadImage(path string) (Image) {
    // registering 2 most common formats
    image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
    image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)

    file, err := os.Open(path)

    if err != nil {
        panic("Can't open file")
    }

    defer file.Close()

    img, err := loadPixels(file)

    if err != nil {
        panic("Can't decode file")
    }

    return img
}

// loadPixels loads pixels of an image
func loadPixels(file io.Reader) (Image, error) {
    img, _, err := image.Decode(file)

    if err != nil {
        return Image{[]byte{}, 0, 0}, err
    }

    // getting image size
    bounds := img.Bounds()
    width, height := bounds.Max.X, bounds.Max.Y

    // iterating and appending every pixel in an image
    var pixels []byte
    for y := height-1; y >= 0; y-- {
        for x := 0; x < width; x++ {
            var c Color = toPix(img.At(x, y).RGBA())
            pixels = append(pixels, c.R)
            pixels = append(pixels, c.G)
            pixels = append(pixels, c.B)
            pixels = append(pixels, c.A)
        }
    }

    return Image{pixels, width, height}, nil
}

// toPix converts uint32 rgba colorspace to Color
func toPix(r, g, b, a uint32) (Color) {
    return Color{byte(r/257), byte(g/257), byte(b/257), byte(a/257)}
}

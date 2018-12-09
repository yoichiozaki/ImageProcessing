/* Package io provides basic image file I/O. */
package io

import (
	"image"
	"image/png"
	"io"
	"os"
)

// Encoder encodes the provided image and writes it out.
type Encoder func(io.Writer, image.Image) error

// Open loads and decodes an image from a file and returns it.
func Open(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// PNGEncoder returns an encoder to PNG format.
func PNGEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}

// Save creates a file and writes to it an image using the provided encoder.
func Save(filename string, img image.Image, encoder Encoder) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return encoder(f, img)
}
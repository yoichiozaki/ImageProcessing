/* Package convolution provides useful functions to create and apply kernel to an image. */
package convolution

import (
	"ImageProcessing/clone"
	"ImageProcessing/parallel"
	"image"
	"math"
)

type Options struct {
	Wrap bool // Wrap sets if indices outside of image dimensions should be taken from the opposite side.
}

func Convolve(img image.Image, kernel Matrix, opt *Options) *image.Gray {
	// log.Println(kernel.String())
	wrap := false
	if opt != nil {
		wrap = opt.Wrap
	}
	return execute(img, kernel, wrap)
}

func execute(img image.Image, kernel Matrix, wrap bool) *image.Gray {
	X := kernel.MaxX()
	Y := kernel.MaxY()
	radiusX := X/2
	radiusY := Y/2
	var source *image.Gray
	if wrap {
		source = clone.Pad(img, radiusX, radiusY, clone.EdgeWrap)
	} else {
		source = clone.Pad(img, radiusX, radiusY, clone.EdgeExtend)
	}
	srcBounds := source.Bounds()
	srcWidth := srcBounds.Dx()
	srcHeight := srcBounds.Dy()
	dst := image.NewGray(img.Bounds())
	_execution := func(start, end int) {
		for y := start + radiusY; y < end+radiusY; y++ {
			for x := radiusX; x < srcWidth-radiusX; x++ {
				var c float64
				for ky := 0; ky < Y; ky++ {
					iy := y - radiusY + ky
					for kx := 0; kx < X; kx++ {
						ix := x - radiusX + kx
						kvalue := kernel.At(kx, ky)
						ipos := iy*source.Stride + ix
						c += float64(source.Pix[ipos]) * kvalue
					}
				}
				pos := (y-radiusY)*dst.Stride + (x-radiusX)
				dst.Pix[pos] = uint8(math.Max(math.Min(c, 255), 0))
			}
		}
	}
	parallel.Line(srcHeight-(radiusY*2), _execution)
	return dst
}



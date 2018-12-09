/* Package clone provides image cloning functions. */
package clone

import (
	"ImageProcessing/parallel"
	"image"
	"image/color"
	"image/draw"
)

type PaddingMethod uint8

const (
	NoFill     = iota // NoFill leaves the new pixels.
	EdgeExtend        // EdgeExtend extends the closest edge pixel for padding new pixels.
	EdgeWrap          // EdgeWrap wraps around the new pixels of an image.
)

func AsRGBA(src image.Image) *image.RGBA {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)
	draw.Draw(img, bounds, src, bounds.Min, draw.Src)
	return img
}

func AsGray(src image.Image) *image.Gray {
	bounds := src.Bounds()
	img := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := color.GrayModel.Convert(src.At(x, y))
			gray, _ := c.(color.Gray)
			img.Set(x, y, gray)
		}
	}
	return img
}

func Pad(src image.Image, padX, padY int, method PaddingMethod) *image.Gray {
	var result *image.Gray
	switch method {
	case EdgeExtend:
		result = extend(src, padX, padY)
	case NoFill:
		result = nofill(src, padX, padY)
	case EdgeWrap:
		result = wrap(src, padX, padY)
	default:
		result = extend(src, padX, padY)

	}
	return result
}

// 							+---+---+---+---+---+---+
// 							| 0 | 0 | 0 | 0 | 0 | 0 |
// +---+---+---+---+		+---+---+---+---+---+---+
// | # | # | # | # |		| 0 | # | # | # | # | 0 |
// +---+---+---+---+		+---+---+---+---+---+---+
// | # | # | # | # |		| 0 | # | # | # | # | 0 |
// +---+---+---+---+ ----->	+---+---+---+---+---+---+
// | # | # | # | # |		| 0 | # | # | # | # | 0 |
// +---+---+---+---+		+---+---+---+---+---+---+
// | # | # | # | # |		| 0 | # | # | # | # | 0 |
// +---+---+---+---+		+---+---+---+---+---+---+
// 							| 0 | 0 | 0 | 0 | 0 | 0 |
// 							+---+---+---+---+---+---+

func nofill(srcImg image.Image, padX, padY int) *image.Gray {
	srcBounds := srcImg.Bounds()
	paddedWidth := srcBounds.Dx() + 2*padX
	paddedHeight := srcBounds.Dy() + 2*padY
	newBounds := image.Rect(0, 0, paddedWidth, paddedHeight)
	fillBounds := image.Rect(padX, padY, padX+srcBounds.Dx(), padY+srcBounds.Dy())
	dst := image.NewGray(newBounds)
	draw.Draw(dst, fillBounds, srcImg, srcBounds.Min, draw.Src)
	return dst
}

// 					 			 +---+---+---+---+---+---+---+---+
// 					 			 | ← | ← | ↑ | ↑ | ↑ | ↑ | → | → |
// 					 			 +---+---+---+---+---+---+---+---+
// 					 			 | ← | ← | ↑ | ↑ | ↑ | ↑ | → | → |
// 	+---+---+---+---+			 +---+---+---+---+---+---+---+---+
// 	| # | # | # | # |			 | ← | ← | # | # | # | # | → | → |
// 	+---+---+---+---+			 +---+---+---+---+---+---+---+---+
// 	| # | # | # | # |			 | ← | ← | # | # | # | # | → | → |
// 	+---+---+---+---+	----->	 +---+---+---+---+---+---+---+---+
// 	| # | # | # | # |			 | ← | ← | # | # | # | # | → | → |
// 	+---+---+---+---+			 +---+---+---+---+---+---+---+---+
// 	| # | # | # | # |			 | ← | ← | # | # | # | # | → | → |
// 	+---+---+---+---+			 +---+---+---+---+---+---+---+---+
// 								 | ← | ← | ↓ | ↓ | ↓ | ↓ | → | → |
// 								 +---+---+---+---+---+---+---+---+
// 								 | ← | ← | ↓ | ↓ | ↓ | ↓ | → | → |
// 								 +---+---+---+---+---+---+---+---+

func extend(srcImg image.Image, padX, padY int) *image.Gray {
	dst := nofill(srcImg, padX, padY)
	paddedWidth := dst.Bounds().Dx()
	paddedHeight := dst.Bounds().Dy()
	padding := func(start, end int) {
		for y := start; y < end; y++ {
			iy := y
			if iy < padY {
				iy = padY
			} else if iy >= paddedHeight-padY {
				iy = paddedHeight - padY - 1
			}

			for x := 0; x < paddedWidth; x++ {
				ix := x
				if ix < padX {
					ix = padX
				} else if x >= paddedWidth-padX {
					ix = paddedWidth - padX - 1
				} else if iy == y {
					// This only enters if we are not in a y-padded area or
					// x-padded area, so nothing to extend here.
					// So simply jump to the next padded-x index.
					x = paddedWidth - padX - 1
					continue
				}
				dstPos := y*dst.Stride + x
				edgePos := iy*dst.Stride + ix
				dst.Pix[dstPos] = dst.Pix[edgePos]
			}
		}
	}
	parallel.Line(paddedHeight, padding)
	return dst
}

// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
//
// +---+---+---+---+        +---+---+---+---+    +---+---+---+---+    +---+---+---+---+
// | # | # | # | # |        | # | # | # | # |    | # | # | # | # |    | # | # | # | # |
// +---+---+---+---+        +---+---+---+---+    +---+---+---+---+    +---+---+---+---+
// | # | # | # | # |        | # | # | # | # |    | # | # | # | # |    | # | # | # | # |
// +---+---+---+---+ -----> +---+---+---+---+    +---+---+---+---+    +---+---+---+---+
// | # | # | # | # |        | # | # | # | # |    | # | # | # | # |    | # | # | # | # |
// +---+---+---+---+        +---+---+---+---+    +---+---+---+---+    +---+---+---+---+
// | # | # | # | # |        | # | # | # | # |    | # | # | # | # |    | # | # | # | # |
// +---+---+---+---+        +---+---+---+---+    +---+---+---+---+    +---+---+---+---+
//
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+
// 												 | # | # | # | # |
// 												 +---+---+---+---+

func wrap(srcImg image.Image, padX, padY int) *image.Gray {
	dst := nofill(srcImg, padX, padY)
	paddedWidth := dst.Bounds().Dx()
	paddedHeight := dst.Bounds().Dy()
	padding := func(start, end int) {
		for y := start; y < end; y++ {
			iy := y
			if iy < padY {
				iy = (paddedHeight - padY) - ((padY - y) % (paddedHeight - padY*2))
			} else if iy >= paddedHeight-padY {
				iy = padY - ((padY - y) % (paddedHeight - padY*2))
			}

			for x := 0; x < paddedWidth; x++ {
				ix := x
				if ix < padX {
					ix = (paddedWidth - padX) - ((padX - x) % (paddedWidth - padX*2))
				} else if ix >= paddedWidth-padX {
					ix = padX - ((padX - x) % (paddedWidth - padX*2))
				} else if iy == y {
					// This only enters if we are not in a y-padded area or
					// x-padded area, so nothing to extend here.
					// So simply jump to the next padded-x index.
					x = paddedWidth - padX - 1
					continue
				}
				dstPos := y*dst.Stride + x
				edgePos := iy*dst.Stride + ix
				dst.Pix[dstPos] = dst.Pix[edgePos]
			}
		}
	}
	parallel.Line(paddedHeight, padding)
	return dst
}

/* Package uils provides various useful helper functions. */
package utils

import (
	"fmt"
	"image"
	"image/color"
)

func SortGray(data []color.Gray, min, max int) {
	if min > max {
		return
	}
	p := partitionGraySlice(data, min, max)
	SortGray(data, min, p-1)
	SortGray(data, p+1, max)
}

func partitionGraySlice(data []color.Gray, min, max int) int {
	pivot := data[max]
	i := min
	r := pivot
	for j := min; j < max; j++ {
		if data[j].Y <= r.Y {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}
	data[i], data[max] = data[max], data[i]
	return i
}

// RGBAToString returns a string representation of the Hex values contained in an image.RGBA.
func RGBAToString(img *image.RGBA) string {
	var result string
	result += fmt.Sprintf("\nBounds: %v", img.Bounds())
	result += fmt.Sprintf("\nStride: %v", img.Stride)
	for y := 0; y < img.Bounds().Dy(); y++ {
		result += "\n"
		for x := 0; x < img.Bounds().Dx(); x++ {
			pos := y * img.Stride + x*4
			result += fmt.Sprintf("%#X, ", img.Pix[pos+0])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+1])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+2])
			result += fmt.Sprintf("%#X, ", img.Pix[pos+3])

		}
	}
	result += "\n"
	return result
}

func GrayToString(img *image.Gray) string {
	var result string
	result += fmt.Sprintf("\nBounds: %v", img.Bounds())
	result += fmt.Sprintf("\nStride: %v", img.Stride)
	for y := 0; y < img.Bounds().Dy(); y++ {
		result += "\n"
		for x := 0; x < img.Bounds().Dx(); x++ {
			pos := y * img.Stride + x
			result += fmt.Sprintf("%#X, ", img.Pix[pos])

		}
	}
	result += "\n"
	return result
}

// RGBASlicesEqual returns true if the parameter RGBA color slices a and b match
// or false if otherwise.
func RGBASlicesEqual(a, b []color.RGBA) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func GraySlicesEqual(a, b []color.Gray) bool {
	if a == nil && b == nil {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// GrayImageEqual returns true if the parameter images a and b match
// or false if otherwise.
func GrayImageEqual(a, b *image.Gray) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}
	for i := 0; i < len(a.Pix); i++ {
		if a.Pix[i] != b.Pix[i] {
			return false
		}
	}
	return true
}

// RGBAImageEqual returns true if the parameter images a and b match
// or false if otherwise.
func RGBAImageEqual(a, b *image.RGBA) bool {
	if !a.Rect.Eq(b.Rect) {
		return false
	}
	for y := 0; y < a.Bounds().Dy(); y++ {
		for x := 0; x < a.Bounds().Dx(); x++ {
			pos := y * a.Stride + x * 4
			if a.Pix[pos+0] != b.Pix[pos+0] {
				return false
			}
			if a.Pix[pos+1] != b.Pix[pos+1] {
				return false
			}
			if a.Pix[pos+2] != b.Pix[pos+2] {
				return false
			}
			if a.Pix[pos+3] != b.Pix[pos+3] {
				return false
			}
		}
	}
	return true
}

func Clamp(value, min, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}
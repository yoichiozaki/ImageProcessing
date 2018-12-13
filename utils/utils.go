/* Package uils provides various useful helper functions. */
package utils

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"os"
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
			pos := y*img.Stride + x*4
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
			pos := y*img.Stride + x
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
			pos := y*a.Stride + x*4
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

func GaussElimination(equations *[][]float64) {

	N := len(*equations)

	for k := 0; k < N; k++ {
		max := 0.0
		s := k
		for j := k; j < N; j++ {
			if math.Abs((*equations)[j][k]) > max {
				max = math.Abs((*equations)[j][k])
				s = j
			}
		}
		if max == 0 {
			log.Println("can not solve this equation.")
			os.Exit(1)
		}
		for j := 0; j <= N; j++ {
			(*equations)[k][j], (*equations)[s][j] = (*equations)[s][j], (*equations)[k][j]
		}
		p := (*equations)[k][k]
		for j := k; j < N+1; j++ {
			(*equations)[k][j] /= p
		}
		for i := 0; i < N; i++ {
			if i != k {
				d := (*equations)[i][k]
				for j := k; j < N+1; j++ {
					(*equations)[i][j] -= d * (*equations)[k][j]
				}
			}
		}
	}

}

func SampleGaussElimination() {
	const N = 4 // 変数個数
	a := [][]float64{
		{1.0, -2.0, 3.0, -4.0, 5.0},
		{-2.0, 5.0, 8.0, -3.0, 9.0},
		{5.0, 4.0, 7.0, 1.0, -1.0},
		{9.0, 7.0, 3.0, 5.0, 4.0},
	}

	for k := 0; k < N-1; k++ {
		for i := k + 1; i < N; i++ {
			d := a[i][k] / a[k][k]
			fmt.Println(d)
			for j := k + 1; j <= N; j++ {
				a[i][j] -= a[k][j] * d
			}
		}
	}

	for i := N - 1; i >= 0; i-- {
		d := a[i][N]
		for j := i + 1; j < N; j++ {
			d -= a[i][j] * a[j][N]
		}
		a[i][N] = d / a[i][i]
	}

	for k := 0; k < N; k++ {
		fmt.Printf("x_%d = %f\n", k, a[k][N])
	}
}

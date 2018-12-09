package convolution

import (
	"fmt"
	"math"
)

type Matrix interface {
	At(x, y int) float64         // At returns the matrix value at (x, y).
	GetNormalizedMatrix() Matrix // GetNormalizedMatrix returns a new normalized matrix.
	MaxX() int                   // MaxX returns the horizontal length of the matrix.
	MaxY() int                   // MaxY returns the vertical length of the matrix.
	String() string              // String returns the string representation of the matrix.
}

// New returns a kernel with provided size.
func New(width, height int) *Kernel {
	return &Kernel{
		Matrix: make([]float64, width*height),
		Width:  width,
		Height: height,
	}
}

// Kernel to bu used as a convolution matrix.
type Kernel struct {
	Matrix []float64
	Width  int
	Height int
}

func (k *Kernel) GetNormalizedMatrix() Matrix {
	absum := absum(k.Matrix)
	normalized := New(k.Width, k.Height)
	if absum == 0 {
		absum = 1
	}
	for i := 0; i < k.Width*k.Height; i++ {
		normalized.Matrix[i] = k.Matrix[i] / absum
	}
	fmt.Println(normalized)
	return normalized
}

func (k *Kernel) MaxX() int {
	return k.Width
}

func (k *Kernel) MaxY() int {
	return k.Height
}

func (k *Kernel) At(x, y int) float64 {
	return k.Matrix[y*k.Width+x]
}

func (k *Kernel) String() string {
	result := ""
	for y := 0; y < k.MaxY(); y++ {
		result += fmt.Sprintf("\n")
		for x := 0; x < k.MaxX(); x++ {
			result += fmt.Sprintf("%-8.4f", k.At(x, y))
		}
	}
	return result
}

func absum(matrix []float64) float64 {
	var sum float64
	for _, v := range matrix {
		sum += math.Abs(v)
	}
	return sum
}

package noise

import (
	"ImageProcessing/clone"
	"ImageProcessing/parallel"
	"image"
	"math/rand"
	"time"
)

type Generator func() uint8

var (
	Uniform  Generator
	Binary   Generator
	Gaussian Generator
	Spike    Generator
)

func init() {
	Uniform = func() uint8 {
		return uint8(rand.Intn(256))
	}
	Binary = func() uint8 {
		return 0xFF * uint8(rand.Intn(2))
	}
	Gaussian = func() uint8 {
		return uint8(rand.NormFloat64()*32.0 + 128.0)
	}
	Spike =func() uint8 {
		if i := rand.Intn(10); i == 7 {
			return 0xFF * uint8(rand.Intn(2))
		} else {
			return 0
		}
	}
}

func GenerateNoiseImage(width, height int, generator Generator) *image.Gray {
	dst := image.NewGray(image.Rect(0, 0, width, height))
	rand.Seed(time.Now().UTC().UnixNano())
	parallel.Line(height, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < width; x++ {
				pos := y*dst.Stride + x
				dst.Pix[pos] = generator()
			}
		}
	})
	return dst
}

func GenerateSpikeNoiseOn(src image.Image) *image.Gray {
	dst := clone.AsGray(src)
	rand.Seed(time.Now().UTC().UnixNano())
	parallel.Line(dst.Bounds().Dy(), func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < dst.Bounds().Dx(); x++ {
				pos := y*dst.Stride + x
				if noise := Spike(); noise == 0 {
					dst.Pix[pos] = dst.Pix[pos]
				} else {
					dst.Pix[pos] = noise
				}
			}
		}
	})
	return dst
}
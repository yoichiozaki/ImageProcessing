package histogram

import (
	"ImageProcessing/clone"
	"image"
)

type Histogram struct {
	Bins []int
}

type GrayHistogram struct {
	Y Histogram
}

func (h *Histogram) Max() int {
	var max int
	if len(h.Bins) > 0 {
		max = h.Bins[0]
		for i := 1; i < len(h.Bins); i++ {
			if h.Bins[i] > max {
				max = h.Bins[i]
			}
		}
	}
	return max
}

func (h *Histogram) Min() int {
	var min int
	if len(h.Bins) > 0 {
		min = h.Bins[0]
		for i := 1; i < len(h.Bins); i++ {
			if h.Bins[i] < min {
				min = h.Bins[i]
			}
		}
	}
	return min
}

// 累積
func (h *Histogram) Cumulate() *Histogram {
	binLength := len(h.Bins)
	ret := &Histogram{make([]int, binLength)}
	if binLength > 0 {
		ret.Bins[0] = h.Bins[0]
	}
	for i := 1; i < binLength; i++ {
		ret.Bins[i] = ret.Bins[i-1] + h.Bins[i]
	}
	return ret
}

func (h *Histogram) Dump() *image.Gray {
	dstW, dstH := len(h.Bins), len(h.Bins)
	dst := image.NewGray(image.Rect(0,0,dstW, dstH))
	max := h.Max()
	if max == 0 {
		max = 1
	}
	for x := 0; x < dstW; x++ {
		value := ((int(h.Bins[x]) << 16 / max) * dstH) >> 16
		for y := dstH - 1; y > dstH-value-1; y-- {
			dst.Pix[y*dst.Stride+x] = 0xFF
		}
	}
	return dst
}

func GetGrayHistogram(img image.Image) *GrayHistogram {
	src := clone.AsGray(img)
	binLength := 256
	grayHistogram := Histogram{make([]int, binLength)}
	for y := 0; y < src.Bounds().Dy(); y++ {
		for x := 0; x < src.Bounds().Dx(); x++ {
			pos := y*src.Stride+x
			grayHistogram.Bins[src.Pix[pos]]++
		}
	}
	return &GrayHistogram{grayHistogram}
}
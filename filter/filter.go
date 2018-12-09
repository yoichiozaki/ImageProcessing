/* Package filter provides various filtering functions. */
package filter

import (
	"ImageProcessing/clone"
	"ImageProcessing/convolution"
	"ImageProcessing/histogram"
	"ImageProcessing/parallel"
	"ImageProcessing/utils"
	"image"
	"image/color"
	"log"
	"math"
)

// 実装するのは
// 空間フィルタリング
// 		[-] 平滑化
// 		[-] ラプラシアンによるエッジ検出
// 		[-] 鮮鋭化
// ノイズ除去
//		[-] 平均化
//		[-] メディアンフィルタ
//		[ ] バイラテラルフィルタ
// ハーフトーン処理
//		[ ] ディザ法
//		[ ] 誤差拡散法
// 濃度変換
//		[-] ガンマ変換
// 		[-] ヒストグラム平坦化

func Grayscale(img image.Image) *image.Gray {
	src := clone.AsRGBA(img)
	bounds := src.Bounds()
	srcW, srcH := bounds.Dx(), bounds.Dy()
	if bounds.Empty() {
		return &image.Gray{}
	}
	dst := image.NewGray(bounds)
	converter := func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < srcW; x++ {
				srcPos := y*src.Stride + x*4
				dstPos := y*dst.Stride + x
				c := 0.3*float64(src.Pix[srcPos+0]) + 0.6*float64(src.Pix[srcPos+1]) + 0.1*float64(src.Pix[srcPos+2])
				dst.Pix[dstPos] = uint8(c + 0.5)
			}
		}
	}
	parallel.Line(srcH, converter)
	return dst
}

func BoxBlur(src image.Image, radius int) *image.Gray {
	if radius <= 0 {
		return clone.AsGray(src)
	}
	length := radius
	kernel := convolution.New(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			kernel.Matrix[y*length+x] = 1
		}
	}
	return convolution.Convolve(src, kernel.GetNormalizedMatrix(), &convolution.Options{Wrap: false})
}

func GaussianBlur(src image.Image, radius int) *image.Gray {
	if radius <= 0 {
		return clone.AsGray(src)
	}
	length := radius
	kernel := convolution.New(length, length)
	gaussian := func(x, y, sigma float64) float64 {
		return math.Exp(-(x*x/sigma + y*y/sigma))
	}
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			kernel.Matrix[y*length+x] = gaussian(float64(x)-float64(radius), float64(y)-float64(radius), 4*float64(radius))
		}
	}
	return convolution.Convolve(src, kernel.GetNormalizedMatrix(), &convolution.Options{Wrap: false})
}

func EdgeDetection(src image.Image, radius int) *image.Gray {
	if radius <= 0 {
		return image.NewGray(src.Bounds())
	}
	length := radius
	kernel := convolution.New(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			v := -1.0
			if x == length/2 && y == length/2 {
				v = float64(length*length) - 1
			}
			kernel.Matrix[y*length+x] = v
		}
	}
	return convolution.Convolve(src, kernel, &convolution.Options{Wrap: false})
}

func FixedEdgeDetection(src image.Image) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			0, 1, 0,
			1, -4, 1,
			0, 1, 0,
		},
		Width:  3,
		Height: 3,
	}
	return convolution.Convolve(src, &kernel, &convolution.Options{Wrap: false})

}

func Sharpen(src image.Image, k float64) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			0, -k, 0,
			-k, 1 + 4*k, -k,
			0, -k, 0,
		},
		Width:  3,
		Height: 3,
	}
	return convolution.Convolve(src, &kernel, &convolution.Options{Wrap: false})
}

func FixedSharpen(src image.Image) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			0, -1, 0,
			-1, 5, -1,
			0, -1, 0,
		},
		Width:  3,
		Height: 3,
	}
	return convolution.Convolve(src, &kernel, &convolution.Options{Wrap: false})
}

// Median returns a new image in which each pixel is the median of its neighbors.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
func Median(img image.Image, radius int) *image.Gray {
	fn := func(neighbors []color.Gray) color.Gray {
		utils.SortGray(neighbors, 0, len(neighbors)-1)
		return neighbors[len(neighbors)/2]
	}
	result := spatialFilter(img, radius, fn)
	return result
}

// Dilate picks the local maxima from the neighbors of each pixel and returns the resulting image.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
func Dilate(img image.Image, radius int) *image.Gray {
	fn := func(neighbors []color.Gray) color.Gray {
		utils.SortGray(neighbors, 0, len(neighbors)-1)
		return neighbors[len(neighbors)-1]
	}
	result := spatialFilter(img, radius, fn)
	return result
}

// Erode picks the local minima from the neighbors of each pixel and returns the resulting image.
// The parameter radius corresponds to the radius of the neighbor area to be searched,
// for example a radius of R will result in a search window length of 2R+1 for each dimension.
func Erode(img image.Image, radius int) *image.Gray {
	fn := func(neighbors []color.Gray) color.Gray {
		utils.SortGray(neighbors, 0, len(neighbors)-1)
		return neighbors[0]
	}
	result := spatialFilter(img, radius, fn)
	return result
}

func spatialFilter(img image.Image, radius int, picker func(neighbors []color.Gray) color.Gray) *image.Gray {
	if radius <= 0 {
		return clone.AsGray(img)
	}
	padding := radius-1
	src := clone.Pad(img, padding, padding, clone.EdgeExtend)
	kernelSize := radius
	bounds := img.Bounds()
	dst := image.NewGray(bounds)
	w, h := bounds.Dx(), bounds.Dy()
	neighborsCount := kernelSize * kernelSize
	parallel.Line(h, func(start, end int) {
		for y := start + padding; y < end+padding; y++ {
			for x := padding; x < w+padding; x++ {
				neighbors := make([]color.Gray, neighborsCount)
				i := 0
				for ky := 0; ky < kernelSize; ky++ {
					for kx := 0; kx < kernelSize; kx++ {
						ix := x - kernelSize>>1 + kx
						iy := y - kernelSize>>1 + ky
						ipos := iy*src.Stride + ix
						neighbors[i] = color.Gray{
							Y: src.Pix[ipos],
						}
						i++
					}
				}
				c := picker(neighbors)
				pos := (y-padding)*dst.Stride + (x-padding)
				dst.Pix[pos] = c.Y
			}
		}
	})
	return dst
}

func Gamma(src image.Image, gamma float64) *image.Gray {
	gamma = math.Max(0.0000001, gamma)
	fn := func(original uint8) uint8 {
		return uint8(utils.Clamp(math.Pow(float64(original)/255, 1.0/gamma)*255, 0, 255))
	}
	bounds := src.Bounds()
	dst := clone.AsGray(src)
	w, h := bounds.Dx(), bounds.Dy()
	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				dstPos := y*dst.Stride + x
				dst.Pix[dstPos] = fn(dst.Pix[dstPos])
			}
		}
	})
	return dst
}

func HistogramEqualization(src image.Image) *image.Gray {
	bounds := src.Bounds()
	dst := clone.AsGray(src)
	w, h := bounds.Dx(), bounds.Dy()
	hist := histogram.GetGrayHistogram(src)
	bins := make([]color.Gray, h*dst.Stride + w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			bins[i*dst.Stride + j], _ = color.GrayModel.Convert(src.At(j, i)).(color.Gray)
		}
	}
	bins = *utils.SortGray_(&bins, 0, len(hist.Y.Bins)-1)
	var max uint8 = 255
	log.Println(max)
	cumulated := hist.Y.Cumulate()
	fn := func(original uint8) uint8 {
		return max*uint8(cumulated.Bins[original])
	}
	parallel.Line(h, func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < w; x++ {
				dstPos := y*dst.Stride + x
				dst.Pix[dstPos] = fn(dst.Pix[dstPos])
			}
		}
	})
	return dst
}

/* Package filter provides various filtering functions. */
package filter

import (
	"ImageProcessing/clone"
	"ImageProcessing/convolution"
	"ImageProcessing/histogram"
	"ImageProcessing/parallel"
	"ImageProcessing/utils"
	"fmt"
	"image"
	"image/color"
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
//		[-] バイラテラルフィルタ
// ハーフトーン処理
//		[-] ディザ法
//		[-] 誤差拡散法
// 濃度変換
//		[-] ガンマ変換
// 		[-] ヒストグラム平坦化
//
// サンプルイメージ: https://homepages.cae.wisc.edu/~ece533/images/

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
				dstPos := y*dst.Stride + x
				c := color.GrayModel.Convert(src.At(x, y))
				gray, _ := c.(color.Gray)
				dst.Pix[dstPos] = gray.Y
			}
		}
	}
	parallel.Line(srcH, converter)
	return dst
}

func BoxBlur(img image.Image, radius int) *image.Gray {
	if radius <= 0 {
		return clone.AsGray(img)
	}
	length := radius
	kernel := convolution.New(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			kernel.Matrix[y*length+x] = 1
		}
	}
	return convolution.Convolve(img, kernel.GetNormalizedMatrix(), &convolution.Options{Wrap: false})
}

func FixedDirectionBlur(img image.Image) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
		},
		Width:  15,
		Height: 15,
	}
	return convolution.Convolve(img, kernel.GetNormalizedMatrix(), &convolution.Options{Wrap: false})

}

func GaussianBlur(img image.Image, radius int) *image.Gray {
	if radius <= 0 {
		return clone.AsGray(img)
	}
	length := radius
	kernel := convolution.New(length, length)
	for x := 0; x < length; x++ {
		for y := 0; y < length; y++ {
			kernel.Matrix[y*length+x] = gaussian2D(float64(x)-float64(radius), float64(y)-float64(radius), 4*float64(radius))
		}
	}
	return convolution.Convolve(img, kernel.GetNormalizedMatrix(), &convolution.Options{Wrap: false})
}

func EdgeDetection(img image.Image, radius int) *image.Gray {
	if radius <= 0 {
		return image.NewGray(img.Bounds())
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
	fmt.Println(kernel.String())
	return convolution.Convolve(img, kernel, &convolution.Options{Wrap: false})
}

func FixedEdgeDetection(img image.Image) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			0, 1, 0,
			1, -4, 1,
			0, 1, 0,
		},
		Width:  3,
		Height: 3,
	}
	return convolution.Convolve(img, &kernel, &convolution.Options{Wrap: false})
}

func Sharpen(img image.Image, k float64) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			0, -k, 0,
			-k, 1 + 4*k, -k,
			0, -k, 0,
		},
		Width:  3,
		Height: 3,
	}
	return convolution.Convolve(img, &kernel, &convolution.Options{Wrap: false})
}

func FixedSharpen(img image.Image) *image.Gray {
	kernel := convolution.Kernel{
		Matrix: []float64{
			0, -1, 0,
			-1, 5, -1,
			0, -1, 0,
		},
		Width:  3,
		Height: 3,
	}
	return convolution.Convolve(img, &kernel, &convolution.Options{Wrap: false})
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
	padding := radius - 1
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
				pos := (y-padding)*dst.Stride + (x - padding)
				dst.Pix[pos] = c.Y
			}
		}
	})
	return dst
}

// TODO: パラメータの調整が必要
func BilateralFilter(img image.Image, radius int, sigma_i, sigma_s float64) *image.Gray {
	if radius <= 0 {
		return clone.AsGray(img)
	}
	src := clone.AsGray(img)
	dst := image.NewGray(img.Bounds())
	for i := 0; i < dst.Bounds().Dy(); i++ {
		for j := 0; j < dst.Bounds().Dx(); j++ {
			bilateral(src, dst, i, j, radius, sigma_i, sigma_s)
		}
	}
	return dst
}

func bilateral(source *image.Gray, dest *image.Gray, y, x, radius int, sigma_i, sigma_s float64) {
	hl := radius / 2
	i_filtered := 0.0
	Wp := 0.0
	for i := 0; i < radius; i++ {
		for j := 0; j < radius; j++ {
			neighbor_x := x - (hl - i)
			neighbor_y := y - (hl - j)
			if dest.Bounds().Dx() <= neighbor_x {
				neighbor_x -= source.Bounds().Dx()
			}
			if dest.Bounds().Dy() <= neighbor_y {
				neighbor_y -= source.Bounds().Dy()
			}
			gi := gaussian(float64(source.GrayAt(neighbor_x, neighbor_y).Y-source.GrayAt(x, y).Y), sigma_i)
			gs := gaussian(distance(float64(neighbor_x), float64(neighbor_y), float64(x), float64(y)), sigma_s)
			w := gi * gs
			i_filtered += float64(source.GrayAt(neighbor_x, neighbor_y).Y) * w
			Wp += w
		}
	}
	i_filtered /= Wp
	dest.Pix[y*dest.Stride+x] = uint8(i_filtered)
}

func Gamma(img image.Image, gamma float64) *image.Gray {
	gamma = math.Max(0.0000001, gamma)
	fn := func(original uint8) uint8 {
		return uint8(utils.Clamp(math.Pow(float64(original)/255, 1.0/gamma)*255, 0, 255))
	}
	bounds := img.Bounds()
	dst := clone.AsGray(img)
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

func HistogramEqualization(img image.Image) *image.Gray {
	bounds := img.Bounds()
	src := clone.AsGray(img)
	dst := clone.AsGray(img)
	w, h := bounds.Dx(), bounds.Dy()
	hist := histogram.GetGrayHistogram(img)
	bins := make([]color.Gray, h*w)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			bins[i*dst.Stride+j] = src.GrayAt(j, i)
		}
	}
	utils.SortGray(bins, 0, len(bins)-1)
	max := bins[len(bins)-1].Y
	cumulated := hist.Y.Cumulate()
	fn := func(original uint8) uint8 {
		// log.Print(int(max) * cumulated.Bins[original]/(w*h))
		return uint8(int(max) * cumulated.Bins[original]/(w*h))
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

func HalftoningWithDitheringMethod(img image.Image) *image.Gray {
	Bayer := []uint{
		0, 8, 2, 10,
		12, 4, 14, 6,
		3, 11, 1, 9,
		15, 7, 3, 5,
	}
	src := clone.AsGray(img)
	bounds := src.Bounds()
	dst := image.NewGray(bounds)
	fn := func(x, y int) uint8 {
		if uint8(Bayer[(y%4)*4 + (x%4)]*16 + 8) <= src.Pix[y*dst.Stride + x] {
			return 255
		} else {
			return 0
		}
	}
	parallel.Line(bounds.Dy(), func(start, end int) {
		for y := start; y < end; y++ {
			for x := 0; x < bounds.Dx(); x++ {
				dstPos := y*dst.Stride + x
				dst.Pix[dstPos] = fn(x, y)
			}
		}
	})
	// for y := 0; y < bounds.Dy(); y++ {
	// 	for x := 0; x < bounds.Dx(); x++ {
	// 		if uint8(Bayer[(y%4)*4 + (x%4)]*16 + 8) <= src.GrayAt(x, y).Y {
	// 			dst.Pix[y*dst.Stride + x] = 255
	// 		} else {
	// 			dst.Pix[y*dst.Stride + x] = 0
	// 		}
	// 	}
	// }
	return dst
}

func HalftoningWithErrorDiffusionMethod(img image.Image) *image.Gray {
	src := clone.AsGray(img)
	bounds := src.Bounds()
	dst := clone.AsGray(img)

	// 並行処理すると切れ目ができちゃう
	// parallel.Line(bounds.Dy(), func(start, end int) {
	// 	var e uint8
	// 	for y := start; y < end; y++ {
	// 		for x := 0; x < bounds.Dx(); x++ {
	// 			dstPos := y*dst.Stride + x
	// 			f := dst.GrayAt(x, y).Y
	// 			if f < 127 {
	// 				dst.Pix[dstPos] = 255
	// 				e = f - 255
	// 			} else {
	// 				dst.Pix[dstPos] = 0
	// 				e = f
	// 			}
	// 			if x != dst.Bounds().Dx()-1 {
	// 				dst.Pix[dstPos+1] += uint8((5.0/16)*float64(e))
	// 				// dst.Pix[dstPos+1] += (7/16)*e
	// 			}
	// 			if x != 0 && y != dst.Bounds().Dy()-1 {
	// 				dst.Pix[(y+1)*dst.Stride + x-1] += uint8((3.0/16)*float64(e))
	// 				// dst.Pix[(y+1)*dst.Stride + x-1] += (3/16)*e
	// 			}
	// 			if y != dst.Bounds().Dy()-1 {
	// 				dst.Pix[(y+1)*dst.Stride + x] += uint8((5.0/16)*float64(e))
	// 				// dst.Pix[(y+1)*dst.Stride + x] += (5/16)*e
	// 			}
	// 			if x != dst.Bounds().Dx()-1 && y != dst.Bounds().Dy()-1 {
	// 				dst.Pix[(y+1)*dst.Stride + x+1] += uint8((3.0/16)*float64(e))
	// 				// dst.Pix[(y+1)*dst.Stride + x+1] += (1/16)*e
	// 			}
	// 		}
	// 	}
	// })

	var e uint8
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			dstPos := y*dst.Stride + x
			f := dst.GrayAt(x, y).Y
			if f < 127 {
				dst.Pix[dstPos] = 255
				e = f - 255
			} else {
				dst.Pix[dstPos] = 0
				e = f
			}
			if x != dst.Bounds().Dx()-1 {
				dst.Pix[dstPos+1] += uint8((5.0/16)*float64(e))
				// dst.Pix[dstPos+1] += (7/16)*e
			}
			if x != 0 && y != dst.Bounds().Dy()-1 {
				dst.Pix[(y+1)*dst.Stride + x-1] += uint8((3.0/16)*float64(e))
				// dst.Pix[(y+1)*dst.Stride + x-1] += (3/16)*e
			}
			if y != dst.Bounds().Dy()-1 {
				dst.Pix[(y+1)*dst.Stride + x] += uint8((5.0/16)*float64(e))
				// dst.Pix[(y+1)*dst.Stride + x] += (5/16)*e
			}
			if x != dst.Bounds().Dx()-1 && y != dst.Bounds().Dy()-1 {
				dst.Pix[(y+1)*dst.Stride + x+1] += uint8((3.0/16)*float64(e))
				// dst.Pix[(y+1)*dst.Stride + x+1] += (1/16)*e
			}
		}
	}
	return dst
}

func gaussian1D(x, sigma float64) float64 {
	return math.Exp(-(x*x / sigma))
}
func gaussian2D(x, y, sigma float64) float64 {
	return math.Exp(-(x*x/sigma + y*y/sigma))
}

func gaussian(x, sigma float64) float64 {
	return (1.0/(2*math.Pi*sigma*sigma)) * math.Exp(-(x*x)/(2*sigma*sigma))
}

func distance(x, y, i, j float64) float64 {
	return math.Sqrt((x-i)*(x-i) + (y-j)*(y-j))
}
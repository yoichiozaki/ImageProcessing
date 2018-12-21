package transform

import (
	"ImageProcessing/clone"
	"image"
)

// 台形補正
// 出力画像のboundsを先に決めておいて、
// 出力画像の位置からそこに相当する入力画像の位置を決定していく
// 出力画像の位置から行列をかけて求められる入力画像の位置が整数値にならない時はガウスを取ってしまえばいい
// ガウスじゃないものを使うともっと正確に描写できるようになるかもしれない
// 出力画像の位置と入力画像の位置を変換するのは行列
func KeystoneEffect(img image.Image, matrix *[][]float64, output_x, output_y int) *image.Gray {
	src := clone.AsGray(img)
	dst := image.NewGray(image.Rect(0, 0, output_x, output_y))
	bounds := dst.Bounds()
	fn := func(x, y int) (int, int) {
		u := (*matrix)[0][0]*float64(x) + (*matrix)[0][1]*float64(y) + (*matrix)[0][2]*float64(1)
		v := (*matrix)[1][0]*float64(x) + (*matrix)[1][1]*float64(y) + (*matrix)[1][2]*float64(1)
		s := (*matrix)[2][0]*float64(x) + (*matrix)[2][1]*float64(y) + (*matrix)[2][2]*float64(1)
		return int(u/s), int(v/s)
	}
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			dstPos := y*dst.Stride + x
			srcX, srcY := fn(x, y)
			dst.Pix[dstPos] = src.GrayAt(srcX, srcY).Y
		}
	}
	return dst
}

func GetUnion(img0 image.Image, img1 image.Image) *image.Gray {
	src0 := clone.AsGray(img0)
	src1 := clone.AsGray(img1)
	bounds := src0.Rect
	dst := image.NewGray(bounds)
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			if src0.Pix[y*src0.Stride+x] < 85 {
				dst.SetGray(x, y, src1.GrayAt(x, y)) // 置換
			} else {
				dst.SetGray(x, y, src0.GrayAt(x, y))
			}
		}
	}
	return dst
}

func Stitching(img0 image.Image, img1 image.Image, img2 image.Image, margin int, matrix01, matrix02 *[][]float64) *image.Gray {
	src0 := clone.AsGray(img0)
	src1 := clone.AsGray(img1)
	src2 := clone.AsGray(img2)
	src0Bounds := src0.Bounds()
	src1Bounds := src1.Bounds()
	src2Bounds := src2.Bounds()
	dstBounds := image.Rectangle{
		Min: struct{ X, Y int }{X: 0, Y: 0},
		Max: struct{ X, Y int }{X: src0Bounds.Dx()+margin*2, Y: src0Bounds.Dy()+margin*2},
	}
	fn01 := func(u, v int) (int, int) {
		x := (*matrix01)[0][0]*float64(u) + (*matrix01)[0][1]*float64(v) + (*matrix01)[0][2]*float64(1)
		y := (*matrix01)[1][0]*float64(u) + (*matrix01)[1][1]*float64(v) + (*matrix01)[1][2]*float64(1)
		s := (*matrix01)[2][0]*float64(u) + (*matrix01)[2][1]*float64(v) + (*matrix01)[2][2]*float64(1)
		return int(x/s), int(y/s)
	}
	fn02 := func(u, v int) (int, int) {
		x := (*matrix02)[0][0]*float64(u) + (*matrix02)[0][1]*float64(v) + (*matrix02)[0][2]*float64(1)
		y := (*matrix02)[1][0]*float64(u) + (*matrix02)[1][1]*float64(v) + (*matrix02)[1][2]*float64(1)
		s := (*matrix02)[2][0]*float64(u) + (*matrix02)[2][1]*float64(v) + (*matrix02)[2][2]*float64(1)
		return int(x/s), int(y/s)
	}
	dst := image.NewGray(dstBounds)
	for y1 := 0; y1 < src1Bounds.Dy(); y1++ {
		for x1 := 0; x1 < src1Bounds.Dx(); x1++ {
			x0, y0 := fn01(x1, y1) // 変換後
			dst.SetGray(x0+margin, y0+margin, src1.GrayAt(x1, y1))
		}
	}
	for y2 := 0; y2 < src2Bounds.Dy(); y2++ {
		for x2 := 0; x2 < src2Bounds.Dx(); x2++ {
			x0, y0 := fn02(x2, y2) // 変換後
			dst.SetGray(x0+margin, y0+margin, src2.GrayAt(x2, y2))
		}
	}
	for i := 0; i < src0Bounds.Dy(); i++ {
		for j := 0; j < src0Bounds.Dx(); j++ {
			dst.SetGray(j+margin, i+margin, src0.GrayAt(j, i))
		}
	}
	return dst
}

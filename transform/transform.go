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

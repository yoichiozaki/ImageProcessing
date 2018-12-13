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
func KeystoneEffect(img image.Image, matrix *[][]float64) *image.Gray {
	src := clone.AsGray(img)
	dst := image.NewGray(image.Rect(0, 0, 400, 300))
	bounds := dst.Bounds()
	fn := func(x, y int) int {
		u := (*matrix)[0][0]*float64(x) + (*matrix)[0][1]*float64(y) + (*matrix)[0][2]*float64(1)
		v := (*matrix)[1][0]*float64(x) + (*matrix)[1][1]*float64(y) + (*matrix)[1][2]*float64(1)
		return int(v)*src.Stride + int(u)
	}
	for y := 0; y < bounds.Dy(); y++ {
		for x := 0; x < bounds.Dx(); x++ {
			dstPos := y*dst.Stride + x
			srcPos := fn(x, y)
			dst.Pix[dstPos] = src.Pix[srcPos]
		}
	}
	return dst
}

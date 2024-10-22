package clone

import (
	"image"
	"image/color"
	"testing"
	"ImageProcessing/utils"
)

func TestCloneAsRGBA(t *testing.T) {
	cases := []struct {
		desc     string
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc: "RGBA",
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "RGBA64",
			value: &image.RGBA64{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "NRGBA",
			value: &image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0x80,
					0xFF, 0xFF, 0xFF, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "NRGBA64",
			value: &image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 8,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x80,
					0xFF, 0xFF, 0xFF, 0x80, 0xFF, 0xFF, 0xFF, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "Gray",
			value: &image.Gray{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 2,
				Pix: []uint8{
					0x80, 0x80,
					0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0xFF,
				},
			},
		},
		{
			desc: "Gray16",
			value: &image.Gray16{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 2,
				Pix: []uint8{
					0x80, 0x80,
					0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0xFF,
				},
			},
		},
		{
			desc: "Alpha",
			value: &image.Alpha{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1,
				Pix: []uint8{
					0x80,
					0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "Alpha16",
			value: &image.Alpha16{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1,
				Pix: []uint8{
					0x80, 0x80,
					0x80, 0x80,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0x80, 0x80, 0x80, 0x80,
				},
			},
		},
		{
			desc: "Paletted",
			value: &image.Paletted{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1,
				Palette: color.Palette{
					color.RGBA{0x00, 0x00, 0x00, 0x00},
					color.RGBA{0x80, 0x80, 0x80, 0x80},
					color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
				},
				Pix: []uint8{
					0x1, 0x2,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0x80,
					0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := AsRGBA(c.value)
		if !utils.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "CloneAsRGBA from "+c.desc, c.expected, actual)
		}
	}
}

func TestPad(t *testing.T) {
	cases := []struct {
		desc     string
		method   PaddingMethod
		x, y     int
		value    image.Image
		expected *image.RGBA
	}{
		{
			desc:   "No fill",
			method: NoFill,
			x:      2,
			y:      1,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 4,
				Pix: []uint8{
					0x80, 0x80, 0x80, 0xFF,
					0x40, 0x40, 0x40, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 5, 4),
				Stride: 5 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x80, 0x80, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x40, 0x40, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
			},
		},
		{
			desc:   "Edge Extend",
			method: EdgeExtend,
			x:      1,
			y:      1,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x40, 0x40, 0x40, 0xFF, 0x10, 0x10, 0x10, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x40, 0x40, 0x40, 0xFF, 0x40, 0x40, 0x40, 0xFF, 0x10, 0x10, 0x10, 0xFF, 0x10, 0x10, 0x10, 0xFF,
					0x40, 0x40, 0x40, 0xFF, 0x40, 0x40, 0x40, 0xFF, 0x10, 0x10, 0x10, 0xFF, 0x10, 0x10, 0x10, 0xFF,
				},
			},
		},
		{
			desc:   "Edge Wrap",
			method: EdgeWrap,
			x:      1,
			y:      1,
			value: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x40, 0x40, 0x40, 0xFF, 0x10, 0x10, 0x10, 0xFF,
				},
			},
			expected: &image.RGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x40, 0x40, 0x40, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x40, 0x40, 0x40, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0x40, 0x40, 0x40, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0x40, 0x40, 0x40, 0xFF, 0x80, 0x80, 0x80, 0xFF,
					0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x80, 0x80, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
		},
	}

	for _, c := range cases {
		actual := Pad(c.value, c.x, c.y, c.method)
		if !utils.RGBAImageEqual(actual, c.expected) {
			t.Errorf("%s:\nexpected:%v\nactual:%v", "Pad "+c.desc, utils.RGBAToString(c.expected), utils.RGBAToString(actual))
		}
	}
}

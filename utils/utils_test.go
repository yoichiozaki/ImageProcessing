package utils

import (
	"image"
	"image/color"
	"testing"
)

func TestQuickSortGray(t *testing.T) {
	cases := []struct {
		value    []color.Gray
		expected []color.Gray
	}{
		{
			value:    []color.Gray{{0}},
			expected: []color.Gray{{0}},
		},
		{
			value:    []color.Gray{{1}, {0}, {0}},
			expected: []color.Gray{{0}, {0}, {1}},
		},
		{
			value:    []color.Gray{{128}, {255}, {0}},
			expected: []color.Gray{{0}, {128}, {255}},
		},
		{
			value:    []color.Gray{{128}, {255}, {0}, {2}, {5}},
			expected: []color.Gray{{0}, {2}, {5}, {128}, {255}},
		},
	}

	for _, c := range cases {
		SortGray(c.value, 0, len(c.value)-1)
		if !GraySlicesEqual(c.value, c.expected) {
			t.Errorf("%s: expected: %#v, actual: %#v", "SortRGBA", c.expected, c.value)
		}
	}
}

func TestRGBASlicesEqual(t *testing.T) {
	cases := []struct {
		a        []color.RGBA
		b        []color.RGBA
		expected bool
	}{
		{
			a:        []color.RGBA{},
			b:        []color.RGBA{},
			expected: true,
		},
		{
			a:        []color.RGBA{{}},
			b:        []color.RGBA{{}},
			expected: true,
		},
		{
			a:        []color.RGBA{{255, 140, 10, 0}},
			b:        []color.RGBA{{255, 140, 10, 0}},
			expected: true,
		},
		{
			a:        []color.RGBA{{255, 128, 10, 0}},
			b:        []color.RGBA{{255, 140, 10, 0}},
			expected: false,
		},
		{
			a:        []color.RGBA{{}},
			b:        []color.RGBA{{255, 140, 10, 0}},
			expected: false,
		},
		{
			a:        []color.RGBA{},
			b:        []color.RGBA{{255, 140, 10, 0}},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := RGBASlicesEqual(c.a, c.b)
		if actual != c.expected {
			t.Errorf("%s: expected: %v actual: %v", "RGBASlicesEqual", c.expected, actual)
		}
	}
}
func TestGrayImageEqual(t *testing.T) {
	cases := []struct {
		a        *image.Gray
		b        *image.Gray
		expected bool
	}{
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0xFF, 0x00, 0xFF,
					0xFF, 0xFF, 0xFF,
				},
			},
			expected: false,
		},
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			expected: false,
		},
		{
			a: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			expected: true,
		},
		{
			a:        &image.Gray{},
			b:        &image.Gray{},
			expected: true,
		},
		{
			a: &image.Gray{},
			b: &image.Gray{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF,
					0xFF, 0xFF,
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := GrayImageEqual(c.a, c.b)
		if actual != c.expected {
			t.Errorf("%s: expected: %v actual: %v", "GrayImageEqual", c.expected, actual)
		}
	}
}

func TestRGBAImageEqual(t *testing.T) {
	cases := []struct {
		a        *image.RGBA
		b        *image.RGBA
		expected bool
	}{
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: false,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: false,
		},
		{
			a: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 3, 2),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: true,
		},
		{
			a:        &image.RGBA{},
			b:        &image.RGBA{},
			expected: true,
		},
		{
			a: &image.RGBA{},
			b: &image.RGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2,
				Pix: []uint8{
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
					0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF,
				},
			},
			expected: false,
		},
	}

	for _, c := range cases {
		actual := RGBAImageEqual(c.a, c.b)
		if actual != c.expected {
			t.Errorf("%s: expected: %v actual: %v", "RGBAImageEqual", c.expected, actual)
		}
	}
}

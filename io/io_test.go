package io

import (
	"bytes"
	"image"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	cases := []struct{
		format string
		encoder Encoder
		value image.Image
	}{
		{
			format: "png",
			encoder: PNGEncoder(),
			value: &image.RGBA{
				Rect: image.Rect(0, 0, 3, 3),
				Stride: 3 * 4,
				Pix: []uint8{
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0x80, 0x00, 0x00, 0xFF,
					0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF, 0xFF, 0x00, 0x00, 0xFF,
				},
			},
		},
	}
	for _, c := range cases {
		buf := bytes.Buffer{}
		c.encoder(&buf, c.value)
		_, outputFormat, err := image.Decode(&buf)
		if err != nil {
			t.Error(err)
		}
		if !strings.Contains(c.format, outputFormat) {
			t.Errorf("Encoder: expected: %#v, actual: %#v", c.format, outputFormat)
		}
	}
}

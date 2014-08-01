package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"os"
	"testing"

	"github.com/nfnt/resize"
)

type mockImage struct{}

func (t mockImage) ColorModel() color.Model {
	return nil
}

func (t mockImage) Bounds() image.Rectangle {
	return image.Rectangle{}
}

func (t mockImage) At(x, y int) color.Color {
	return nil
}

func Test_resizeFile(t *testing.T) {
	filename := "test.file"
	open = func(s string) (*os.File, error) {
		if s != filename {
			t.Fatal("unexpected:", s)
		}
		return nil, nil
	}
	create = func(s string) (*os.File, error) {
		if s != "resized/"+filename {
			t.Fatal("unexpected:", s)
		}
		return nil, nil
	}
	decode = func(io.Reader) (image.Image, error) {
		return mockImage{}, nil
	}
	encode = func(w io.Writer, m image.Image, o *jpeg.Options) error {
		return nil
	}
	resizer = func(width, height uint, img image.Image, interp resize.InterpolationFunction) image.Image {
		return mockImage{}
	}
	resizeFile(filename, 10, 0)
}

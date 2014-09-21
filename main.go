package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"

	"path/filepath"

	"github.com/nfnt/resize"
)

const PATH = "resized"

var y = flag.Float64("y", 0, "desired height, 0 for auto, 0 > y > 1 for percentage")
var x = flag.Float64("x", 0, "desired width, 0 for auto, 0 > x > 1 for percentage")

func main() {
	flag.Parse()
	if err := validate(*x, *y, flag.Args()); err != nil {
		fmt.Println(err)
		usage()
		return
	}

	files, err := filepath.Glob(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Mkdir(PATH, 0777)

	for _, file := range files {
		resizeFile(file, *x, *y)
	}
}

func validate(x, y float64, args []string) error {
	if x == 0 && y == 0 {
		return errors.New("specify either x or y flag or both")
	}

	if x < 0 || y < 0 {
		return errors.New("x and y flags must be positive values")
	}

	if len(args) != 1 {
		return errors.New("specify filename or wildcard of filenames to resize")
	}

	return nil
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage example:\n\tresizer -x=0.3 *.jpg\n\n")
	fmt.Fprintf(os.Stderr, "All possible flags:\n")
	flag.CommandLine.PrintDefaults()

}

func resizeFile(filename string, xPx, yPx float64) {
	fmt.Println("Processing", filename)

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	img, imgFmt, err := image.Decode(file)
	check(err)

	if xPx < 1 {
		xPx = float64(img.Bounds().Size().X) * xPx
	}
	if yPx < 1 {
		yPx = float64(img.Bounds().Size().Y) * yPx
	}
	m := resize.Resize(uint(xPx), uint(yPx), img, resize.Bilinear)

	out, err := os.Create(PATH + "/" + filename)
	check(err)
	defer out.Close()

	if imgFmt == "jpeg" {
		err = jpeg.Encode(out, m, nil)
	} else if imgFmt == "gif" {
		err = gif.Encode(out, m, nil)
	} else if imgFmt == "png" {
		err = png.Encode(out, m)
	} else {
		fmt.Println("Unrecognized format:", imgFmt)
		return
	}
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

const PATH = "resized"

var y = flag.Float64("y", 0, "desired height, 0 for auto, 0 > y > 1 for percentage")
var x = flag.Float64("x", 0, "desired width, 0 for auto, 0 > x > 1 for percentage")

func main() {
	flag.Parse()

	if (*x == 0 && *y == 0) || *x < 0 || *y < 0 {
		usage()
		return
	}

	args := flag.Args()
	if len(args) != 1 {
		usage()
		return
	}

	files, err := filepath.Glob(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Mkdir(PATH, 0666)

	for _, file := range files {
		resizeFile(file, *x, *y)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage example:\n\tresizer -x=0.3 *.jpg\n\n")
	fmt.Fprintf(os.Stderr, "All possible flags:\n")
	flag.CommandLine.PrintDefaults()

}

// mockables
var open = os.Open
var create = os.Create
var decode = jpeg.Decode
var encode = jpeg.Encode
var resizer = resize.Resize

func resizeFile(filename string, xPx, yPx float64) {
	fmt.Println("Processing", filename)

	file, err := open(filename)
	check(err)
	defer file.Close()

	img, err := decode(file)
	check(err)

	if xPx < 1 {
		xPx = float64(img.Bounds().Size().X) * xPx
	}
	if yPx < 1 {
		yPx = float64(img.Bounds().Size().Y) * yPx
	}
	m := resizer(uint(xPx), uint(yPx), img, resize.Bilinear)

	out, err := create(PATH + "/" + filename)
	check(err)
	defer out.Close()

	err = encode(out, m, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

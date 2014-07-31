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

var yPx = flag.Uint("ypx", 0, "height of resized image in pixels, 0 = auto")
var xPx = flag.Uint("xpx", 0, "width of resized image in pixels, 0 = auto")

func main() {
	flag.Parse()

	if *xPx == 0 && *yPx == 0 {
		flag.Usage()
		return
	}

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		return
	}

	files, err := filepath.Glob(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Mkdir(PATH, 0666)

	for _, file := range files {
		resizeFile(file, *xPx, *yPx)
	}
}

// mockables
var open = os.Open
var create = os.Create
var decode = jpeg.Decode
var encode = jpeg.Encode
var resizer = resize.Resize

func resizeFile(filename string, xPx, yPx uint) {
	fmt.Println("Processing", filename)

	file, err := open(filename)
	check(err)
	defer file.Close()

	img, err := decode(file)
	check(err)

	m := resizer(xPx, yPx, img, resize.Lanczos3)

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

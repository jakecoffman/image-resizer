package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

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

func resizeFile(filename string, xPx, yPx uint) {
	if !strings.HasSuffix(filename, ".JPG") {
		return
	}
	fmt.Println("Processing", filename)

	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	img, err := jpeg.Decode(file)
	check(err)

	m := resize.Resize(xPx, yPx, img, resize.Lanczos3)

	out, err := os.Create(PATH + "/" + filename)
	check(err)
	defer out.Close()

	jpeg.Encode(out, m, nil)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

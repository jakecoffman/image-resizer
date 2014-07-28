package main

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

const PATH = "resized"

func main() {

	os.Mkdir(PATH, 0666)

	files, err := ioutil.ReadDir(".")
	check(err)

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".JPG") {
			continue
		}
		fmt.Println("Processing", file.Name())

		file, err := os.Open(file.Name())
		check(err)
		defer file.Close()

		img, err := jpeg.Decode(file)
		check(err)

		newSize := float64(img.Bounds().Size().X) * 0.3

		m := resize.Resize(uint(newSize), 0, img, resize.Lanczos3)

		out, err := os.Create(PATH + "/" + file.Name())
		check(err)
		defer out.Close()

		jpeg.Encode(out, m, nil)
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

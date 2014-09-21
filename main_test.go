package main

import (
	"errors"
	"image"
	"image/jpeg"
	"os"
	"os/exec"
	"testing"
)

func TestParse(t *testing.T) {
	data := []struct {
		x        float64
		y        float64
		args     []string
		expected error
	}{
		{0, 0, []string{""}, errors.New("specify either x or y flag or both")},
		{-1, 0, []string{""}, errors.New("x and y flags must be positive values")},
		{0, -1, []string{""}, errors.New("x and y flags must be positive values")},
		{0, 1, []string{}, errors.New("specify filename or wildcard of filenames to resize")},
		{0, 1, []string{""}, nil},
	}

	for i, d := range data {
		if err := validate(d.x, d.y, d.args); !errEq(d.expected, err) {
			t.Errorf("%v - expected %v got %v", i, d.expected, err.Error())
		}
	}

}

func errEq(e1, e2 error) bool {
	if (e1 == nil && e2 != nil) || (e1 != nil && e2 == nil) {
		return false
	}
	if e1 != nil && e2 != nil && e1.Error() != e2.Error() {
		return false
	}
	return true
}

func TestFunctional(t *testing.T) {
	// Create new image of known size
	original := image.NewRGBA(image.Rect(0, 0, 800, 600))
	file, err := os.Create("test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	err = jpeg.Encode(file, original, nil)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	// Execute resize
	cmd := exec.Command("go", "run", "main.go", "-x=80", "*.jpg")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	t.Log(string(out))

	// Check the new image to see if it is the expected size
	file, err = os.Open("resized/test.jpg")
	defer file.Close()
	if err != nil {
		t.Fatal(err)
	}
	result, format, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}
	if format != "jpeg" {
		t.Fatal("Expected jpeg got ", format)
	}
	if !result.Bounds().Eq(image.Rect(0, 0, 80, 60)) {
		t.Error("Not equal: ", result.Bounds())
	}
}

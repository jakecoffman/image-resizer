package main

import (
	"errors"
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

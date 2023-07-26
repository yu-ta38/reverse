package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var board Board
	board.init()
	expected := 1
	result := check("O")
	if expected != result {
		t.Errorf("Test fail expected:%s, result: %s\n", expected, result)
	}
}

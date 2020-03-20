package main

import "testing"

func TestMinimumSwap(t *testing.T) {
	expected := 2
	actual := minimumSwap([]int{1, 5, 4, 3, 2})
	if expected != actual {
		t.Error("Expected was", expected, "while actual is", actual)
	}
}

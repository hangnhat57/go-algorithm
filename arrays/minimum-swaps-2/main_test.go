package main

import "testing"

func TestMinimumSwap(t *testing.T) {
	expected := 3
	actual := minimumSwap([]int{4, 3, 1, 2})
	if expected != actual {
		t.Error("Expected was", expected, "while actual is", actual)
	}
}

package hourglass

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestHourGlass(t *testing.T) {
	arr := [][]int{
		[]int{1, 1, 1, 0, 0, 0},
		[]int{0, 1, 0, 0, 0, 0},
		[]int{1, 1, 1, 0, 0, 0},
		[]int {0, 0, 2, 4, 4, 0},
		[]int{0, 0, 0, 2, 0, 0},
		[]int{0, 0, 1, 2, 4, 0},
	}
	expect := 19
	actual:= HourGlass(arr)
	assert.Equal(t,expect,actual,"Not correct")
}
func TestHourGlass2(t *testing.T) {
	arr := [][]int{
		[]int{0, -4, -6, 0, -7, -6},
		[]int{-1, -2, -6, -8, -3, -1},
		[]int{-8, -4, -2, -8, -8, -6},
		[]int{-3, -1, -2, -5, -7, -4},
		[]int{-3, -5, -3, -6, -6, -6},
		[]int{-3, -6, 0, -8, -6, -7},
	}
	expect := -19
	actual:= HourGlass(arr)
	assert.Equal(t,expect,actual,"Not correct")
}

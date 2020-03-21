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

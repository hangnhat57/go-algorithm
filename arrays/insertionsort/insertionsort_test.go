package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	expect := []int{1, 2, 3, 4, 5}
	actual := InsertionSort([]int{3, 2, 1, 4, 5})
	assert.Exactly(t, expect, actual, "Does not match")

}
func TestInsertionSortCase0(t *testing.T) {
	expect := []int{0}
	actual := InsertionSort([]int{0})
	assert.Exactly(t, expect, actual, "Does not match")

}

func TestInsertionSortCase9(t *testing.T) {
	expect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	actual := InsertionSort([]int{9, 6, 5, 4, 7, 8, 2, 3, 1})
	assert.Exactly(t, expect, actual, "Does not match")

}

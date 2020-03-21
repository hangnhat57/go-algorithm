package binaryinsertionsort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinaryInsertionSort(t *testing.T) {
	expect := []int{1,2,3,4,5,6,7,8}
	actual := BinaryInsertionSort([]int{8,5,6,4,7,3,1,2})
	assert.Exactly(t,expect,actual,"Not Match")
}
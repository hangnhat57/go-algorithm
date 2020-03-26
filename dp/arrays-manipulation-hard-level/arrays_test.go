package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraysManipulationS(t *testing.T)  {
	expect := 200
	n := 5
	arr := [][]int{
		[]int{1 ,2, 100,},
		[]int{2, 5, 100,},
		[]int{3 ,4 ,100,},
	}
	actual := arrayManipulationStupid(n,arr)
	assert.Equal(t,expect,actual)

}

func TestArraysManipulationDP(t *testing.T)  {
	expect := 200
	n := 5
	arr := [][]int{
		[]int{1 ,2, 100,},
		[]int{2, 5, 100,},
		[]int{3 ,4 ,100,},
	}
	actual := arrayManipulation(n,arr)
	assert.Equal(t,expect,actual)

}
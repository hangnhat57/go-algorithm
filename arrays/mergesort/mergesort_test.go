package mergesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T)  {
	actual:=Sort([]int{5,6,4,2,3,1})
	expect:= []int{1,2,3,4,5,6}
	assert.Exactly(t,expect,actual,"Not Match")
}
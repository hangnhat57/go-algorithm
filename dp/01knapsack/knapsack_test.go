package knapsack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnapsackNaive(t *testing.T)  {
	weight := []int{1,3,2,4,5,7}
	value :=  []int{4,3,5,6,1,2}
	cap := 6
	actual := KnapSackNaive(weight,value,cap)
	assert.Equal(t,12,actual)
}

func TestBottomUpKnapsack(t *testing.T)  {
	weight := []int{1,3,2,4,5,7}
	value :=  []int{4,3,5,6,1,2}
	cap := 6
	actual := KnapsackBottomUp(weight,value,cap)
	assert.Equal(t,12,actual)
}


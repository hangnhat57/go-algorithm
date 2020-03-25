package heapsort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
	"time"
)


func TestHeapSort(t *testing.T) {
	for i:= 10;i<30;i++{
		arr := generateRandomSlice(i)
		expect := make([]int,len(arr))
		copy(expect,arr)
		sort.Ints(expect)

		actual :=HeapSort(arr)
		assert.Exactly(t,expect,actual,"Not Match")
	}

}

func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
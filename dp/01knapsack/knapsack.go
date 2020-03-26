package knapsack

import (
	"github.com/hangnhat57/go-algorithm/arrays/mergesort"
)

//[1 3 2 4 5 7]
//[4 3 5 6 1 2]



func KnapSackNaive(weight, value []int, capacity int) int {
	var arr []int
	for i := 0; i < len(weight); i++ {
		arr = append(arr, naiveKap(weight, value, i, capacity))
	}
	return mergesort.Sort(arr)[len(arr)-1]
}

func naiveKap(w, v []int, idx, capacity int) int {
	if capacity == 0 || idx > len(w)-1 {
		return 0
	}
	if w[idx] > capacity {
		return naiveKap(w, v, idx+1, capacity)
	}
	if w[idx] < capacity {
		atmp1 := naiveKap(w, v, idx+1, capacity)
		atmp2 := v[idx] + naiveKap(w, v, idx+1, capacity-w[idx])
		if atmp1 > atmp2 {
			return atmp1
		} else {
			return atmp2
		}
	} else {
		return v[idx]
	}
}

func max(a,b int) int {
	if a > b {return a}else {return b}
}

func KnapsackBottomUp(weight, value []int, capacity int) int {

	var K = new2DArray(len(weight)+1,capacity+1)

	for i:= 0 ; i < len(value)+1;i++{
		for c := 0;c < capacity+1;c ++{
			if i == 0 || c == 0{
				K[i][c] = 0
			} else if c < weight[i-1] {
				K[i][c] = K[i-1][c]
			}else {
				K[i][c] = max(
					//value[1] + K[1][0]
					value[i-1] + K[i-1][c-weight[i-1]],
					K[i-1][c])
			}
			if i != 0 && c != 0 {
			}
		}
	}

	return K[len(weight)][capacity]
}

func new2DArray(x, y int) [][]int {
	a := make([][]int, y)
	for i := range a {
		a[i] = make([]int, x)
	}
	return a
}

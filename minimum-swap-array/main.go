package main

import "fmt"

func main() {
	fmt.Println(minimumSwap([]int{1, 2, 3, 4, 6, 5}))
}

/*
Given an array of n distinct elements, find the minimum number of swaps required to sort the array.
*/

/*
Solution:
1. Store array with matched indicator in a hash ( map )
2. If value - key match => This element already in correct position. Do not swap. Mask as visited
3. If value - key does not match => swap position between
*/

func minimumSwap(arr []int) int {
	var counter int                 //=> For counting number of swap
	var hashArr = make(map[int]int) //=> store all given arr elements with their position/indicator

	//Pushed all elements && indicators to the hash
	for i := 0; i < len(arr); i++ {
		hashArr[i+1] = arr[i]
	}
	// Make another hash to mask if element is visited/checked. By default, Golang init all value to false
	var visited = make(map[int]bool)

	//Loop all element in element hash to check, start
	for k, v := range hashArr {
		// If element has not been checked/visited before
		if !visited[k] {
			// Mask as checked
			visited[k] = true
			// If the value && indicator do not match, they need to be swap
			if k != v {
				// Because we already swap between 2 number, check the next position if it's correct position
				nextNumber := v
				// If next number haven't checked
				// This is `while` in golang way :D
				for !visited[nextNumber] {
					//Masked as checked this position
					visited[nextNumber] = true
					// change nextNumber to next position until it checked.
					nextNumber = hashArr[nextNumber]
					counter++
				}
			}
		}
	}
	return counter
}

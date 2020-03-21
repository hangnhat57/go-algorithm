package insertionsort

//InsertionSort implement of insertion sort algorithm
// Complexity : O(n^2)
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		idx := i
		for idx >= 1 && arr[idx-1] > arr[idx] {
			tmp := arr[idx]
			arr[idx] = arr[idx-1]
			arr[idx-1] = tmp
			idx--
		}
	}
	return arr
}

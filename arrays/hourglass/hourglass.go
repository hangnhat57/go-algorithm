package hourglass

import "github.com/hangnhat57/go-algorithm/arrays/insertionsort"

func HourGlass(arr [][]int) int {
	rows := len(arr)
	columns := len(arr[0])
	var values []int
	for i:= 0;i<rows-2;i++{
		for j:= 0;j<columns-2;j++{
			sum := arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] +arr[i+2][j]+arr[i+2][j+1]+arr[i+2][j+2]
			values = append(values,sum)
		}
	}
	values = insertionsort.InsertionSort(values)
	return values[len(values)-1]
}
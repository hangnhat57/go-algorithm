package mergesort

import "fmt"

func Sort(arr []int)  []int{
	right := len(arr)
	fmt.Println(arr)
	return subSort(arr,0,right)
}

func subSort(arr []int,left,right int)  []int{
	if len(arr) == 1 {
		return arr
	}
	mid := (left + right )/2
	arr1 := make([]int,mid - left)
	arr2 := make([]int,right-mid)
	copy(arr1,arr[left:mid])
	copy(arr2,arr[mid:right])
	//fmt.Println("Mid",mid,"Left",left,"Right",right,"Left Arr",arr1,"Right Arr",arr2)
	//subSort(arr1,left,mid)
	//subSort(arr2,0,right-mid)
	return mergeSub(subSort(arr1,left,mid),subSort(arr2,0,right-mid))
}

func mergeSub(left,right []int)  []int{
	size,i,j := len(left)+len(right),0,0
	arr := make([]int,size)
		for idx := 0;idx <size;idx++{
			if i > len(left)-1 && j <= len(right)-1{
				arr[idx] = right[j]
				j++
			}else if j > len(right)-1 && i <= len(left)-1{
				arr[idx] = left[i]
				i++
			}else if left[i] > right[j]{
				arr[idx] = right[j]
				j++
			} else {
					arr[idx] = left[i]
					i++
			}
		}
	return arr
}
package heapsort


// [1,2,3,4,5,6]
// parent == 2i
// left child == 2i+1
// right child == 2i +2
//last parent node == n/2 -1
//Algorithm explain at https://www.youtube.com/watch?v=MtQL_ll5KhQ

func HeapSort(arr []int)  []int{
	buildHeap(arr)
	for i:= len(arr)-1;i>=0;i-- {
		swap(arr,0,i)
		heapify(arr[:i],0,i-1)
	}
	return arr
}

func buildHeap(arr []int)  {
	for i:= len(arr)/2 -1 ; i >= 0 ; i--{
		heapify(arr,i,len(arr)-1)
	}
}
func left(idx int) int {
	return 2*idx +1
}

func  right(idx int) int {
	return 2*idx +2
}

func swap(arr []int,des,src int){
		tmp := arr[des]
		arr[des] = arr[src]
		arr[src] = tmp
}
func heapify(arr []int,idx int,capacity int )  {
	var min int
	if left(idx) <= capacity && arr[left(idx)] > arr[idx] {
		min = left(idx)
	}else {
		min = idx
	}
	if right(idx) <= capacity && arr[right(idx)] > arr[min]{
		min = right(idx)
	}
	if min != idx {
		swap(arr,min,idx)
		heapify(arr,min,capacity)
	}
}




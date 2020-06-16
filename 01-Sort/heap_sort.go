package Sort

func heap_sort(arr []int) {
	buildMaxHeap(arr)

	len := len(arr)
	for i := len - 1; i > 0; i-- {
		swap(arr, 0, i)
		len--
		heapify(arr, 0)
	}
}

func buildMaxHeap(arr []int) { // 建立大顶堆
	len := len(arr)

	for i := int(len / 2); i >= 0; i-- {
		heapify(arr, i)
	}
}

func heapify(arr []int, i int) { // 堆调整

	left := 2*i + 1
	right := 2*i + 2
	largest := i
	len := len(arr)

	if left < len && arr[left] > arr[largest] {
		largest = left
	}

	if right < len && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest)
	}
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

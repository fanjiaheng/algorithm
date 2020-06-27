package Sort

func merge_sort(arr []int, low, high int) {
	if low < high {
		mid := (low + high) / 2
		merge_sort(arr, low, mid)
		merge_sort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}

func merge(arr []int, low, mid, high int) {
	leftLen := mid - low + 1
	rightLen := high - mid

	arrLeft := make([]int, leftLen)
	for i := 0; i < leftLen; i++ {
		arrLeft[i] = arr[low+i]
	}

	arrRight := make([]int, rightLen)
	for j := 0; j < rightLen; j++ {
		arrRight[j] = arr[mid+j+1]
	}

	i, j, k := 0, 0, low
	for ; k <= high && i < leftLen && j < rightLen; k++ {
		if arrLeft[i] <= arrRight[j] {
			arr[k] = arrLeft[i]
			i++
		} else {
			arr[k] = arrRight[j]
			j++
		}
	}

	for ; i < leftLen && k <= high; k++ {
		arr[k] = arrLeft[i]
		i++
	}

	for ; j < rightLen && k <= high; k++ {
		arr[k] = arrRight[j]
		j++
	}

}

package Sort

func select_sort(arr []int) {
	len := len(arr)
	minIndex := 0
	tmp := 0

	for i := 0; i < len-1; i++ {
		minIndex = i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		tmp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = tmp
	}
}

// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
// 不稳定
// 例子: [5, 8, 5, 2, 9]

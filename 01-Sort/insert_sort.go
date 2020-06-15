package Sort

func insert_sort(arr []int) {

	len := len(arr)

	preIndex := 0
	current := 0
	for i := 1; i < len; i++ {

		preIndex = i - 1
		current = arr[i]

		for {
			if preIndex >= 0 && arr[preIndex] > current {
				arr[preIndex+1] = arr[preIndex]
				preIndex--
			}
		}
		arr[preIndex+1] = current
	}
}

// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
// 稳定

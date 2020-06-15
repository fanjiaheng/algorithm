package Sort

func shell_sort(arr []int) {
	len := len(arr)

	for gap := int(len / 2); gap > 0; gap = int(gap / 2) {
		for i := gap; i < len; i++ {
			j := i
			current := arr[i]
			for {
				if j-gap >= 0 && current < arr[j-gap] {
					arr[j] = arr[j-gap]
					j = j - gap
				} else {
					break
				}
			}
			arr[j] = current
		}
	}
}

// 时间复杂度 O(n^(1.3—2))
// 空间复杂度 O(1)
// 不稳定

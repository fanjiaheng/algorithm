package Sort

// 冒泡排序
func bubble_sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
// 稳定

package Search

func sequence_search(arr []int, wantVal int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == wantVal {
			return i
		}
	}

	return -1
}

// 时间复杂度 O(n)

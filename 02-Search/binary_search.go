package Search

//二分查找，递归版本
func binary_search(arr []int, val, low, high int) int {
	mid := int(low + (high-low)/2)
	if arr[mid] == val {
		return mid
	}

	if arr[mid] > val {
		return binary_search(arr, val, low, mid-1)
	}

	if arr[mid] < val {
		return binary_search(arr, val, mid+1, high)
	}

	return -1
}

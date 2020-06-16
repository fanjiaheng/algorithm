package Search

import "fmt"

func inseart_search(arr []int, val, low, high int) int {

	mid := low + int((val-arr[low])/(arr[high]-arr[low])*(high-low))

	fmt.Println("***********************")
	fmt.Println(mid)

	if arr[mid] == val {
		return mid
	}

	if arr[mid] > val {
		return inseart_search(arr, val, low, mid-1)
	}

	if arr[mid] < val {
		return inseart_search(arr, val, mid+1, high)
	}

	return -1
}

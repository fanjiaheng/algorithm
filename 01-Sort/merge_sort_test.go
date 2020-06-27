package Sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(T *testing.T) {
	fmt.Println("*******二路归并排序测试*******")

	arr := []int{10, 14, 3, 4, 1}
	merge_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

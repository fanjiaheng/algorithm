package Sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(T *testing.T) {
	fmt.Println("*******快速排序测试*******")

	arr := []int{10, 14, 3, 4, 1, 100, 99}
	quick_sort(arr, 0, 4)
	fmt.Println(arr)
}

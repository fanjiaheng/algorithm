package Sort

import (
	"fmt"
	"testing"
)

func TestHeapRecursionSort(T *testing.T) {
	fmt.Println("*******递归堆排序测试*******")

	arr := []int{10, 14, 3, 4, 1, 10}
	heap_recursion_sort(arr)
	fmt.Println(arr)
}

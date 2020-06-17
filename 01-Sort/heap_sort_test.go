package Sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(T *testing.T) {
	fmt.Println("*******堆排序测试*******")

	arr := []int{10, 14, 3, 4, 1, 11, 23, 44, 1, 21, 33}
	heap_sort(arr)
	fmt.Println(arr)
}

// 输出
// === RUN   TestHeapSort
// *******堆排序测试*******
// [1 1 3 4 10 11 14 21 23 33 44]
// --- PASS: TestHeapSort (0.00s)
// PASS
// ok      command-line-arguments  0.003s

package Sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(T *testing.T) {
	fmt.Println("*******选择排序测试*******")

	arr := []int{10, 14, 3, 4, 1}
	select_sort(arr)
	fmt.Println(arr)
}

// 输出
// *******选择排序测试*******
// [1 3 4 10 14]
// PASS
// ok      algorithm/01-Sort       0.002s

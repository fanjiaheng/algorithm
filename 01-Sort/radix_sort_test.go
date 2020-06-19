package Sort

import (
	"fmt"
	"testing"
)

func TestRadixSort(T *testing.T) {
	fmt.Println("*******基数排序测试*******")

	arr := []int{10, 14, 3, 4, 1}

	fmt.Println(arr)
}

// 输出
// *******选择排序测试*******
// [1 3 4 10 14]
// PASS
// ok      algorithm/01-Sort       0.002s

package Sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	fmt.Println("*******冒泡排序测试*******")

	arr := []int{10, 14, 3, 4, 1}
	bubble_sort(arr)
	fmt.Println(arr)
}

// 输出
// *******冒泡排序测试*******
// [1 3 4 10 14]
// PASS
// ok      algorithm/01-Sort       0.003s

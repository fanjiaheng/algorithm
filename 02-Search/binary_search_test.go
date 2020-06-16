package Search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(T *testing.T) {
	fmt.Println("*******二分查找/折半查找测试*******")

	arr := []int{10, 14, 3, 4, 1}
	fmt.Println(binary_search(arr, 4, 0, 4))
}

// 输出
// *******顺序查找测试*******
// 3
// PASS
// ok      algorithm/01-Sort       0.003s

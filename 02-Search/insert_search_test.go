package Search

import (
	"fmt"
	"testing"
)

func TestInsertSearch(T *testing.T) {
	fmt.Println("*******插值查找测试*******")

	arr := []int{10, 14, 3, 4, 1}
	fmt.Println(inseart_search(arr, 4, 0, 4))
}

// 输出
// *******顺序查找测试*******
// 3
// PASS
// ok      algorithm/01-Sort       0.003s

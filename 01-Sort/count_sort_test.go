package Sort

import (
	"fmt"
	"testing"
)

func TestCountSort(T *testing.T) {
	fmt.Println("*******计数排序测试*******")

	arr := []int{10, 14, 3, 4, 1, 11, 23, 44, 1, 21, 33}
	count_sort(arr)
	fmt.Println(arr)
}

package Search

import (
	"fmt"
	"testing"
)

func TestHashSearch(T *testing.T) {
	fmt.Println("*******哈希查找测试*******")

	h_tbl := &hash_table{}
	hash_init(h_tbl) // 初始化

	arr := []int{10, 14, 3, 4, 1}
	for i := 0; i < len(arr); i++ {
		insert(h_tbl, arr[i])
	}
	show(h_tbl)

	index := hash_search(h_tbl, 4)
	if index == -1 {
		fmt.Println("哈希查找失败!")
		return
	}

	fmt.Println(index)
}

package Search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(T *testing.T) {
	fmt.Println("*******二叉排序树查找测试*******")

	arr := []int{10, 14, 3, 4, 1}

	root := &BTnode{}
	CreateBST(root, arr)

	BST_preorder_traverse(root)

	fmt.Println(BSTSearch(root, 14))
}

// 输出
// *******顺序查找测试*******
// 3
// PASS
// ok      algorithm/01-Sort       0.003s

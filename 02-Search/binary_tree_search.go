package Search

import (
	"fmt"
)

type BTnode struct {
	Key    int // 关键字
	LChild *BTnode
	RChild *BTnode
}

// 二叉排序树查找算法
func BSTSearch(bt *BTnode, key int) *BTnode {

	if bt == nil {
		return nil
	}

	if bt.Key == key {
		return bt
	} else if bt.Key < key {
		return BSTSearch(bt.RChild, key)
	} else {
		return BSTSearch(bt.LChild, key)
	}

	return nil
}

// 构建二叉排序树
func BSTInsert(bt *BTnode, key int) *BTnode {

	if bt == nil {
		fmt.Println("1111111111111")
		bt = &BTnode{
			Key:    key,
			LChild: nil,
			RChild: nil,
		}

		return bt
	} else {
		fmt.Println("22222222")
		if bt.Key == key {
			return bt
		} else if bt.Key > key {
			return BSTInsert(bt.LChild, key)
		} else {
			return BSTInsert(bt.RChild, key)
		}
	}

	return bt
}

func CreateBST(bt *BTnode, arr []int) {

	fmt.Println("************************")
	fmt.Println(arr)

	bt = nil
	for i := 0; i < len(arr); i++ {
		BSTInsert(bt, arr[i])
	}
}

// 遍历二叉排序树
func BST_preorder_traverse(bt *BTnode) {
	if bt == nil {
		return
	}

	// visit
	fmt.Println(bt.Key)

	if bt.LChild != nil {
		BST_preorder_traverse(bt.LChild)
	}
	if bt.RChild != nil {
		BST_preorder_traverse(bt.RChild)
	}
}

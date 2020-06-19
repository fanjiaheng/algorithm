package Tree

import "testing"

func TestBTreeMiddleorderTraverse(T *testing.T) {
	arr := []byte{'A', 'B', 'D', '#', '#', 'E', '#', 'H', '#', '#', 'C', 'F', '#', '#', 'G'}

	index := 0
	root := create_bTree(arr, &index)
	bTree_middleorder_traverse(root)
}

package Tree

import "testing"

func TestBTreePreorderTraverse(T *testing.T) {
	arr := []byte{'A', 'B', 'D', '#', '#', 'E', '#', 'H', '#', '#', 'C', 'F', '#', '#', 'G'}

	index := 0
	root := create_bTree(arr, &index)
	bTree_preorder_traverse(root)
}

package Tree

import "testing"

func TestThreadBTree(T *testing.T) {
	arr := []byte{'A', 'B', 'D', '#', '#', 'E', '#', 'H', '#', '#', 'C', 'F', '#', '#', 'G'}

	index := 0
	root := create_threadBTree(arr, &index)
	threadBTree_preorder_traverse(root)
	inThreading(root)
	// threadBTree_preorder_traverse(root)
	// inOrderThreadBTree(root)
}

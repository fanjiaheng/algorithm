package Tree

import "fmt"

func bTree_middleorder_traverse(root *BTree) {

	if root == nil {
		return
	}

	if root.Left != nil {
		bTree_middleorder_traverse(root.Left)
	}

	// visit
	fmt.Println(root.Value)

	if root.Right != nil {
		bTree_middleorder_traverse(root.Right)
	}
}

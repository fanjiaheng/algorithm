package Tree

import "fmt"

func bTree_postorder_traverse(root *BTree) {

	if root == nil {
		return
	}

	if root.Left != nil {
		bTree_postorder_traverse(root.Left)
	}
	if root.Right != nil {
		bTree_postorder_traverse(root.Right)
	}

	// visit
	fmt.Println(root.Value)
}

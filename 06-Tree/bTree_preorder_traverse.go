package Tree

import (
	"fmt"
)

func bTree_preorder_traverse(root *BTree) {

	if root == nil {
		return
	}

	// visit
	fmt.Println(root.Value)

	if root.Left != nil {
		bTree_preorder_traverse(root.Left)
	}
	if root.Right != nil {
		bTree_preorder_traverse(root.Right)
	}
}

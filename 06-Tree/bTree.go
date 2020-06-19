package Tree

// 二叉树的节点
type BTree struct {
	Value byte
	Left  *BTree
	Right *BTree
}

// 创建二叉树（先序创建二叉树）
func create_bTree(arr []byte, index *int) *BTree {

	if len(arr) == *index {
		return nil
	}

	if arr[*index] == '#' {
		(*index)++
		return nil
	}

	root := &BTree{}
	root.Value = arr[*index]
	(*index)++
	root.Left = create_bTree(arr, index)
	root.Right = create_bTree(arr, index)

	return root
}

package Tree

// 二叉树的节点
type BTree struct {
	Value int
	Left  *BTree
	Right *BTree
}

// 创建二叉树
func create_bTree(arr []int, index int) *BTree {

	if len(arr) == 0 {
		// fmt.Println("输入的序列为空, 不能创建二叉树!")
		return nil
	}

	// BTree *node = new BTree()
	// node.value = arr[index]
	// index++
	// if len(arr) > index {
	// 	node.Left = create_bTree(arr, index)
	// }

	// if len(arr) > index {
	// 	node.Left = create_bTree(arr, index)
	// }

}

func insert_node() *BTree {

}

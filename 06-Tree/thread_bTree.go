package Tree

import "fmt"

// 线索二叉树是为了加快查找节点前驱和后继的速度
type ThreadBTree struct {
	Value byte
	Left  *ThreadBTree
	Right *ThreadBTree
	Ltag  byte
	Rtag  byte
}

// 记ptr指向二叉链表中的一个结点，以下是建立线索的规则：
// 如果ptr->lchild为空，则存放指向中序遍历序列中该结点的前驱结点。这个结点称为ptr的中序前驱；
// 如果ptr->rchild为空，则存放指向中序遍历序列中该结点的后继结点。这个结点称为ptr的中序后继；

// ltag为0时指向该结点的左孩子，为1时指向该结点的前驱；
// rtag为0时指向该结点的右孩子，为1时指向该结点的后继；

// 线索化的实质就是将二叉链表中的空指针改为指向前驱或后继的线索。
// 由于前驱和后继信息只有在遍历该二叉树时才能得到，所以，线索化的
// 过程就是在遍历的过程中修改空指针的过程。

//全局变量，始终指向刚刚访问过的结点
//注意：此处一定要初始化!!!
var pre *ThreadBTree = &ThreadBTree{}

//中序遍历进行中序线索化
func inThreading(root *ThreadBTree) {

	if root != nil {
		inThreading(root.Left) //递归左子树线索化

		if root.Left == nil { //没有左孩子
			root.Ltag = 1   //前驱线索
			root.Left = pre //左孩子指针指向前驱
		}

		if root.Right == nil { //没有右孩子
			pre.Rtag = 1     //后继线索
			pre.Right = root //前驱右孩子指针指向后继(当前结点root)
		}

		pre = root
		inThreading(root.Right) //递归右子树线索化
	}
}

// 创建二叉树（先序创建线索结构的二叉树）
func create_threadBTree(arr []byte, index *int) *ThreadBTree {

	if len(arr) == *index {
		return nil
	}

	if arr[*index] == '#' {
		(*index)++
		return nil
	}

	root := &ThreadBTree{}
	root.Value = arr[*index]
	(*index)++
	root.Left = create_threadBTree(arr, index)
	root.Right = create_threadBTree(arr, index)

	return root
}

// 先序遍历线索二叉树
func threadBTree_preorder_traverse(root *ThreadBTree) {

	if root == nil {
		return
	}

	// visit
	fmt.Println(root.Value)

	if root.Left != nil {
		threadBTree_preorder_traverse(root.Left)
	}

	if root.Right != nil {
		threadBTree_preorder_traverse(root.Right)
	}
}

// 中序遍历线索二叉树
func inOrderThreadBTree(root *ThreadBTree) {

	if root != nil {
		var p *ThreadBTree = &ThreadBTree{}

		p = root.Left
		for {
			if p != root {
				for {
					if p.Ltag == 0 {
						p = p.Left
					} else {
						break
					}
				}

				fmt.Println(p.Value)

				for {
					if p.Rtag == 1 && p.Right != root {
						p = p.Right
						fmt.Println(p.Value)
					} else {
						break
					}
				}

				p = p.Right
			} else {
				break
			}
		}
	}
}

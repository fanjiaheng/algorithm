package Sort

// 递归的堆排序，堆其实就是一棵完全二叉树，且具有自己的特性。
// 如果给定一个序列的位置（数组下标），则可以通过下列公式查找
// 其父节点和左右孩子节点。
// parent = (i-1)/2 --- 向下取整
// lchild = 2*i + 1
// rchild = 2*i + 2
func heap_recursion_sort(tree []int) {

	n := len(tree)

	build_heap(tree, n)
	for i := n - 1; i >= 0; i-- {
		swap(tree, i, 0)
		heapify(tree, i, 0)
	}
}

// 建堆
func build_heap(tree []int, n int) {
	last_node := n - 1
	parent := int((last_node - 1) / 2)

	for i := parent; i >= 0; i-- {
		heapify(tree, n, i)
	}
}

// n: 切片中未排序关键字的个数
func heapify(tree []int, n, i int) {

	if i >= n {
		return
	}

	c1 := 2*i + 1
	c2 := 2*i + 2
	max := i
	if c1 < n && tree[c1] > tree[max] {
		max = c1
	}

	if c2 < n && tree[c2] > tree[max] {
		max = c2
	}

	if max != i {
		swap(tree, max, i)
		heapify(tree, n, max)
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

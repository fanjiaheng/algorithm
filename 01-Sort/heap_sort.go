package Sort

func heap_sort(arr []int) {

	for i := int(len(arr)/2 - 1); i >= 0; i-- {
		adjust_heap(arr, i, len(arr))
	}

	for j := len(arr) - 1; j > 0; j-- {
		swap(arr, 0, j)        //将堆顶元素与末尾元素进行交换
		adjust_heap(arr, 0, j) //重新对堆进行调整
	}
}

func adjust_heap(arr []int, i, length int) {
	tmp := arr[i]

	//从i结点的左子结点开始，也就是2i+1处开始
	for k := i*2 + 1; k < length; k = k*2 + 1 {

		//如果左子结点小于右子结点，k指向右子结点
		if k+1 < length && arr[k] < arr[k+1] {
			k++
		}

		//如果子节点大于父节点，将子节点值赋给父节点（不用进行交换）
		if arr[k] > tmp {
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	arr[i] = tmp // 将tmp值放到最终的位置
}

func swap(arr []int, a, b int) {
	tmp := arr[a]
	arr[a] = arr[b]
	arr[b] = tmp
}

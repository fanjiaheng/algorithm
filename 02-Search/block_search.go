package Search

// 单个分块的结构体
type bIndex struct {
	start  int // 块开始索引
	length int // 块长度
	max    int // 块最大值
}

/*
arr 满足分块查找的数组
bs 分块信息数组
val 查找的值
*/
func BlockSearch(arr []int, bs []bIndex, val int) int {
	// 先找到所属块
	var bi bIndex
	for _, b := range bs {
		if b.max > val {
			bi = b
			break
		}
	}

	// 再在该块中查找该元素
	for k, v := range arr[bi.start : bi.start+bi.length] {
		if v == val {
			return bi.start + k
		}

	}

	return -1
}

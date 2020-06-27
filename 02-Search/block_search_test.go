package Search

import (
	"fmt"
	"testing"
)

func TestBlockSearch(T *testing.T) {
	fmt.Println("*******分块查找测试*******")

	// 先构建符合条件的数组和分块
	blockSlice := []int{3, 4, 6, 2, 5, 7, 14, 12, 16, 13, 19, 17, 25, 21, 36, 23, 22, 29}
	indices := make([]bIndex, 3)
	indices[1] = bIndex{start: 0, length: 6, max: 7}
	indices[1] = bIndex{start: 6, length: 6, max: 19}
	indices[1] = bIndex{start: 12, length: 6, max: 36}

	i := BlockSearch(blockSlice, indices, 21)
	j := BlockSearch(blockSlice, indices, 1)
	T.Logf("BlockSearch:%v, existVal:21,existIndex:%d, noExistVal:1,noExistIndex:%d",
		blockSlice, i, j)
}

// 输出
// *******分块查找测试*******
// === RUN   TestBlockSearch
// --- PASS: TestBlockSearch (0.00s)
//     block_search_test.go:15: BlockSearch:[3 4 6 2 5 7 14 12 16 13 19 17 25 21 36 23 22 29], existVal:21,existIndex:13, noExistVal:1,noExistIndex:-1
// PASS
// ok      command-line-arguments  0.182s

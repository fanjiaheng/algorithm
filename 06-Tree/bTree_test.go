package Tree

import "testing"

func TestCreateBTree(T *testing.T) {

	arr := []byte{'A', 'B', 'D', '#', '#', 'E', '#', 'H', '#', '#', 'C', 'F', '#', '#', 'G'}

	index := 0
	_ = create_bTree(arr, &index)

}

package Search

//顺序查找
// int SequenceSearch(int a[], int value, int n)
// {
//     int i;
//     for(i=0; i<n; i++)
//         if(a[i]==value)
//             return i;
//     return -1;
// }

func sequence_search(arr []int, wantVal int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == wantVal {
			return i
		}
	}

	return -1
}

// 时间复杂度 O(n)

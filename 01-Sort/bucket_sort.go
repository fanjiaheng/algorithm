package Sort

import "fmt"

func bucket_sort(arr []int) {
	len := len(arr)

	if len == 0 {
		fmt.Println("输入的切片为空!")
		return
	}

	minValue := arr[0]
	maxValue := arr[0]
	for i := 1; i < len; i++ {
		if arr[i] < minValue { // 输入数据的最小值
			minValue = arr[i]
		} else if arr[i] > maxValue { // 输入数据的最大值
			maxValue = arr[i]
		} else {

		}
	}

	// 桶的初始化
	// bucketSize := 5 // 设置桶的默认数量为5
	// bucketCount := int((maxValue-minValue)/bucketSize) + 1
	// buckets := [bucketCount]int{}
	// for i := 0; i < bucketCount; i++ {
	//     buckets[i] = []
	// }

	// 利用映射函数将数据分配到各个桶中
	for i := 0; i < len; i++ {
		buckets[int((maxValue-minValue)/bucketSize)] = append(buckets[int((maxValue-minValue)/bucketSize)], arr[i])
	}

}

// function bucketSort(arr, bucketSize) {
//     if (arr.length === 0) {
//       return arr;
//     }

//     var i;
//     var minValue = arr[0];
//     var maxValue = arr[0];
//     for (i = 1; i < arr.length; i++) {
//       if (arr[i] < minValue) {
//           minValue = arr[i];                // 输入数据的最小值
//       } else if (arr[i] > maxValue) {
//           maxValue = arr[i];                // 输入数据的最大值
//       }
//     }

//     // 桶的初始化
//     var DEFAULT_BUCKET_SIZE = 5;            // 设置桶的默认数量为5
//     bucketSize = bucketSize || DEFAULT_BUCKET_SIZE;
//     var bucketCount = Math.floor((maxValue - minValue) / bucketSize) + 1;
//     var buckets = new Array(bucketCount);
//     for (i = 0; i < buckets.length; i++) {
//         buckets[i] = [];
//     }

//     // 利用映射函数将数据分配到各个桶中
//     for (i = 0; i < arr.length; i++) {
//         buckets[Math.floor((arr[i] - minValue) / bucketSize)].push(arr[i]);
//     }

//     arr.length = 0;
//     for (i = 0; i < buckets.length; i++) {
//         insertionSort(buckets[i]);                      // 对每个桶进行排序，这里使用了插入排序
//         for (var j = 0; j < buckets[i].length; j++) {
//             arr.push(buckets[i][j]);
//         }
//     }

//     return arr;
// }
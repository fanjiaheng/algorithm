package Sort

import "fmt"

var counter []int

func radix_sort(arr []int, maxDigit int) {
	fmt.Println("基数排序")

	// mod := 10
	// dev := 1

	// for i := 0; i < maxDigit; i++, dev *= 10, mod *= 10 {
	// 	for j = 0; j < arr.length; j++ {
	// 		bucket := int((arr[j] % mod) / dev)
	// 		if counter[bucket] == 0 {

	// 		}
	// 		counter == append(counter, arr[j])
	// 	}

	// 	pos := 0
	// 	for j:= 0; j < len(counter); j++ {
	// 		value := 0
	// 		if counter[j] != 0 {
	// 			// for {
	// 			// 	// if
	// 			// }
	// 		}
	// 	}
	// }

}

// var counter = [];
// function radixSort(arr, maxDigit) {
//     var mod = 10;
//     var dev = 1;
//     for (var i = 0; i < maxDigit; i++, dev *= 10, mod *= 10) {
//         for (var j = 0; j < arr.length; j++) {
//             var bucket = parseInt((arr[j] % mod) / dev);
//             if(counter[bucket]==null) {
//                 counter[bucket] = [];
//             }
//             counter[bucket].push(arr[j]);
//         }
//         var pos = 0;
//         for (var j = 0; j < counter.length; j++) {
//             var value = null;
//             if(counter[j]!=null) {
//                 while ((value = counter[j].shift()) != null) {
//                       arr[pos++] = value;
//                 }
//           }
//         }
//     }
//     return arr;
// }

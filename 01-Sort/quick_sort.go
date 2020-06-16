package Sort

import "fmt"

func quick_sort(arr []int, low, high int) {

	i := low
	j := high
	tmp := 0

	if low < high {

		for {
			if i < j {
				tmp = arr[i] // 枢柚
				for {
					if j > i && arr[j] >= tmp { //从右向左扫描，找到一个小于tmp的关键字
						j--
					} else {
						break
					}
				}

				if i < j {
					arr[i] = arr[j]
					i++
				}

				fmt.Println(arr[i])

				for {
					if i < j && arr[i] < tmp { //从左向右扫描，找到一个大于tmp的关键字
						i++
					} else {
						break
					}
				}

				if i < j {
					arr[j] = arr[i]
					j--
				}
				fmt.Println(arr[j])
			} else {
				break
			}
		}
		// {10, 14, 3, 4, 1}
		arr[i] = tmp // 将tmp放到最终的位置

		fmt.Println("1--------")
		fmt.Println(i)
		fmt.Println(arr[i])

		quick_sort(arr, low, i-1)  // 递归的对tmp左边的关键字进行排序
		quick_sort(arr, i+1, high) // 递归的对tmp右边的关键字进行排序
	}
}

// 时间复杂度 O(nlog2n)
// 空间复杂度 O(nlog2n)
// 不稳定

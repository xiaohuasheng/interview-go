package sort

func quickSort(arr []int, begin, end int) {
	if begin >= end {
		return
	}
	l, r := begin, end
	temp := arr[begin]
	for l < r {
		// 先移动右边的
		// 这里也要判断 l < r
		for l < r && arr[r] >= temp {
			r--
		}
		for l < r && arr[l] <= temp {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	// 最后交换基准数与指针相遇位置的数
	arr[begin], arr[l] = arr[l], arr[begin]
	quickSort(arr, begin, l-1)
	quickSort(arr, l+1, end)
}

//func quickSort(arr []int, begin, end int) {
//	// 递归出口
//	if begin >= end {
//		return
//	}
//	// l, r 左右指针
//	l, i, r := begin, begin+1, end
//	// 基准数
//	v := arr[begin]
//	// 左右指针相遇的时候退出扫描循环
//	for i <= r {
//		if arr[i] > v {
//			swap(arr, i, r)
//			r--
//		} else if arr[i] < v {
//			swap(arr, l, i)
//			l++
//			i++
//		} else {
//			i++
//		}
//	}
//	quickSort(arr, begin, l-1)
//	quickSort(arr, r+1, end)
//}
//
//func swap(arr []int, index1, index2 int) {
//	tmp := arr[index1]
//	arr[index1] = arr[index2]
//	arr[index2] = tmp
//}

package sort

import "fmt"

func sink(arr []int, i, length int) {
	for left := 2*i + 1; left < length; left = 2*i + 1 {
		if left+1 < length && arr[left] < arr[left+1] {
			left++
		}
		if arr[i] >= arr[left] {
			break
		}
		arr[left], arr[i] = arr[i], arr[left]
		i = left
	}
}

func heapsort() {
	arr := []int{7, 9, 4, 10, 3, 5}
	// 构造堆
	length := len(arr)
	for i := length/2 - 1; i >= 0; i-- {
		sink(arr, i, length)
	}
	// 排序
	for i := length - 1; i > 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		sink(arr, 0, i)
	}
	fmt.Println(arr)
}

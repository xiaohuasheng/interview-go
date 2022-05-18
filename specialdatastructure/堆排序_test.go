package specialdatastructure

import (
	"fmt"
	"testing"
)

func Test_heapSort(t *testing.T) {
	// 测试代码
	arr := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 0}
	// arr := []int{3, 27, 36, 27}
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println(arr)
}

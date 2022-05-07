package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	//Integer[] arr = new Integer[];
	//quickSortCore3(arr, 0, arr.length - 1);
	//PrintUtil.printArr(arr);
	//arr := []int{2, 3, 1, 1, 1, 1, 1, 1, 1}
	arr := []int{47, 29, 71, 99, 78, 19, 24, 47}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

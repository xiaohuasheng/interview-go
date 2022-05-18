package specialdatastructure

import (
	"fmt"
	"testing"
)

func Test_heapSort2(t *testing.T) {
	arr := []int{4, 6, 2, 7, 9}
	heapSort2(arr)
	fmt.Println(arr)
}

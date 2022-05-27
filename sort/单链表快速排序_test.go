package sort

import (
	"fmt"
	"testing"
)

func Test_quickSortMain(t *testing.T) {
	myList := new(Node)
	arr := []int{5, 1, 7, 4, 2, 9, 6, 3, 8}
	cur := myList
	for _, val := range arr {
		newNode := &Node{
			Val: val,
		}
		cur.Next = newNode
		cur = cur.Next
	}
	quickSortMain(myList)
	p := myList.Next
	for p != nil {
		fmt.Print(p.Val, " ")
		p = p.Next
	}
	fmt.Println()
}

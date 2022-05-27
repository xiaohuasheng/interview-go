package sort

type Node struct {
	Val  int
	Next *Node
}

func quickSortMain(myList *Node) {
	if myList.Next == nil || myList.Next.Next == nil {
		return
	}
	start := myList.Next // 第一个不用排序吗？
	quick_sort(start, nil)
}

func quick_sort(start *Node, end *Node) {
	if start == end || start.Next == end {
		return
	}
	mid := partion(start, end)
	quick_sort(start, mid)
	quick_sort(mid.Next, end)
}

func partion(start *Node, end *Node) *Node {
	if start == end || start.Next == end {
		return start
	}
	p, q, refer := start, start, start
	for q != end {
		if q.Val < refer.Val {
			p = p.Next
			swap(p, q)
		}
		q = q.Next
	}
	swap(p, refer)
	return p
}

func swap(p, q *Node) {
	temp := p.Val
	p.Val = q.Val
	q.Val = temp
}

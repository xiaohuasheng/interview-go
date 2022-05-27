package specialdatastructure

import "fmt"

// 涉及的知识点，完全二叉树的数组表示
// 完全二叉树的定义：除了最后一层都是满的，所以下沉或上浮操作的时间复杂度都是log(n)
// 技巧：不用数组第一个元素，这样通过根节点找左右子树就有规律

type maxPQ struct {
	Data []int
	Len  int
}

func New(k int) *maxPQ {
	data := make([]int, k+1)
	pq := new(maxPQ)
	pq.Data = data
	pq.Len = 0 // 只是开辟了空间，但还没放数据
	return pq
}

func (p *maxPQ) insert(elem int) {
	p.Len++
	p.Data[p.Len] = elem
	p.swim(p.Len)
}

func (p *maxPQ) swim(k int) {
	// 上浮第k个元素，跟父节点比较，比父节点大就交换
	// k == 1 表示已经到根节点了，无法再上浮
	for k > 1 && p.Data[parent(k)] < p.Data[k] {
		// 交换
		p.Data[parent(k)], p.Data[k] = p.Data[k], p.Data[parent(k)]
		k = parent(k)
	}
}

func parent(k int) int {
	return k / 2
}

func (p *maxPQ) delMax() int {
	max := p.Data[1]
	// 第一个和最后一个元素交换
	p.Data[1], p.Data[p.Len] = p.Data[p.Len], p.Data[1]
	p.Len--
	// 最后一个元素在顶了，不符合定义，要下沉调整
	p.sink(1)
	return max
}

// 下沉第k个元素
func (p *maxPQ) sink(k int) {
	for left(k) < p.Len {
		// 看看左右孩子哪个更大，跟大的比较
		bigger := left(k)
		// 要保证右孩子也不超过长度
		if right(k) > p.Len {
			break
		}
		if p.Data[right(k)] > p.Data[left(k)] {
			bigger = right(k)
		}
		if p.Data[bigger] < p.Data[k] {
			// 父节点更大，不用交换了
			break
		}
		p.Data[k], p.Data[bigger] = p.Data[bigger], p.Data[k]
		k = bigger
	}
}

func left(k int) int {
	return 2 * k
}

func right(k int) int {
	return 2*k + 1
}

// 最大堆，排序是从大到小
func heapSort3() {
	data := []int{7, 9, 4, 10, 3, 5}
	pq := New(len(data))
	// 先构造堆
	// 这里构造用了新的切片，如果我要在原数组上构造呢
	for _, elem := range data {
		pq.insert(elem)
	}
	var sortedData []int
	for pq.Len > 0 {
		elem := pq.delMax()
		sortedData = append(sortedData, elem)
	}
	fmt.Println(sortedData)
}

// 这里构造用了新的切片，如果我要在原数组上构造呢
// 考虑在原数组构造
// 


package stack

import (
	"container/list"
)

func test() {
	// 初始化
	queue := list.New()
	stack := list.New()

	// 入队 入栈
	queue.PushBack(123)
	stack.PushBack(123)

	// 出队 出栈 返回的数据是结构类型 Value 需要断言成相应的类型
	num1 := queue.Front()
	queue.Remove(num1)

	num2 := queue.Back()
	stack.Remove(num2)
}

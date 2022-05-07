package tree

import (
	"container/list"
)

func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

func traversal(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	traversal(node.Left, res)
	traversal(node.Right, res)
}

func preorderTraversal4(root *TreeNode) []int {
	result := make([]int, 0)
	// 少了一步！当 root 为 nil 时的处理
	if root == nil {
		return result
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		// 弹出元素处理
		node := stack.Remove(stack.Back()).(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return result
}

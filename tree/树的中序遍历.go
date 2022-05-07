package tree

import (
	"container/list"
)

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/binary-tree-inorder-traversal/solution/er-cha-shu-de-zhong-xu-bian-li-by-leetcode-solutio/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// 跟着思路写，画个图
func inorderTraversal3(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := list.New()
	cur := root
	for stack.Len() != 0 || cur != nil {
		if cur != nil {
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			// 弹出
			node := stack.Remove(stack.Back()).(*TreeNode)
			result = append(result, node.Val)
			cur = node.Right
		}
	}
	return result
}

package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(nums []int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}
		index := 0
		maxx := nums[0]
		root := &TreeNode{}
		for i := range nums {
			if nums[i] > maxx {
				maxx = nums[i]
				index = i
			}
		}
		root.Val = maxx
		root.Left = dfs(nums[:index])
		root.Right = dfs(nums[index+1:])
		return root
	}
	return dfs(nums)
}

// 作者：mythicmyuu
// 链接：https://leetcode-cn.com/problems/maximum-binary-tree/solution/golang-di-gui-by-mythicmyuu-seya/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

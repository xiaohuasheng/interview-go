package tree

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mp := map[string]int{}
	strMp := map[string]*TreeNode{}
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root == nil {
			return "|"
		}
		var str string
		str += "." + strconv.Itoa(root.Val)
		str += dfs(root.Left)
		str += dfs(root.Right)
		mp[str]++
		strMp[str] = root
		return str
	}
	dfs(root)
	res := []*TreeNode{}
	for i := range mp {
		if mp[i] > 1 {
			res = append(res, strMp[i])
		}
	}
	return res
}

// 作者：mythicmyuu
// 链接：https://leetcode-cn.com/problems/find-duplicate-subtrees/solution/golang-dfs-by-mythicmyuu-iuxu/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

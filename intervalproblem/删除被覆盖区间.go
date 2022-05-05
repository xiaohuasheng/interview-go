package intervalproblem

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	ret := 0
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] >= pre[0] && cur[1] <= pre[1] {
			ret++
			continue
		}
		pre = cur
	}
	return len(intervals) - ret
}

// 作者：QoIrigRcUB
// 链接：https://leetcode-cn.com/problems/remove-covered-intervals/solution/by-qoirigrcub-jnkp/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

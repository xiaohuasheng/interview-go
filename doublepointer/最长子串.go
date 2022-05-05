package doublepointer

func lengthOfLongestSubstring(s string) int {
	location := [255]int{}
	for i := range location {
		location[i] = -1
	}
	maxLen, left := 0, 0

	for i, v := range s {
		if location[v] >= left {
			left = location[v] + 1
		} else if i-left+1 > maxLen {
			maxLen = i - left + 1
		}
		location[v] = i
	}

	return maxLen
}

// 作者：lryong
// 链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/solution/lie-biao-zuo-ha-xi-biao-shuang-zhi-zhen-xjp68/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

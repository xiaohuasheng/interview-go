package string

//给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/implement-strstr
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	next := getNext2(needle)
	bytes := []byte(haystack)
	j := 0
	for i := 0; i < len(bytes); i++ {
		for j > 0 && bytes[i] != bytes[j] {
			//j = next[j]
			j = next[j-1]
		}
		if bytes[i] == bytes[j] {
			j++
		}
		//if j == len(needle) - 1 {
		//	return i - len(needle)
		//}
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

func getNext2(needle string) []int {
	next := make([]int, len(needle))
	// TODO 考虑needle为空的情况
	j := 0
	bytes := []byte(needle)
	for i := 1; i < len(bytes); i++ {
		for j > 0 && bytes[i] != bytes[j] {
			// j = next[j]
			j = next[j-1] // 是回到上一个结果
		}
		if bytes[i] == bytes[j] {
			j++
		}
		next[i] = j
	}
	return next
}

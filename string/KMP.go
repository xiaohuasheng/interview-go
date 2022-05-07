package string

// 方法二: 前缀表无减一或者右移

// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext(next []int, s string) {
	// next是最长相等前后缀的长度
	j := 0 // 因为一个字符的时候，没有前缀也没有后缀
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] { // 为什么要这样跳呢？觉得是利用了前一个结果，归纳法，这不就是动态规划吗？
			// 相等则+1,不等为什么是放在循环里呢？
			// j的含义，既是长度，也是位置，是什么的位置呢？是以i(包括i)为结束符的子串的最长前后缀中前缀末尾的位置
			// 那就可以解释为什么是循环了，从j的定义出发，最长前后缀，s[i]必然等于s[j]，或者是j = 0
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		// 如果 s[i] 与 t[j + 1] 不相同，j就要从next数组里寻找下一个匹配的位置
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		// 如果 s[i] 与 t[j + 1] 相同，那么i 和 j 同时向后移动, i的增加在for循环里
		if haystack[i] == needle[j] {
			j++
		}
		// 如何判断在文本串s里出现了模式串t呢，如果j指向了模式串t的末尾，那么就说明模式串t完全匹配文本串s里的某个子串了
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

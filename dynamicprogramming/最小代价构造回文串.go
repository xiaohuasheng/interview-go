package dynamicprogramming

func minInsertions(s string) int {
	n := len(s)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][i] = 0
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[0][n-1]
}

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// 作者：oliverzero
// 链接：https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/solution/bei-wang-lu-memodptable-by-oliverzero-sl5y/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

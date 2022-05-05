package buyandsellstock

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 作者：carlsun-2
// 链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/dai-ma-sui-xiang-lu-121-mai-mai-gu-piao-odhle/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

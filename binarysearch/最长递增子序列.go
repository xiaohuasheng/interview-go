package dichotomy

func lengthOfLIS(nums []int) int {
	stack := []int{}
	for _, num := range nums {
		l := 0
		for r := len(stack); l < r; {
			mid := (l + r) / 2
			if stack[mid] >= num {
				r = mid
			} else {
				l = mid + 1
			}
		}
		if l == len(stack) {
			stack = append(stack, num)
		} else {
			stack[l] = num
		}
	}
	return len(stack)
}

// 作者：himymBen
// 链接：https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/pythonjavajavascriptgo-lis-zui-chang-sha-e36x/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

// func lengthOfLIS(nums []int ) int {
// 	dp := []int{}
// 	for _, num := range nums {
// 		if len(dp) ==0 || dp[len(dp) - 1] < num {
// 		dp = append(dp, num)
// 	  } else {
// 			l, r := 0, len(dp) - 1
// 			pos := r
// 			for l <= r {
// 				mid := (l + r) >> 1
// 				if dp[mid] >= num {
// 					pos = mid;
// 					r = mid - 1
// 				} else {
// 					l = mid + 1
// 				}
// 			}
// 			dp[pos] = num
// 		}//二分查找
// 	}
// 	  return len(dp)
//   }

//   作者：carlsun-2
//   链接：https://leetcode-cn.com/problems/longest-increasing-subsequence/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-dpzi-i1kh6/
//   来源：力扣（LeetCode）
//   著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

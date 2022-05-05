package dichotomy

// 左闭右闭区间
// https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right { // 我的区间是左闭右闭的，left == right 时有意义，所以这里是<=而不是<
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1 // 因为等于mid时找不到，而区间右边是闭的，所以right变为mid-1而不是mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

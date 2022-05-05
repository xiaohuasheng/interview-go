package array

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/
// 双指针，这里有三个指针
func sortedSquares(nums []int) []int {
	len := len(nums)
	result := make([]int, len)
	left := 0
	right := len - 1
	k := len - 1
	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare > rightSquare {
			result[k] = leftSquare
			left++
		} else {
			result[k] = rightSquare
			right--
		}
		k--
	}
	return result
}

package array

// https://leetcode-cn.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	var slowPointer int
	for fastPointer := 0; fastPointer < len(nums); fastPointer++ {
		// 果然是知易行难，看的时候没注意这里是等于还是不等于，自己写的时候就在想等于的时候怎么处理，没法搞
		// 还把循环结束条件改成了len(nums)-1，经验告诉我，答案都不用额外去判断边界的
		if nums[fastPointer] != val {
			nums[slowPointer] = nums[fastPointer]
			slowPointer++
		}
	}
	return slowPointer
}

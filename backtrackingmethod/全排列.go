package backtrackingmethod

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums, len(nums), []int{})
	return res
}
func backTrack(nums []int, numsLen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
		backTrack(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}

// 作者：carlsun-2
// 链接：https://leetcode-cn.com/problems/permutations/solution/dai-ma-sui-xiang-lu-dai-ni-xue-tou-hui-s-mfrp/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。


0000010 0101100
0101100 0000010
10101100 00000010
每个字节的第一位，叫做msb（most significant bit），
用于标识下一个字节是否还属于这个整数（1：属于；0：不属于）
所以，最后，填充msb后，第二个字节的第一位是0，表示下一个字节不属于这个整数了

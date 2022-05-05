package dichotomy

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, int(1e9)+1
	var check func(t int) int
	check = func(t int) int {
		res := 0
		for _, v := range piles {
			res += v / t
			if v%t > 0 {
				res++
			}
		}
		return res
	}
	for l < r {
		mid := (l + r) >> 1
		if check(mid) <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// 作者：vincent-492
// 链接：https://leetcode-cn.com/problems/koko-eating-bananas/solution/goer-fen-fa-by-vincent-492-sbn9/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

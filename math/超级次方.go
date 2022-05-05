package math

const mod = 1337

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}

func superPow(a int, b []int) int {
	ans := 1
	for i := len(b) - 1; i >= 0; i-- {
		ans = ans * pow(a, b[i]) % mod
		a = pow(a, 10)
	}
	return ans
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/super-pow/solution/chao-ji-ci-fang-by-leetcode-solution-ow8j/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

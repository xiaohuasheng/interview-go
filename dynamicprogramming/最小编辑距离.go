package dynamicprogramming

func minDistance(word1 string, word2 string) int {
	if len(word2)*len(word1) == 0 {
		return len(word1) + len(word2)
	}

	first, second := make([]int, len(word2)+1), make([]int, len(word2)+1)
	for i := range first {
		first[i] = i
	}
	for i := range word1 {
		second[0] = i + 1
		for j := range word2 {
			b := j + 1
			if word1[i] == word2[j] {
				second[b] = first[b-1]
			} else {
				second[b] = min(first[b]+1, min(first[b-1]+1, second[b-1]+1))
			}
		}
		copy(first, second)
	}
	return second[len(word2)]
}

// func min(i,j int)int{
// 	if i>j{
// 		return j
// 	}
// 	return i
// }

// 作者：acrafle
// 链接：https://leetcode-cn.com/problems/edit-distance/solution/dong-tai-gui-hua-gun-dong-shu-zu-by-acra-htab/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

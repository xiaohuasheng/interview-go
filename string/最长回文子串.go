package string

func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        left1, right1 := expandAroundCenter(s, i, i)
        left2, right2 := expandAroundCenter(s, i, i + 1)
        if right1 - left1 > end - start {
            start, end = left1, right1
        }
        if right2 - left2 > end - start {
            start, end = left2, right2
        }
    }
    return s[start:end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
    for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1 , right+1 { }
    return left + 1, right - 1
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
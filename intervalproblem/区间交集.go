package intervalproblem

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}
	a := 0
	b := 0
	count := [][]int{}
	for a < len(firstList) && b < len(secondList) {
		if firstList[a][1] < secondList[b][1] {
			if ext(firstList[a], secondList[b]) != nil {
				count = append(count, ext(firstList[a], secondList[b]))
			}
			a++
		} else {
			if ext(secondList[b], firstList[a]) != nil {
				count = append(count, ext(secondList[b], firstList[a]))
			}
			b++
		}
	}
	return count
}

func ext(a []int, b []int) []int {
	if a[1] < b[0] {
		return nil
	}
	if a[0] >= b[0] {
		return []int{a[0], a[1]}
	}
	if a[0] <= b[0] {
		return []int{b[0], a[1]}
	}
	return nil
}

// 作者：jiu-xin-8v
// 链接：https://leetcode-cn.com/problems/interval-list-intersections/solution/by-jiu-xin-8v-zcod/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

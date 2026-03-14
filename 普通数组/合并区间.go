import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 先按每个区间的start从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := [][]int{}
	for _, interval := range intervals {
		n := len(res)
		// 如果结果为空或者当前区间和上一个区间不重叠，直接加入
		if n == 0 || res[n-1][1] < interval[0] {
			res = append(res, interval)
		} else {
			// 否则两者有重叠，合并区间
			res[n-1][1] = max(res[n-1][1], interval[1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
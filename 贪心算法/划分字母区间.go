package main

func partitionLabels(s string) []int {
	// 记录每个字母最后出现的位置
	last := make([]int, 26)
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	res := make([]int, 0)
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		if last[s[i]-'a'] > end {
			end = last[s[i]-'a']
		}
		// 到达当前片段的最右边界，切分
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

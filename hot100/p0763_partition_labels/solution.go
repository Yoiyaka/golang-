package p0763_partition_labels

func partitionLabels(s string) []int {
	last := [26]int{}
	for i, c := range s {
		last[c-'a'] = i
	}
	result := []int{}
	start, end := 0, 0
	for i, c := range s {
		if last[c-'a'] > end {
			end = last[c-'a']
		}
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}
	return result
}

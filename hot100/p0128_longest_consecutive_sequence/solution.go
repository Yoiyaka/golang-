package p0128_longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, n := range nums {
		set[n] = true
	}
	best := 0
	for n := range set {
		if !set[n-1] {
			cur := n
			streak := 1
			for set[cur+1] {
				cur++
				streak++
			}
			if streak > best {
				best = streak
			}
		}
	}
	return best
}

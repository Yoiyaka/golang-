package p0300_longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, n := range nums {
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := (lo + hi) / 2
			if tails[mid] < n {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if lo == len(tails) {
			tails = append(tails, n)
		} else {
			tails[lo] = n
		}
	}
	return len(tails)
}

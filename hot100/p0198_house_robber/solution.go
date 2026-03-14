package p0198_house_robber

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev2, prev1 := 0, 0
	for _, n := range nums {
		curr := prev1
		if prev2+n > curr {
			curr = prev2 + n
		}
		prev2, prev1 = prev1, curr
	}
	return prev1
}

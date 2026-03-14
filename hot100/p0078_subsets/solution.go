package p0078_subsets

func subsets(nums []int) [][]int {
	result := [][]int{}
	var backtrack func(start int, curr []int)
	backtrack = func(start int, curr []int) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		result = append(result, tmp)
		for i := start; i < len(nums); i++ {
			curr = append(curr, nums[i])
			backtrack(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	backtrack(0, []int{})
	return result
}

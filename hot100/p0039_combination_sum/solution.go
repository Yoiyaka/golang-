package p0039_combination_sum

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	var backtrack func(start, remain int, curr []int)
	backtrack = func(start, remain int, curr []int) {
		if remain == 0 {
			tmp := make([]int, len(curr))
			copy(tmp, curr)
			result = append(result, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] <= remain {
				curr = append(curr, candidates[i])
				backtrack(i, remain-candidates[i], curr)
				curr = curr[:len(curr)-1]
			}
		}
	}
	backtrack(0, target, []int{})
	return result
}

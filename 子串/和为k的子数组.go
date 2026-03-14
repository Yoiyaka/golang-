func subarraySum(nums []int, k int) int {
	count := 0
	preSum := 0
	m := make(map[int]int)
	m[0] = 1 // 前缀和为0出现一次（即前缀本身就是k的情况）

	for _, v := range nums {
		preSum += v
		if freq, ok := m[preSum-k]; ok {
			count += freq
		}
		m[preSum] += 1
	}
	return count
}
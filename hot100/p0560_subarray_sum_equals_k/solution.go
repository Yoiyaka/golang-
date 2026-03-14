package p0560_subarray_sum_equals_k

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0
	prefixSums := map[int]int{0: 1}
	for _, n := range nums {
		sum += n
		count += prefixSums[sum-k]
		prefixSums[sum]++
	}
	return count
}

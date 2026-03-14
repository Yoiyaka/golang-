package p0268_missing_number

func missingNumber(nums []int) int {
	n := len(nums)
	expected := n * (n + 1) / 2
	actual := 0
	for _, v := range nums {
		actual += v
	}
	return expected - actual
}

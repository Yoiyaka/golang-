func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for num, cnt := range m {
		if cnt > len(nums)/2 {
			return num
		}
	}
	return -1
}

/*
func majorityElement(nums []int) int {
    candidate := 0
    count := 0
    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}
*/
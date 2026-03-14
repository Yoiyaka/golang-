func findDuplicate(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for num, cnt := range m {
		if cnt != 1 {
			return num
		}
	}
	return -1 // 如果输入一定有一个只出现一次，可以不要这句
}

/*快慢指针
func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[0]
    // 第一次相遇
    for {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            break
        }
    }
    // 找入口
    fast = nums[0]
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}
*/
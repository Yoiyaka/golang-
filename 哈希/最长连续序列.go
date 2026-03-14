package main

import "fmt"

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0
	for num := range numSet { //遍历哈希表的key，不是原数组
		if !numSet[num-1] { //只从起点开始
			cur := num
			length := 1
			for numSet[cur+1] {
				cur++
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums)) // 输出4
}

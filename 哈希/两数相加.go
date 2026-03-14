package main

import "fmt"

// twoSum 接受一个整数切片和目标值，返回下标的切片
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 创建哈希表
	for i, num := range nums {  // 遍历每个数
		complement := target - num // 目标值 - 当前数 = 另一个数
		if index, found := numMap[complement]; found {
			// 如果哈希表中有 complement，返回两个下标
			return []int{index, i}
		}
		numMap[num] = i // 存当前数到哈希表
	}
	return nil // 没找到返回 nil
}

func main() {
	var n int
	fmt.Print("请输入数组长度: ")
	fmt.Scan(&n)

	nums := make([]int, n)
	fmt.Printf("请输入 %d 个数字（用空格分隔）: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var target int
	fmt.Print("请输入目标数字: ")
	fmt.Scan(&target)

	result := twoSum(nums, target)
	fmt.Println(result)
}

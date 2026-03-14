package main

import (
	"fmt"
	"math"
)

func longestConsecutiveFactors(n int) []int {
	maxLen := 0   // 最长长度
	startNum := 0 // 最佳起始数字

	upper := int(math.Sqrt(float64(n))) + 1

	// 枚举起始数字
	for i := 2; i <= upper; i++ {
		product := 1
		j := i
		for {
			product *= j
			if n%product != 0 {
				break
			}
			if product == n {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
					startNum = i
				}
				break
			}
			j++
			if product > n {
				break
			}
		}
	}

	if maxLen == 0 {
		return []int{n}
	}

	result := make([]int, maxLen)
	for i := 0; i < maxLen; i++ {
		result[i] = startNum + i
	}
	return result
}

func main() {
	var n int
	fmt.Print("请输入一个正整数：")
	fmt.Scan(&n)
	seq := longestConsecutiveFactors(n)
	fmt.Printf("最长连续因数序列为: %v\n", seq)
}

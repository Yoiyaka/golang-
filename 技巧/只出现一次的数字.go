package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	nums := []int{2, 2, 1}
	result := singleNumber(nums)
	fmt.Println(result) // 输出: 1
}

/*出现k次
func singleNumber(nums []int, k int) int {
    result := 0
    for i := 0; i < 32; i++ {
        cnt := 0
        for _, num := range nums {
            // 统计当前bit位出现的次数
            if (num>>i)&1 == 1 {
                cnt++
            }
        }
        if cnt%k != 0 {
            // 这位只在那个出现一次的数上
            result |= 1 << i
        }
    }
    return result
}
*/
/*出现都不一样
func singleNumber(nums []int) int {
    m := make(map[int]int)
    for _, num := range nums {
        m[num]++
    }
    for num, cnt := range m {
        if cnt == 1 {
            return num
        }
    }
    return -1 // 如果输入一定有一个只出现一次，可以不要这句
}
*/

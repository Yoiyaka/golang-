package main

import (
	"fmt"
	"sort"
)

// 字母异位词分组主函数
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		// 1. 转成字符切片并排序
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

		// 2. 排序后的字符切片转回字符串，作为map的key
		key := string(chars)

		// 3. 把原字符串加入该key下的分组
		anagrams[key] = append(anagrams[key], str)
	}

	// 4. 收集所有结果
	result := [][]string{}
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

// main 函数用于测试
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

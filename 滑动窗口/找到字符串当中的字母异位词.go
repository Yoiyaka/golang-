package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := []int{}
	if len(s) < len(p) {
		return res
	}
	// 字母计数，pFreq为目标频率，winFreq为滑动窗口中频率
	pFreq := [26]int{}
	winFreq := [26]int{}

	for _, ch := range p {
		pFreq[ch-'a']++
	}

	for i := 0; i < len(s); i++ {
		// 增加窗口右侧字符
		winFreq[s[i]-'a']++
		// 窗口大小超过p长度时，把左侧字符踢出去
		if i >= len(p) {
			winFreq[s[i-len(p)]-'a']--
		}
		// 比较频率是否一致
		if winFreq == pFreq {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc")) // 输出: [0 6]
	fmt.Println(findAnagrams("abab", "ab"))        // 输出: [0 1 2]
}

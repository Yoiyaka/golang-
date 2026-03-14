func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mapping := []string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}
	var res []string
	var backtrack func(index int, path string)
	backtrack = func(index int, path string) {
		if index == len(digits) {
			res = append(res, path)
			return
		}
		letters := mapping[digits[index]-'0']
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, path+string(letters[i]))
		}
	}
	backtrack(0, "")
	return res
}
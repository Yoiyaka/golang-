package p0017_letter_combinations_of_a_phone_number

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	phoneMap := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	result := []string{}
	var backtrack func(idx int, curr string)
	backtrack = func(idx int, curr string) {
		if idx == len(digits) {
			result = append(result, curr)
			return
		}
		for _, c := range phoneMap[digits[idx]] {
			backtrack(idx+1, curr+string(c))
		}
	}
	backtrack(0, "")
	return result
}

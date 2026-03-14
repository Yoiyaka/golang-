package p0438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return nil
	}
	need := [26]int{}
	window := [26]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
		window[s[i]-'a']++
	}
	result := []int{}
	if need == window {
		result = append(result, 0)
	}
	for i := len(p); i < len(s); i++ {
		window[s[i]-'a']++
		window[s[i-len(p)]-'a']--
		if need == window {
			result = append(result, i-len(p)+1)
		}
	}
	return result
}

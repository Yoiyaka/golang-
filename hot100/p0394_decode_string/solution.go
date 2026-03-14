package p0394_decode_string

func decodeString(s string) string {
	countStack := []int{}
	strStack := []string{}
	cur := ""
	k := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			k = k*10 + int(c-'0')
		} else if c == '[' {
			countStack = append(countStack, k)
			strStack = append(strStack, cur)
			cur = ""
			k = 0
		} else if c == ']' {
			n := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			repeated := ""
			for i := 0; i < n; i++ {
				repeated += cur
			}
			cur = prev + repeated
		} else {
			cur += string(c)
		}
	}
	return cur
}

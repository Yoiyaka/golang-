package p0438_find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		s, p string
		want []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}
	for _, tc := range tests {
		got := findAnagrams(tc.s, tc.p)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("findAnagrams(%q,%q) = %v, want %v", tc.s, tc.p, got, tc.want)
		}
	}
}

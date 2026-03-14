package p0049_group_anagrams

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	normalize := func(groups [][]string) [][]string {
		for i := range groups {
			sort.Strings(groups[i])
		}
		sort.Slice(groups, func(i, j int) bool {
			if len(groups[i]) == 0 {
				return true
			}
			if len(groups[j]) == 0 {
				return false
			}
			return groups[i][0] < groups[j][0]
		})
		return groups
	}
	got := normalize(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	if len(got) != 3 {
		t.Errorf("expected 3 groups, got %d", len(got))
	}
	got2 := groupAnagrams([]string{""})
	if len(got2) != 1 {
		t.Errorf("expected 1 group, got %d", len(got2))
	}
}

package p0015_3sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	normalize := func(res [][]int) [][]int {
		for _, r := range res {
			sort.Ints(r)
		}
		sort.Slice(res, func(i, j int) bool {
			for k := 0; k < len(res[i]) && k < len(res[j]); k++ {
				if res[i][k] != res[j][k] {
					return res[i][k] < res[j][k]
				}
			}
			return false
		})
		return res
	}
	got := normalize(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	want := normalize([][]int{{-1, -1, 2}, {-1, 0, 1}})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("threeSum got %v, want %v", got, want)
	}
	got2 := threeSum([]int{0, 1, 1})
	if len(got2) != 0 {
		t.Errorf("expected empty, got %v", got2)
	}
	got3 := threeSum([]int{0, 0, 0})
	if !reflect.DeepEqual(normalize(got3), normalize([][]int{{0, 0, 0}})) {
		t.Errorf("expected [[0,0,0]], got %v", got3)
	}
}

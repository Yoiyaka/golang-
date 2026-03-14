package p0347_top_k_frequent_elements

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	sort.Ints(got)
	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got2 := topKFrequent([]int{1}, 1)
	if len(got2) != 1 || got2[0] != 1 {
		t.Errorf("got %v, want [1]", got2)
	}
}

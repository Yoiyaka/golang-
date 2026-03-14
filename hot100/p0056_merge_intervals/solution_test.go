package p0056_merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	got := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	want := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	got2 := merge([][]int{{1, 4}, {4, 5}})
	want2 := [][]int{{1, 5}}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %v, want %v", got2, want2)
	}
}

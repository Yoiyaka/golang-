package p0210_course_schedule_ii

import "testing"

func TestFindOrder(t *testing.T) {
	got := findOrder(2, [][]int{{1, 0}})
	if len(got) != 2 {
		t.Errorf("expected order of length 2, got %v", got)
	}
	got2 := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	if len(got2) != 4 {
		t.Errorf("expected order of length 4, got %v", got2)
	}
	got3 := findOrder(2, [][]int{{1, 0}, {0, 1}})
	if len(got3) != 0 {
		t.Errorf("expected empty order for cycle, got %v", got3)
	}
}

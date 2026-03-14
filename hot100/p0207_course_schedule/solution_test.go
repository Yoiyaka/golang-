package p0207_course_schedule

import "testing"

func TestCanFinish(t *testing.T) {
	if !canFinish(2, [][]int{{1, 0}}) {
		t.Error("expected true")
	}
	if canFinish(2, [][]int{{1, 0}, {0, 1}}) {
		t.Error("expected false for cycle")
	}
}

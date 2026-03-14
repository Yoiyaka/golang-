package p0055_jump_game

import "testing"

func TestCanJump(t *testing.T) {
	if !canJump([]int{2, 3, 1, 1, 4}) {
		t.Error("expected true")
	}
	if canJump([]int{3, 2, 1, 0, 4}) {
		t.Error("expected false")
	}
	if !canJump([]int{0}) {
		t.Error("expected true for single element")
	}
}

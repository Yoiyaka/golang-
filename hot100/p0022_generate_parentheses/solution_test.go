package p0022_generate_parentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	got := generateParenthesis(3)
	if len(got) != 5 {
		t.Errorf("expected 5 combinations for n=3, got %d: %v", len(got), got)
	}
	got2 := generateParenthesis(1)
	if len(got2) != 1 || got2[0] != "()" {
		t.Errorf("expected ['()'], got %v", got2)
	}
}

package p0017_letter_combinations_of_a_phone_number

import "testing"

func TestLetterCombinations(t *testing.T) {
	got := letterCombinations("23")
	if len(got) != 9 {
		t.Errorf("expected 9 combinations, got %d: %v", len(got), got)
	}
	if letterCombinations("") != nil {
		t.Error("expected nil for empty input")
	}
	got2 := letterCombinations("2")
	if len(got2) != 3 {
		t.Errorf("expected 3 combinations for '2', got %d", len(got2))
	}
}

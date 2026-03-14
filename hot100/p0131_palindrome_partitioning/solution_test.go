package p0131_palindrome_partitioning

import "testing"

func TestPartition(t *testing.T) {
	got := partition("aab")
	if len(got) != 2 {
		t.Errorf("expected 2 partitions, got %d: %v", len(got), got)
	}
	got2 := partition("a")
	if len(got2) != 1 {
		t.Errorf("expected 1 partition for 'a', got %d", len(got2))
	}
}

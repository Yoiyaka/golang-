package p0416_partition_equal_subset_sum

import "testing"

func TestCanPartition(t *testing.T) {
	if !canPartition([]int{1, 5, 11, 5}) {
		t.Error("expected true")
	}
	if canPartition([]int{1, 2, 3, 5}) {
		t.Error("expected false")
	}
}

package p0763_partition_labels

import (
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	got := partitionLabels("ababcbacadefegdehijhklij")
	want := []int{9, 7, 8}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
	got2 := partitionLabels("eccbbbbdec")
	want2 := []int{10}
	if !reflect.DeepEqual(got2, want2) {
		t.Errorf("got %v, want %v", got2, want2)
	}
}

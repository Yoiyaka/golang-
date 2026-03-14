package p0023_merge_k_sorted_lists

import (
	"reflect"
	"testing"
)

func makeList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestMergeKLists(t *testing.T) {
	lists := []*ListNode{
		makeList([]int{1, 4, 5}),
		makeList([]int{1, 3, 4}),
		makeList([]int{2, 6}),
	}
	got := listToSlice(mergeKLists(lists))
	want := []int{1, 1, 2, 3, 4, 4, 5, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	got2 := mergeKLists([]*ListNode{})
	if got2 != nil {
		t.Error("expected nil")
	}
}

package p0148_sort_list

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

func TestSortList(t *testing.T) {
	got := listToSlice(sortList(makeList([]int{4, 2, 1, 3})))
	if !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("got %v, want [1,2,3,4]", got)
	}
	got2 := listToSlice(sortList(makeList([]int{-1, 5, 3, 4, 0})))
	if !reflect.DeepEqual(got2, []int{-1, 0, 3, 4, 5}) {
		t.Errorf("got %v, want [-1,0,3,4,5]", got2)
	}
}

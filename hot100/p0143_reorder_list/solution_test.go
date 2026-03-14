package p0143_reorder_list

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

func TestReorderList(t *testing.T) {
	head := makeList([]int{1, 2, 3, 4})
	reorderList(head)
	got := listToSlice(head)
	if !reflect.DeepEqual(got, []int{1, 4, 2, 3}) {
		t.Errorf("got %v, want [1,4,2,3]", got)
	}

	head2 := makeList([]int{1, 2, 3, 4, 5})
	reorderList(head2)
	got2 := listToSlice(head2)
	if !reflect.DeepEqual(got2, []int{1, 5, 2, 4, 3}) {
		t.Errorf("got %v, want [1,5,2,4,3]", got2)
	}
}

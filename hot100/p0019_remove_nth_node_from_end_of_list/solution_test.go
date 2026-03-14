package p0019_remove_nth_node_from_end_of_list

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

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		input []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	}
	for _, tc := range tests {
		got := listToSlice(removeNthFromEnd(makeList(tc.input), tc.n))
		if !reflect.DeepEqual(got, tc.want) {
			if len(got) == 0 && len(tc.want) == 0 {
				continue
			}
			t.Errorf("removeNthFromEnd(%v,%d) = %v, want %v", tc.input, tc.n, got, tc.want)
		}
	}
}

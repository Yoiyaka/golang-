package p0141_linked_list_cycle

import "testing"

func TestHasCycle(t *testing.T) {
	// no cycle
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	if hasCycle(head) {
		t.Error("expected no cycle")
	}

	// with cycle
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	if !hasCycle(n1) {
		t.Error("expected cycle")
	}
}

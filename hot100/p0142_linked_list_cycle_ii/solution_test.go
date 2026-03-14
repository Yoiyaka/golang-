package p0142_linked_list_cycle_ii

import "testing"

func TestDetectCycle(t *testing.T) {
	// no cycle
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	if detectCycle(head) != nil {
		t.Error("expected nil")
	}

	// cycle at pos 1
	n1 := &ListNode{Val: 3}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 0}
	n4 := &ListNode{Val: -4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	got := detectCycle(n1)
	if got != n2 {
		t.Errorf("expected cycle entry at node with val 2, got %v", got)
	}
}

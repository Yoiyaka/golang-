package p0160_intersection_of_two_linked_lists

import "testing"

func TestGetIntersectionNode(t *testing.T) {
	// intersect at node with val 8
	intersect := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	headA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersect}}
	headB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersect}}}
	got := getIntersectionNode(headA, headB)
	if got != intersect {
		t.Errorf("expected intersection node with val 8")
	}

	// no intersection
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	if getIntersectionNode(a, b) != nil {
		t.Error("expected nil")
	}
}

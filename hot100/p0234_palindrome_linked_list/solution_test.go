package p0234_palindrome_linked_list

import "testing"

func makeList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func TestIsPalindrome(t *testing.T) {
	if !isPalindrome(makeList([]int{1, 2, 2, 1})) {
		t.Error("expected true for [1,2,2,1]")
	}
	if isPalindrome(makeList([]int{1, 2})) {
		t.Error("expected false for [1,2]")
	}
	if !isPalindrome(makeList([]int{1})) {
		t.Error("expected true for [1]")
	}
}

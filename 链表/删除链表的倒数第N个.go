/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// fast先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// slow和fast一起走，直到fast到尾巴
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除slow后面那个节点
	slow.Next = slow.Next.Next

	return dummy.Next
}
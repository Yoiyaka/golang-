/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head

	// 第一阶段：快慢指针查找是否有环（相遇点）
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 如果没有环
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 第二阶段：一指针回头，然后齐步走到入环点
	p1 := head
	p2 := slow
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
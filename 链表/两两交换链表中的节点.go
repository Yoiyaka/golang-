/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		// 交换
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// 推进prev和head，准备下一组
		prev = first
		head = first.Next
	}
	return dummy.Next
}
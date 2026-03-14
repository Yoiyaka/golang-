/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		next := curr.Next // 先保存下一个节点
		curr.Next = prev  // 当前节点反转指向前一个
		prev = curr       // prev 前进到当前
		curr = next       // curr 前进到原来下一个
	}
	return prev
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 辅助节点，方便操作头部
	dummy := &ListNode{Next: head}
	pre := dummy

	for {
		tail := pre
		// 检查后面是否有k个节点
		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}

		next := tail.Next
		// 翻转pre.Next-tail
		headK := pre.Next
		prev := next
		curr := headK
		for curr != next {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}
		// pre 的 next 指向新的头
		pre.Next = tail
		// 下轮pre移动到本段末尾（之前的headK）
		pre = headK
	}
	return dummy.Next
}
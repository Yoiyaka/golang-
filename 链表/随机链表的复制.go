/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 第一步：克隆节点穿插在原节点后面
	cur := head
	for cur != nil {
		copy := &Node{Val: cur.Val}
		copy.Next = cur.Next
		cur.Next = copy
		cur = copy.Next
	}

	// 第二步：处理 random 指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 第三步：拆分新旧链表
	cur = head
	copyHead := head.Next
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		cur = cur.Next
	}
	return copyHead
}
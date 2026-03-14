package p0114_flatten_binary_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			rightmost := curr.Left
			for rightmost.Right != nil {
				rightmost = rightmost.Right
			}
			rightmost.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}

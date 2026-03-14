package p0110_balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := height(node.Left)
	if l == -1 {
		return -1
	}
	r := height(node.Right)
	if r == -1 {
		return -1
	}
	diff := l - r
	if diff < -1 || diff > 1 {
		return -1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

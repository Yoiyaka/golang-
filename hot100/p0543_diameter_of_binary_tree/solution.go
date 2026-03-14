package p0543_diameter_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiam := 0
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		if l+r > maxDiam {
			maxDiam = l + r
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	depth(root)
	return maxDiam
}

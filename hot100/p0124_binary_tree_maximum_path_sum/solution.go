package p0124_binary_tree_maximum_path_sum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftGain := max(dfs(node.Left), 0)
		rightGain := max(dfs(node.Right), 0)
		pathSum := node.Val + leftGain + rightGain
		if pathSum > maxSum {
			maxSum = pathSum
		}
		if leftGain > rightGain {
			return node.Val + leftGain
		}
		return node.Val + rightGain
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	// 返回当前节点的最大深度，并更新全局直径
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := dfs(node.Left)
		rightDepth := dfs(node.Right)
		// 每层都更新最大直径
		if leftDepth+rightDepth > diameter {
			diameter = leftDepth + rightDepth
		}
		// 返回当前节点的最大深度
		return max(leftDepth, rightDepth) + 1
	}
	dfs(root)
	return diameter
}

// 工具函数：返回两个数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
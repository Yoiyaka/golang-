/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	maxSum := -1 << 31 // 最小int，保证更新
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0) // 如果子树贡献为负则取0
		right := max(dfs(node.Right), 0)
		// 以当前节点为最高点的最大路径和，可能经过左右两侧
		currSum := node.Val + left + right
		if currSum > maxSum {
			maxSum = currSum
		}
		// 只能返回包含当前节点和左或右的最大路径分支给父节点
		return node.Val + max(left, right)
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
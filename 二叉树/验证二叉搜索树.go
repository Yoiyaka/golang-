/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// 辅助递归函数，传入节点及其值域范围的上下界
	var helper func(node *TreeNode, lower, upper int64) bool
	helper = func(node *TreeNode, lower, upper int64) bool {
		if node == nil {
			return true
		}
		val := int64(node.Val)
		if val <= lower || val >= upper {
			return false
		}
		return helper(node.Left, lower, val) && helper(node.Right, val, upper)
	}
	// 使用 int64 防止边界溢出
	return helper(root, -1<<63, 1<<63-1)
}
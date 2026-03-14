/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	// 辅助递归函数，判断两子树是否镜像对称
	var isMirror func(left, right *TreeNode) bool
	isMirror = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		// 当前节点值相等，且左右各自对称
		return left.Val == right.Val 
		&& isMirror(left.Left, right.Right) 
		&& isMirror(left.Right, right.Left)
	}
	return isMirror(root.Left, root.Right)
}
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 递归交换左右子树
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
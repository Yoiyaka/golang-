/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var count int
	var result int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil || count >= k {
			return
		}
		inorder(node.Left)
		count++
		if count == k {
			result = node.Val
			return
		}
		inorder(node.Right)
	}

	inorder(root)
	return result
}
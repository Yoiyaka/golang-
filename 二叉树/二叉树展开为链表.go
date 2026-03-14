/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 原地先序遍历展开二叉树为链表
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 递归展开左右子树
	flatten(root.Left)
	flatten(root.Right)

	// 将左子树插入到右子树位置
	left := root.Left
	right := root.Right

	if left != nil {
		root.Right = left
		root.Left = nil

		// 找到左子树最右节点，将原右子树接到它后面
		curr := root.Right
		for curr.Right != nil {
			curr = curr.Right
		}
		curr.Right = right
	}
}
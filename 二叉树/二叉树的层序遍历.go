/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var levelorder func(node *TreeNode, level int)
	levelorder = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 如果当前层还没有被初始化，则创建一个新的层
		if len(result) <= level {
			result = append(result, []int{})
		}
		// 将当前节点值添加到对应层中
		result[level] = append(result[level], node.Val)
		levelorder(node.Left, level+1)
		levelorder(node.Right, level+1)
	}
	levelorder(root, 0)
	return result
}
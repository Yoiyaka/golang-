/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	preSum := map[int]int{0: 1} // 前缀和为0出现一次(起点)
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, currSum int) int {
		if node == nil {
			return 0
		}
		currSum += node.Val
		count := preSum[currSum-targetSum] // 查找有多少条路径以当前节点为终点，和为targetSum
		preSum[currSum]++
		count += dfs(node.Left, currSum)
		count += dfs(node.Right, currSum)
		preSum[currSum]-- // 状态回溯
		return count
	}
	return dfs(root, 0)
}
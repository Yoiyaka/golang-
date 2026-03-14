package p0437_path_sum_iii

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	prefixSums := map[int]int{0: 1}
	return dfs(root, 0, targetSum, prefixSums)
}

func dfs(node *TreeNode, currSum, target int, prefixSums map[int]int) int {
	if node == nil {
		return 0
	}
	currSum += node.Val
	count := prefixSums[currSum-target]
	prefixSums[currSum]++
	count += dfs(node.Left, currSum, target, prefixSums)
	count += dfs(node.Right, currSum, target, prefixSums)
	prefixSums[currSum]--
	return count
}

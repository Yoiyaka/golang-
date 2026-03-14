/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS 层序遍历，记录每层最后一个节点的值
func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			// 记录每层最后一个节点的值
			if i == levelSize-1 {
				result = append(result, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return result
}

/*
func rightSideView(root *TreeNode) []int {
    var result []int
    var dfs func(node *TreeNode, depth int)
    dfs = func(node *TreeNode, depth int) {
        if node == nil {
            return
        }
        // 首次到达这一层，则记录当前节点
        if depth == len(result) {
            result = append(result, node.Val)
        }
        // 优先遍历右边
        dfs(node.Right, depth+1)
        dfs(node.Left, depth+1)
    }
    dfs(root, 0)
    return result
}
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// 用map优化inorder查找，提升效率到O(n)
	idxMap := make(map[int]int, len(inorder))
	for i, v := range inorder {
		idxMap[v] = i
	}

	var build func(preLeft, preRight, inLeft, inRight int) *TreeNode
	build = func(preLeft, preRight, inLeft, inRight int) *TreeNode {
		if preLeft > preRight {
			return nil
		}
		rootVal := preorder[preLeft]
		root := &TreeNode{Val: rootVal}
		inRootIdx := idxMap[rootVal]
		leftSize := inRootIdx - inLeft
		root.Left = build(preLeft+1, preLeft+leftSize, inLeft, inRootIdx-1)
		root.Right = build(preLeft+leftSize+1, preRight, inRootIdx+1, inRight)
		return root
	}
	return build(0, len(preorder)-1, 0, len(inorder)-1)
}
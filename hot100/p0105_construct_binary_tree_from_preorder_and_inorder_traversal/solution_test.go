package p0105_construct_binary_tree_from_preorder_and_inorder_traversal

import "testing"

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, inorder(root.Right)...)
	return res
}

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorderArr := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorderArr)
	got := inorder(root)
	want := []int{9, 3, 15, 20, 7}
	for i, v := range want {
		if got[i] != v {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

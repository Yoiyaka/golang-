package p0226_invert_binary_tree

import "testing"

func makeTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := inorder(root.Left)
	result = append(result, root.Val)
	result = append(result, inorder(root.Right)...)
	return result
}

func TestInvertTree(t *testing.T) {
	root := makeTree([]interface{}{4, 2, 7, 1, 3, 6, 9})
	invertTree(root)
	// after invert inorder should be [9,7,6,4,3,2,1]
	got := inorder(root)
	want := []int{9, 7, 6, 4, 3, 2, 1}
	for i, v := range want {
		if got[i] != v {
			t.Errorf("inorder after invert: got %v, want %v", got, want)
			break
		}
	}
}

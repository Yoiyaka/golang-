package p0297_serialize_and_deserialize_binary_tree

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

func TestSerializeDeserialize(t *testing.T) {
	c := Constructor()
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
	}
	serialized := c.serialize(root)
	restored := c.deserialize(serialized)

	orig := inorder(root)
	rest := inorder(restored)
	if len(orig) != len(rest) {
		t.Errorf("inorder mismatch: %v vs %v", orig, rest)
		return
	}
	for i := range orig {
		if orig[i] != rest[i] {
			t.Errorf("inorder mismatch at %d: %v vs %v", i, orig, rest)
		}
	}

	if c.deserialize("null") != nil {
		t.Error("expected nil for null")
	}
}

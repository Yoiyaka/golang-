package p0297_serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	left := c.serialize(root.Left)
	right := c.serialize(root.Right)
	return strconv.Itoa(root.Val) + "," + left + "," + right
}

func (c *Codec) deserialize(data string) *TreeNode {
	parts := strings.Split(data, ",")
	idx := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if idx >= len(parts) || parts[idx] == "null" {
			idx++
			return nil
		}
		val, _ := strconv.Atoi(parts[idx])
		idx++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}
	return build()
}

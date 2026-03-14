package p0230_kth_smallest_element_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root
	count := 0
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		if count == k {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}

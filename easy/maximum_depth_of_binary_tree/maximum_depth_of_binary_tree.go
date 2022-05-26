package maximum_depth_of_binary_tree

import . "hamologist.com/leetcode-go/shared"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := 0

	var inner func(int, *TreeNode)
	inner = func(level int, node *TreeNode) {
		if node == nil {
			if level > result {
				result = level
			}
			return
		}

		inner(level+1, node.Left)
		inner(level+1, node.Right)
	}
	inner(0, root)

	return result
}

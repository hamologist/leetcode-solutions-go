package shared

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTreeNode(slice []interface{}) *TreeNode {
	var root *TreeNode = nil
	var current *TreeNode
	var queue []*TreeNode
	checkLeft, checkRight := true, true

	for _, num := range slice {
		asInt, isInt := num.(int)
		if root == nil && isInt {
			root = &TreeNode{Val: asInt}
			current = root
			continue
		}

		if checkLeft {
			checkLeft = false
			if num != nil && isInt {
				current.Left = &TreeNode{Val: asInt}
				queue = append(queue, current.Left)
			}
			continue
		}

		if checkRight {
			checkRight = false
			if num != nil && isInt {
				current.Right = &TreeNode{Val: asInt}
				queue = append(queue, current.Right)
			}
			continue
		}

		checkLeft, checkRight = true, true
		current = queue[0]
		queue = queue[1:]
	}

	return root
}

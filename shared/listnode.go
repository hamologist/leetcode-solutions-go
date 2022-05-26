package shared

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateListNode(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}
	var head *ListNode = nil
	var current *ListNode

	for _, num := range slice {
		if head == nil {
			head = &ListNode{
				Val:  num,
				Next: nil,
			}
			current = head
			continue
		}

		current.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		current = current.Next
	}

	return head
}

func ListToSlice(node *ListNode) []int {
	var result []int

	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}

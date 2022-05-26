package delete_node_in_a_linked_list

import . "hamologist.com/leetcode-go/shared"

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

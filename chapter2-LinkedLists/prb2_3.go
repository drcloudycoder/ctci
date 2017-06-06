package linkedlist

/*
  PROBLEM: To delete a node in the middle (anything but first or last node) of a linked list
  Time Complexity: O(n)
  Space Complexity: O(1)
*/

func deleteMiddleNode(node *Node) bool {
	if node.next == nil || node == nil {
		return false
	}
	// Copy data from next node to current node
	node.data = node.next.data
	// Delete the next node
	node.next = node.next.next
	return true
}

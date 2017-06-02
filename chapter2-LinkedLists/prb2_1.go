package linkedlist

/*
  PROBLEM: To remove duplicates from an unsorted linked list
  Time Complexity: O(n)
  Space Complexity: O(n)
*/
func (list *List) removeDuplicates() {
	current := list.Head()
	var previous *Node
	previous = nil
	m := make(map[int]int)

	for current != nil {
		if _, alreadyPresent := m[current.data]; alreadyPresent {
			previous.next = current.next
		} else {
			m[current.data] = 1
			previous = current
		}
		current = current.next
	}
}

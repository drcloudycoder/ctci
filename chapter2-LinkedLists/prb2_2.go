package linkedlist

/*
  PROBLEM: To find the kth to last element of a singly linked list
  Time Complexity: O(n)
  Space Complexity: O(1)
*/

func (list *List) findKthElementToLast(k int) int {
	var slowRunner *Node = nil
	var fastRunner *Node = nil

	for i := 0; i < k; i++ {
		if fastRunner == nil {
			fastRunner = list.Head()
		} else {
			fastRunner = fastRunner.next
		}
	}

	for fastRunner != nil {
		fastRunner = fastRunner.next
		if slowRunner == nil {
			slowRunner = list.Head()
		} else {
			slowRunner = slowRunner.next
		}
	}

	return slowRunner.data
}

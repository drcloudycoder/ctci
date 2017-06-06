package linkedlist

import (
	"fmt"
	"testing"
)

/*
  PROBLEM: To delete a node in the middle (anything but first or last node) of a linked list
*/
func TestDeleteMiddleNode(t *testing.T) {
	fmt.Println("---------- Testing deleteMiddleNode ----------")
	list := New()
	e := &Node{data: 5, next: nil}
	d := &Node{data: 4, next: e}
	c := &Node{data: 3, next: d}
	b := &Node{data: 2, next: c}
	a := &Node{data: 1, next: b}
	list.head = a
	list.size = 5
	if result := deleteMiddleNode(c); result {
		current := list.Head()
		for current != nil {
			fmt.Println("Element: ", current.data)
			current = current.next
		}
	} else {
		fmt.Println("deleteMiddleNode function returned false")
	}
}

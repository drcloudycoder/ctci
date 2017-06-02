package linkedlist

import (
	"fmt"
	"testing"
)

/*
  PROBLEM: To remove duplicates from an unsorted linked list
*/
func TestRemoveDuplicates(t *testing.T) {
	fmt.Println("---------- Testing removeDuplicates ----------")
	list := New()
	d := &Node{data: 1, next: nil}
	c := &Node{data: 2, next: d}
	b := &Node{data: 2, next: c}
	a := &Node{data: 4, next: b}
	list.head = a
	list.size = 4
	list.removeDuplicates()
	current := list.Head()
	for current != nil {
		fmt.Println("Element: ", current.data)
		current = current.next
	}
}

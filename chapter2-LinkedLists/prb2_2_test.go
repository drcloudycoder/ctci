package linkedlist

import (
	"fmt"
	"testing"
)

/*
  PROBLEM: To find the kth to last element of a singly linked list
*/
func TestFindKthElementToLast(t *testing.T) {
	fmt.Println("---------- Testing findKthElementToLast ----------")
	list := New()
	e := &Node{data: 5, next: nil}
	d := &Node{data: 4, next: e}
	c := &Node{data: 3, next: d}
	b := &Node{data: 2, next: c}
	a := &Node{data: 1, next: b}
	list.head = a
	list.size = 5
	fmt.Println("2nd element from last: ", list.findKthElementToLast(2))
}

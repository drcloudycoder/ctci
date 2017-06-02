package linkedlist

// Represents structure of single item of list
type Node struct {
	data int
	next *Node
}

// Represents structure of the List
type List struct {
	head *Node
	size int
}

/*
  Create an empty list
*/
func New() *List {
	return &List{size: 0}
}

/*
  Get the list's current length
*/
func (list *List) Size() int {
	return list.size
}

/*
  Get the head of the list
*/
func (list *List) Head() *Node {
	return list.head
}

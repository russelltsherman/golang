package linkedlist

// List the linked list structure
type List struct {
	head *Node
	tail *Node
}

// Node a node of a linked list structure
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// First return the first node in the list
func (l *List) First() *Node {
	return l.head
}

// Last return the first node in the list
func (l *List) Last() *Node {
	return l.tail
}

// Value return the value of the node
func (n *Node) Value() int {
	return n.value
}

// Next return the next node in the list
func (n *Node) Next() *Node {
	return n.next
}

// Prev return the previous node in the list
func (n *Node) Prev() *Node {
	return n.prev
}

// Push a node into the list
func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}
	l.tail = node
}

package structs

// List of nodes
type List struct {
	root *Node
}

// Node of the list
type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

// Append Node at the end of the list
func (l *List) Append(n *Node) {
	if l.root == nil {
		l.root = n
		return
	}

	currNode := l.root

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Next = n
}

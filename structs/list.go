package structs

import (
	"errors"
	"reflect"
)

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
func (l *List) Append(v interface{}) error {
	nodeNew := &Node{
		Next:  nil,
		Prev:  nil,
		Value: v,
	}

	if l.root == nil {
		l.root = nodeNew
		return nil
	}

	currNode := l.root

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Next = nodeNew
	currNode.Next.Prev = currNode

	return nil
}

// Remove Node by value
func (l *List) Remove(v interface{}) error {
	if l.root == nil {
		return errors.New("list is empty")
	}

	node := &Node{
		Next:  nil,
		Value: v,
	}

	currNode := l.root

	for {
		if reflect.DeepEqual(currNode.Value, node.Value) {
			currNode.Prev.Next = currNode.Next
			if currNode.Next != nil {
				currNode.Next.Prev = currNode.Prev
			}
			currNode.Next, currNode.Prev = nil, nil
			return nil
		}
		if currNode.Next == nil {
			break
		}
		currNode = currNode.Next
	}
	return errors.New("cant find node with target value")
}

func (l *List) Insert(index int, v interface{}) error {
	return nil
}

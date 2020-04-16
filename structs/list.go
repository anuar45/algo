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

// NewList creates new empty list
func NewList() *List {
	return &List{nil}
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

// RemoveByValue removes Nodes by value
func (l *List) RemoveByValue(v interface{}) error {
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

// Insert Node at specified position
func (l *List) Insert(index uint, v interface{}) error {
	nodeNew := &Node{Value: v}
	if l.root == nil {
		if index == 0 {
			l.root = nodeNew
			return nil
		}
		return errors.New("list empty")
	}

	currNode := l.root

	for i := 0; i < int(index); i++ {
		if currNode.Next == nil {
			return errors.New("index out of range")
		}
		currNode = currNode.Next
	}

	if currNode.Prev != nil {
		nodeNew.Prev = currNode.Prev
		nodeNew.Prev.Next = nodeNew
	}

	currNode.Prev = nodeNew
	nodeNew.Next = currNode

	return nil
}

// Last returns last item value
func (l *List) Last() (interface{}, error) {
	if l.root == nil {
		return nil, errors.New("list is empty")
	}

	currNode := l.root

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	return currNode.Value, nil
}

// RemoveLast removes last Node in the list
func (l *List) RemoveLast() error {
	if l.root == nil {
		return errors.New("list is empty")
	}

	currNode := l.root

	if currNode.Next == nil {
		l.root = nil
		return nil
	}

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	currNode.Prev.Next = nil
	currNode.Prev = nil

	return nil
}

// Len returns lenght of list
func (l *List) Len() int {
	if l.root == nil {
		return 0
	}

	count := 1
	currNode := l.root

	for currNode.Next != nil {
		currNode = currNode.Next
	}

	return count
}

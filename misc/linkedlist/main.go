package main

import (
	"fmt"
)

type List struct {
	Head *Node
}

type Node struct {
	Value int
	Next  *Node
}

func main() {
	fmt.Println("vim-go")

	n9 := &Node{3, nil}
	n8 := &Node{3, n9}
	n7 := &Node{4, n8}
	n6 := &Node{3, n7}
	n5 := &Node{3, n6}
	n4 := &Node{3, n5}
	n3 := &Node{2, n4}
	n2 := &Node{3, n3}
	head := &Node{3, n2}

	l := List{head}

	fmt.Println(l.String())
	l.Remove(3)
	fmt.Println(l.String())
}

func (l *List) Remove(n int) {

	for l.Head.Value == n {
		l.Head = l.Head.Next
	}

	curNode := l.Head.Next
	prevNode := l.Head
	for curNode != nil {
		if curNode.Value == n {
			if prevNode != nil {
				prevNode.Next = curNode.Next
			}
		} else {
			prevNode = curNode
		}
		curNode = curNode.Next
	}
}

func (l *List) String() string {
	var values []int
	curNode := l.Head
	for curNode != nil {
		values = append(values, curNode.Value)
		curNode = curNode.Next
	}

	return fmt.Sprint(values)
}

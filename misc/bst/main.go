package main

import "fmt"

func main() {
	root := &Node{Value: 100}
	root.Insert(200)
	root.Insert(300)
	root.Insert(20)
	root.Insert(400)
	root.Insert(50)
	root.Insert(70)
	root.Insert(600)
	root.Insert(10)

	root.Print()
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(v int) {
	if v > n.Value {
		if n.Right == nil {
			n.Right = &Node{Value: v}
		} else {
			n.Right.Insert(v)
		}
	} else if v < n.Value {
		if n.Left == nil {
			n.Left = &Node{Value: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

func (n *Node) Print() {
	if n == nil {
		return
	} else {
		fmt.Println(n.Value)
		n.Left.Print()
		n.Right.Print()
	}
}

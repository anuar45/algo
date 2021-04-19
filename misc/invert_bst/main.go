package main

import "fmt"

func main() {
	root := &TreeNode{Val: 100}
	root.Insert(200)
	root.Insert(300)
	root.Insert(20)
	root.Insert(400)
	root.Insert(50)
	root.Insert(70)
	root.Insert(600)
	root.Insert(10)

	root.Print()

	invertTree(root)

	root.Print()
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left :=  invertTree(root.Left)
	right := invertTree(root.Right)

	root.Right = left
	root.Left = right

	return root
}

func (n *TreeNode) Insert(v int) {
	if v > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: v}
		} else {
			n.Right.Insert(v)
		}
	} else if v < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: v}
		} else {
			n.Left.Insert(v)
		}
	}
}

func (n *TreeNode) Print() {
	if n == nil {
		return
	} else {
		fmt.Println(n.Val)
		n.Left.Print()
		n.Right.Print()
	}
}

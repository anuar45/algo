package main

import "fmt"

func main() {
	node1 := &ListNode{1, nil}
	node2 := &ListNode{2, nil}
	node3 := &ListNode{3, nil}
	node4 := &ListNode{4, nil}
	node5 := &ListNode{5, nil}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	deleteNode(node3)

	printList(node1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func printList(head *ListNode) {
	curNode := head
	for curNode != nil {
		fmt.Println(curNode.Val)
		curNode = curNode.Next
	}
}

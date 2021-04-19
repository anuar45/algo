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

	head := removeNthFromEnd(node1, 5)

	printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var nodeArr []*ListNode

	for curNode := head; curNode != nil; curNode = curNode.Next {
		nodeArr = append(nodeArr, curNode)
	}

	if len(nodeArr) == 1 {
		return nil
	}

	if nodeArr[len(nodeArr)-n].Next == nil {
		nodeArr[len(nodeArr)-n-1].Next = nil
	} else if n == len(nodeArr) {
		head = nodeArr[1]
	} else {
		nodeArr[len(nodeArr)-n-1].Next = nodeArr[len(nodeArr)-n].Next
	}

	return head
}

func printList(head *ListNode) {
	curNode := head
	for curNode != nil {
		fmt.Println(curNode.Val)
		curNode = curNode.Next
	}
}

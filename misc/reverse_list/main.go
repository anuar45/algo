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

	head := reverseList(node1)

	printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var nodeArr []*ListNode

	for curNode := head; curNode != nil; curNode = curNode.Next {
		nodeArr = append(nodeArr, curNode)
	}

	if len(nodeArr) == 0 || len(nodeArr) == 1 {
		return head
	}

	head = nodeArr[len(nodeArr)-1]
	nodeArr[0].Next = nil

	for i := len(nodeArr) - 1; i > 0; i-- {
		nodeArr[i].Next = nodeArr[i-1]
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

type Stack struct {
	items []*ListNode
}

func (s *Stack) Pop() *ListNode {
	if len(s.items) == 0 {
		return nil
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return item
}

func (s *Stack) Push(l *ListNode) {
	s.items = append(s.items, l)
}
